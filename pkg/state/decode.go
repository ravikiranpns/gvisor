// Copyright 2018 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package state

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"reflect"

	"gvisor.dev/gvisor/pkg/state/wire"
)

// internalCallback is a interface called on object completion.
//
// There are two implementations: objectDecodeState & userCallback.
type internalCallback interface {
	// source returns the dependent object. May be nil.
	source() *objectDecodeState

	// callbackRun executes the callback.
	callbackRun()
}

// userCallback is an implementation of internalCallback.
type userCallback func()

// source implements internalCallback.source.
func (userCallback) source() *objectDecodeState {
	return nil
}

// callbackRun implements internalCallback.callbackRun.
func (uc userCallback) callbackRun() {
	uc()
}

// objectDecodeState represents an object that may be in the process of being
// decoded. Specifically, it represents either a decoded object, or an an
// interest in a future object that will be decoded. When that interest is
// registered (via register), the storage for the object will be created, but
// it will not be decoded until the object is encountered in the stream.
type objectDecodeState struct {
	// id is the id for this object.
	id objectID

	// typ is the id for this typeID. This may be zero if this is not a
	// type-registered structure.
	typ typeID

	// obj is the object. This may or may not be valid yet, depending on
	// whether complete returns true. However, regardless of whether the
	// object is valid, obj contains a final storage location for the
	// object. This is immutable.
	//
	// Note that this must be addressable (obj.Addr() must not panic).
	//
	// The obj passed to the decode methods below will equal this obj only
	// in the case of decoding the top-level object. However, the passed
	// obj may represent individual fields, elements of a slice, etc. that
	// are effectively embedded within the reflect.Value below but with
	// distinct types.
	obj reflect.Value

	// blockedBy is the number of dependencies this object has.
	blockedBy int

	// callbacksInline is inline storage for callbacks.
	callbacksInline [2]internalCallback

	// callbacks is a set of callbacks to execute on load.
	callbacks []internalCallback

	completeEntry
}

// addCallback adds a callback to the objectDecodeState.
func (ods *objectDecodeState) addCallback(ic internalCallback) {
	if ods.callbacks == nil {
		ods.callbacks = ods.callbacksInline[:0]
	}
	ods.callbacks = append(ods.callbacks, ic)
}

// findCycleFor returns when the given object is found in the blocking set.
func (ods *objectDecodeState) findCycleFor(target *objectDecodeState) []*objectDecodeState {
	for _, ic := range ods.callbacks {
		other := ic.source()
		if other != nil && other == target {
			return []*objectDecodeState{target}
		} else if childList := other.findCycleFor(target); childList != nil {
			return append(childList, other)
		}
	}

	// This should not occur.
	Failf("no deadlock found?")
	panic("unreachable")
}

// findCycle finds a dependency cycle.
func (ods *objectDecodeState) findCycle() []*objectDecodeState {
	return append(ods.findCycleFor(ods), ods)
}

// source implements internalCallback.source.
func (ods *objectDecodeState) source() *objectDecodeState {
	return ods
}

// callbackRun implements internalCallback.callbackRun.
func (ods *objectDecodeState) callbackRun() {
	ods.blockedBy--
}

// decodeState is a graph of objects in the process of being decoded.
//
// The decode process involves loading the breadth-first graph generated by
// encode. This graph is read in it's entirety, ensuring that all object
// storage is complete.
//
// As the graph is being serialized, a set of completion callbacks are
// executed. These completion callbacks should form a set of acyclic subgraphs
// over the original one. After decoding is complete, the objects are scanned
// to ensure that all callbacks are executed, otherwise the callback graph was
// not acyclic.
type decodeState struct {
	// ctx is the decode context.
	ctx context.Context

	// r is the input stream.
	r wire.Reader

	// types is the type database.
	types typeDecodeDatabase

	// objectByID is the set of objects in progress.
	objectsByID []*objectDecodeState

	// deferred are objects that have been read, by no interest has been
	// registered yet. These will be decoded once interest in registered.
	deferred map[objectID]wire.Object

	// pending is the set of objects that are not yet complete.
	pending completeList

	// stats tracks time data.
	stats Stats
}

// lookup looks up an object in decodeState or returns nil if no such object
// has been previously registered.
func (ds *decodeState) lookup(id objectID) *objectDecodeState {
	if len(ds.objectsByID) < int(id) {
		return nil
	}
	return ds.objectsByID[id-1]
}

// checkComplete checks for completion.
func (ds *decodeState) checkComplete(ods *objectDecodeState) bool {
	// Still blocked?
	if ods.blockedBy > 0 {
		return false
	}

	// Track stats if relevant.
	if ods.callbacks != nil && ods.typ != 0 {
		ds.stats.start(ods.typ)
		defer ds.stats.done()
	}

	// Fire all callbacks.
	for _, ic := range ods.callbacks {
		ic.callbackRun()
	}

	// Mark completed.
	cbs := ods.callbacks
	ods.callbacks = nil
	ds.pending.Remove(ods)

	// Recursively check others.
	for _, ic := range cbs {
		if other := ic.source(); other != nil && other.blockedBy == 0 {
			ds.checkComplete(other)
		}
	}

	return true // All set.
}

// wait registers a dependency on an object.
//
// As a special case, we always allow _useable_ references back to the first
// decoding object because it may have fields that are already decoded. We also
// allow trivial self reference, since they can be handled internally.
func (ds *decodeState) wait(waiter *objectDecodeState, id objectID, callback func()) {
	switch id {
	case waiter.id:
		// Trivial self reference.
		fallthrough
	case 1:
		// Root object; see above.
		if callback != nil {
			callback()
		}
		return
	}

	// Mark as blocked.
	waiter.blockedBy++

	// No nil can be returned here.
	other := ds.lookup(id)
	if callback != nil {
		// Add the additional user callback.
		other.addCallback(userCallback(callback))
	}

	// Mark waiter as unblocked.
	other.addCallback(waiter)
}

// waitObject notes a blocking relationship.
func (ds *decodeState) waitObject(ods *objectDecodeState, encoded wire.Object, callback func()) {
	if rv, ok := encoded.(*wire.Ref); ok && rv.Root != 0 {
		// Refs can encode pointers and maps.
		ds.wait(ods, objectID(rv.Root), callback)
	} else if sv, ok := encoded.(*wire.Slice); ok && sv.Ref.Root != 0 {
		// See decodeObject; we need to wait for the array (if non-nil).
		ds.wait(ods, objectID(sv.Ref.Root), callback)
	} else if iv, ok := encoded.(*wire.Interface); ok {
		// It's an interface (wait recursively).
		ds.waitObject(ods, iv.Value, callback)
	} else if callback != nil {
		// Nothing to wait for: execute the callback immediately.
		callback()
	}
}

// walkChild returns a child object from obj, given an accessor path. This is
// the decode-side equivalent to traverse in encode.go.
//
// For the purposes of this function, a child object is either a field within a
// struct or an array element, with one such indirection per element in
// path. The returned value may be an unexported field, so it may not be
// directly assignable. See decode_unsafe.go.
func walkChild(path []wire.Dot, obj reflect.Value) reflect.Value {
	// See wire.Ref.Dots. The path here is specified in reverse order.
	for i := len(path) - 1; i >= 0; i-- {
		switch pc := path[i].(type) {
		case *wire.FieldName: // Must be a pointer.
			if obj.Kind() != reflect.Struct {
				Failf("next component in child path is a field name, but the current object is not a struct. Path: %v, current obj: %#v", path, obj)
			}
			obj = obj.FieldByName(string(*pc))
		case wire.Index: // Embedded.
			if obj.Kind() != reflect.Array {
				Failf("next component in child path is an array index, but the current object is not an array. Path: %v, current obj: %#v", path, obj)
			}
			obj = obj.Index(int(pc))
		default:
			panic("unreachable: switch should be exhaustive")
		}
	}
	return obj
}

// register registers a decode with a type.
//
// This type is only used to instantiate a new object if it has not been
// registered previously. This depends on the type provided if none is
// available in the object itself.
func (ds *decodeState) register(r *wire.Ref, typ reflect.Type) reflect.Value {
	// Grow the objectsByID slice.
	id := objectID(r.Root)
	if len(ds.objectsByID) < int(id) {
		ds.objectsByID = append(ds.objectsByID, make([]*objectDecodeState, int(id)-len(ds.objectsByID))...)
	}

	// Does this object already exist?
	ods := ds.objectsByID[id-1]
	if ods != nil {
		return walkChild(r.Dots, ods.obj)
	}

	// Create the object.
	if len(r.Dots) != 0 {
		typ = ds.findType(r.Type)
	}
	v := reflect.New(typ)
	ods = &objectDecodeState{
		id:  id,
		obj: v.Elem(),
	}
	ds.objectsByID[id-1] = ods
	ds.pending.PushBack(ods)

	// Process any deferred objects & callbacks.
	if encoded, ok := ds.deferred[id]; ok {
		delete(ds.deferred, id)
		ds.decodeObject(ods, ods.obj, encoded)
	}

	return walkChild(r.Dots, ods.obj)
}

// objectDecoder is for decoding structs.
type objectDecoder struct {
	// ds is decodeState.
	ds *decodeState

	// ods is current object being decoded.
	ods *objectDecodeState

	// reconciledTypeEntry is the reconciled type information.
	rte *reconciledTypeEntry

	// encoded is the encoded object state.
	encoded *wire.Struct
}

// load is helper for the public methods on Source.
func (od *objectDecoder) load(slot int, objPtr reflect.Value, wait bool, fn func()) {
	// Note that we have reconciled the type and may remap the fields here
	// to match what's expected by the decoder. The "slot" parameter here
	// is in terms of the local type, where the fields in the encoded
	// object are in terms of the wire object's type, which might be in a
	// different order (but will have the same fields).
	v := *od.encoded.Field(od.rte.FieldOrder[slot])
	od.ds.decodeObject(od.ods, objPtr.Elem(), v)
	if wait {
		// Mark this individual object a blocker.
		od.ds.waitObject(od.ods, v, fn)
	}
}

// aterLoad implements Source.AfterLoad.
func (od *objectDecoder) afterLoad(fn func()) {
	// Queue the local callback; this will execute when all of the above
	// data dependencies have been cleared.
	od.ods.addCallback(userCallback(fn))
}

// decodeStruct decodes a struct value.
func (ds *decodeState) decodeStruct(ods *objectDecodeState, obj reflect.Value, encoded *wire.Struct) {
	if encoded.TypeID == 0 {
		// Allow anonymous empty structs, but only if the encoded
		// object also has no fields.
		if encoded.Fields() == 0 && obj.NumField() == 0 {
			return
		}

		// Propagate an error.
		Failf("empty struct on wire %#v has field mismatch with type %q", encoded, obj.Type().Name())
	}

	// Lookup the object type.
	rte := ds.types.Lookup(typeID(encoded.TypeID), obj.Type())
	ods.typ = typeID(encoded.TypeID)

	// Invoke the loader.
	od := objectDecoder{
		ds:      ds,
		ods:     ods,
		rte:     rte,
		encoded: encoded,
	}
	ds.stats.start(ods.typ)
	defer ds.stats.done()
	if sl, ok := obj.Addr().Interface().(SaverLoader); ok {
		// Note: may be a registered empty struct which does not
		// implement the saver/loader interfaces.
		sl.StateLoad(Source{internal: od})
	}
}

// decodeMap decodes a map value.
func (ds *decodeState) decodeMap(ods *objectDecodeState, obj reflect.Value, encoded *wire.Map) {
	if obj.IsNil() {
		// See pointerTo.
		obj.Set(reflect.MakeMap(obj.Type()))
	}
	for i := 0; i < len(encoded.Keys); i++ {
		// Decode the objects.
		kv := reflect.New(obj.Type().Key()).Elem()
		vv := reflect.New(obj.Type().Elem()).Elem()
		ds.decodeObject(ods, kv, encoded.Keys[i])
		ds.decodeObject(ods, vv, encoded.Values[i])
		ds.waitObject(ods, encoded.Keys[i], nil)
		ds.waitObject(ods, encoded.Values[i], nil)

		// Set in the map.
		obj.SetMapIndex(kv, vv)
	}
}

// decodeArray decodes an array value.
func (ds *decodeState) decodeArray(ods *objectDecodeState, obj reflect.Value, encoded *wire.Array) {
	if len(encoded.Contents) != obj.Len() {
		Failf("mismatching array length expect=%d, actual=%d", obj.Len(), len(encoded.Contents))
	}
	// Decode the contents into the array.
	for i := 0; i < len(encoded.Contents); i++ {
		ds.decodeObject(ods, obj.Index(i), encoded.Contents[i])
		ds.waitObject(ods, encoded.Contents[i], nil)
	}
}

// findType finds the type for the given wire.TypeSpecs.
func (ds *decodeState) findType(t wire.TypeSpec) reflect.Type {
	switch x := t.(type) {
	case wire.TypeID:
		typ := ds.types.LookupType(typeID(x))
		rte := ds.types.Lookup(typeID(x), typ)
		return rte.LocalType
	case *wire.TypeSpecPointer:
		return reflect.PtrTo(ds.findType(x.Type))
	case *wire.TypeSpecArray:
		return reflect.ArrayOf(int(x.Count), ds.findType(x.Type))
	case *wire.TypeSpecSlice:
		return reflect.SliceOf(ds.findType(x.Type))
	case *wire.TypeSpecMap:
		return reflect.MapOf(ds.findType(x.Key), ds.findType(x.Value))
	default:
		// Should not happen.
		Failf("unknown type %#v", t)
	}
	panic("unreachable")
}

// decodeInterface decodes an interface value.
func (ds *decodeState) decodeInterface(ods *objectDecodeState, obj reflect.Value, encoded *wire.Interface) {
	if _, ok := encoded.Type.(wire.TypeSpecNil); ok {
		// Special case; the nil object. Just decode directly, which
		// will read nil from the wire (if encoded correctly).
		ds.decodeObject(ods, obj, encoded.Value)
		return
	}

	// We now need to resolve the actual type.
	typ := ds.findType(encoded.Type)

	// We need to imbue type information here, then we can proceed to
	// decode normally. In order to avoid issues with setting value-types,
	// we create a new non-interface version of this object. We will then
	// set the interface object to be equal to whatever we decode.
	origObj := obj
	obj = reflect.New(typ).Elem()
	defer origObj.Set(obj)

	// With the object now having sufficient type information to actually
	// have Set called on it, we can proceed to decode the value.
	ds.decodeObject(ods, obj, encoded.Value)
}

// isFloatEq determines if x and y represent the same value.
func isFloatEq(x float64, y float64) bool {
	switch {
	case math.IsNaN(x):
		return math.IsNaN(y)
	case math.IsInf(x, 1):
		return math.IsInf(y, 1)
	case math.IsInf(x, -1):
		return math.IsInf(y, -1)
	default:
		return x == y
	}
}

// isComplexEq determines if x and y represent the same value.
func isComplexEq(x complex128, y complex128) bool {
	return isFloatEq(real(x), real(y)) && isFloatEq(imag(x), imag(y))
}

// decodeObject decodes a object value.
func (ds *decodeState) decodeObject(ods *objectDecodeState, obj reflect.Value, encoded wire.Object) {
	switch x := encoded.(type) {
	case wire.Nil: // Fast path: first.
		// We leave obj alone here. That's because if obj represents an
		// interface, it may have been imbued with type information in
		// decodeInterface, and we don't want to destroy that.
	case *wire.Ref:
		// Nil pointers may be encoded in a "forceValue" context. For
		// those we just leave it alone as the value will already be
		// correct (nil).
		if id := objectID(x.Root); id == 0 {
			return
		}

		// Note that if this is a map type, we go through a level of
		// indirection to allow for map aliasing.
		if obj.Kind() == reflect.Map {
			v := ds.register(x, obj.Type())
			if v.IsNil() {
				// Note that we don't want to clobber the map
				// if has already been decoded by decodeMap. We
				// just make it so that we have a consistent
				// reference when that eventually does happen.
				v.Set(reflect.MakeMap(v.Type()))
			}
			obj.Set(v)
			return
		}

		// Normal assignment: authoritative only if no dots.
		v := ds.register(x, obj.Type().Elem())
		obj.Set(reflectValueRWAddr(v))
	case wire.Bool:
		obj.SetBool(bool(x))
	case wire.Int:
		obj.SetInt(int64(x))
		if obj.Int() != int64(x) {
			Failf("signed integer truncated from %v to %v", int64(x), obj.Int())
		}
	case wire.Uint:
		obj.SetUint(uint64(x))
		if obj.Uint() != uint64(x) {
			Failf("unsigned integer truncated from %v to %v", uint64(x), obj.Uint())
		}
	case wire.Float32:
		obj.SetFloat(float64(x))
	case wire.Float64:
		obj.SetFloat(float64(x))
		if !isFloatEq(obj.Float(), float64(x)) {
			Failf("floating point number truncated from %v to %v", float64(x), obj.Float())
		}
	case *wire.Complex64:
		obj.SetComplex(complex128(*x))
	case *wire.Complex128:
		obj.SetComplex(complex128(*x))
		if !isComplexEq(obj.Complex(), complex128(*x)) {
			Failf("complex number truncated from %v to %v", complex128(*x), obj.Complex())
		}
	case *wire.String:
		obj.SetString(string(*x))
	case *wire.Slice:
		// See *wire.Ref above; same applies.
		if id := objectID(x.Ref.Root); id == 0 {
			return
		}
		// Note that it's fine to slice the array here and assume that
		// contents will still be filled in later on.
		typ := reflect.ArrayOf(int(x.Capacity), obj.Type().Elem()) // The object type.
		v := ds.register(&x.Ref, typ)
		obj.Set(reflectValueRWSlice3(v, 0, int(x.Length), int(x.Capacity)))
	case *wire.Array:
		ds.decodeArray(ods, obj, x)
	case *wire.Struct:
		ds.decodeStruct(ods, obj, x)
	case *wire.Map:
		ds.decodeMap(ods, obj, x)
	case *wire.Interface:
		ds.decodeInterface(ods, obj, x)
	default:
		// Should not happen, not propagated as an error.
		Failf("unknown object %#v for %q", encoded, obj.Type().Name())
	}
}

// Load deserializes the object graph rooted at obj.
//
// This function may panic and should be run in safely().
func (ds *decodeState) Load(obj reflect.Value) {
	ds.stats.init()
	defer ds.stats.fini(func(id typeID) string {
		return ds.types.LookupName(id)
	})

	// Create the root object.
	rootOds := &objectDecodeState{
		id:  1,
		obj: obj,
	}
	ds.objectsByID = append(ds.objectsByID, rootOds)
	ds.pending.PushBack(rootOds)

	// Read the number of objects.
	numObjects, object, err := ReadHeader(ds.r)
	if err != nil {
		Failf("header error: %w", err)
	}
	if !object {
		Failf("object missing")
	}

	// Decode all objects.
	var (
		encoded wire.Object
		ods     *objectDecodeState
		id      objectID
		tid     = typeID(1)
	)
	if err := safely(func() {
		// Decode all objects in the stream.
		//
		// Note that the structure of this decoding loop should match the raw
		// decoding loop in state/pretty/pretty.printer.printStream().
		for i := uint64(0); i < numObjects; {
			// Unmarshal either a type object or object ID.
			encoded = wire.Load(ds.r)
			switch we := encoded.(type) {
			case *wire.Type:
				ds.types.Register(we)
				tid++
				encoded = nil
				continue
			case wire.Uint:
				id = objectID(we)
				i++
				// Unmarshal and resolve the actual object.
				encoded = wire.Load(ds.r)
				ods = ds.lookup(id)
				if ods != nil {
					// Decode the object.
					ds.decodeObject(ods, ods.obj, encoded)
				} else {
					// If an object hasn't had interest registered
					// previously or isn't yet valid, we deferred
					// decoding until interest is registered.
					ds.deferred[id] = encoded
				}
				// For error handling.
				ods = nil
				encoded = nil
			default:
				Failf("wanted type or object ID, got %T", encoded)
			}
		}
	}); err != nil {
		// Include as much information as we can, taking into account
		// the possible state transitions above.
		if ods != nil {
			Failf("error decoding object ID %d (%T) from %#v: %w", id, ods.obj.Interface(), encoded, err)
		} else if encoded != nil {
			Failf("error decoding from %#v: %w", encoded, err)
		} else {
			Failf("general decoding error: %w", err)
		}
	}

	// Check if we have any deferred objects.
	numDeferred := 0
	for id, encoded := range ds.deferred {
		numDeferred++
		if s, ok := encoded.(*wire.Struct); ok && s.TypeID != 0 {
			typ := ds.types.LookupType(typeID(s.TypeID))
			Failf("unused deferred object: ID %d, type %v", id, typ)
		} else {
			Failf("unused deferred object: ID %d, %#v", id, encoded)
		}
	}
	if numDeferred != 0 {
		Failf("still had %d deferred objects", numDeferred)
	}

	// Scan and fire all callbacks. We iterate over the list of incomplete
	// objects until all have been finished. We stop iterating if no
	// objects become complete (there is a dependency cycle).
	//
	// Note that we iterate backwards here, because there will be a strong
	// tendendcy for blocking relationships to go from earlier objects to
	// later (deeper) objects in the graph. This will reduce the number of
	// iterations required to finish all objects.
	if err := safely(func() {
		for ds.pending.Back() != nil {
			thisCycle := false
			for ods = ds.pending.Back(); ods != nil; {
				if ds.checkComplete(ods) {
					thisCycle = true
					break
				}
				ods = ods.Prev()
			}
			if !thisCycle {
				break
			}
		}
	}); err != nil {
		Failf("error executing callbacks for %#v: %w", ods.obj.Interface(), err)
	}

	// Check if we have any remaining dependency cycles. If there are any
	// objects left in the pending list, then it must be due to a cycle.
	if ods := ds.pending.Front(); ods != nil {
		// This must be the result of a dependency cycle.
		cycle := ods.findCycle()
		var buf bytes.Buffer
		buf.WriteString("dependency cycle: {")
		for i, cycleOS := range cycle {
			if i > 0 {
				buf.WriteString(" => ")
			}
			fmt.Fprintf(&buf, "%q", cycleOS.obj.Type())
		}
		buf.WriteString("}")
		Failf("incomplete graph: %s", string(buf.Bytes()))
	}
}

// ReadHeader reads an object header.
//
// Each object written to the statefile is prefixed with a header. See
// WriteHeader for more information; these functions are exported to allow
// non-state writes to the file to play nice with debugging tools.
func ReadHeader(r wire.Reader) (length uint64, object bool, err error) {
	// Read the header.
	err = safely(func() {
		length = wire.LoadUint(r)
	})
	if err != nil {
		// On the header, pass raw I/O errors.
		if sErr, ok := err.(*ErrState); ok {
			return 0, false, sErr.Unwrap()
		}
	}

	// Decode whether the object is valid.
	object = length&objectFlag != 0
	length &^= objectFlag
	return
}
