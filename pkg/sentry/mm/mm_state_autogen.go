// automatically generated by stateify.

package mm

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (a *aioManager) StateTypeName() string {
	return "pkg/sentry/mm.aioManager"
}

func (a *aioManager) StateFields() []string {
	return []string{
		"contexts",
	}
}

func (a *aioManager) beforeSave() {}

// +checklocksignore
func (a *aioManager) StateSave(stateSinkObject state.Sink) {
	a.beforeSave()
	stateSinkObject.Save(0, &a.contexts)
}

func (a *aioManager) afterLoad() {}

// +checklocksignore
func (a *aioManager) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &a.contexts)
}

func (i *ioResult) StateTypeName() string {
	return "pkg/sentry/mm.ioResult"
}

func (i *ioResult) StateFields() []string {
	return []string{
		"data",
		"ioEntry",
	}
}

func (i *ioResult) beforeSave() {}

// +checklocksignore
func (i *ioResult) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.data)
	stateSinkObject.Save(1, &i.ioEntry)
}

func (i *ioResult) afterLoad() {}

// +checklocksignore
func (i *ioResult) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.data)
	stateSourceObject.Load(1, &i.ioEntry)
}

func (ctx *AIOContext) StateTypeName() string {
	return "pkg/sentry/mm.AIOContext"
}

func (ctx *AIOContext) StateFields() []string {
	return []string{
		"results",
		"maxOutstanding",
		"outstanding",
	}
}

func (ctx *AIOContext) beforeSave() {}

// +checklocksignore
func (ctx *AIOContext) StateSave(stateSinkObject state.Sink) {
	ctx.beforeSave()
	if !state.IsZeroValue(&ctx.dead) {
		state.Failf("dead is %#v, expected zero", &ctx.dead)
	}
	stateSinkObject.Save(0, &ctx.results)
	stateSinkObject.Save(1, &ctx.maxOutstanding)
	stateSinkObject.Save(2, &ctx.outstanding)
}

// +checklocksignore
func (ctx *AIOContext) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &ctx.results)
	stateSourceObject.Load(1, &ctx.maxOutstanding)
	stateSourceObject.Load(2, &ctx.outstanding)
	stateSourceObject.AfterLoad(ctx.afterLoad)
}

func (m *aioMappable) StateTypeName() string {
	return "pkg/sentry/mm.aioMappable"
}

func (m *aioMappable) StateFields() []string {
	return []string{
		"aioMappableRefs",
		"mfp",
		"fr",
	}
}

func (m *aioMappable) beforeSave() {}

// +checklocksignore
func (m *aioMappable) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.aioMappableRefs)
	stateSinkObject.Save(1, &m.mfp)
	stateSinkObject.Save(2, &m.fr)
}

func (m *aioMappable) afterLoad() {}

// +checklocksignore
func (m *aioMappable) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.aioMappableRefs)
	stateSourceObject.Load(1, &m.mfp)
	stateSourceObject.Load(2, &m.fr)
}

func (r *aioMappableRefs) StateTypeName() string {
	return "pkg/sentry/mm.aioMappableRefs"
}

func (r *aioMappableRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *aioMappableRefs) beforeSave() {}

// +checklocksignore
func (r *aioMappableRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *aioMappableRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(r.afterLoad)
}

func (l *ioList) StateTypeName() string {
	return "pkg/sentry/mm.ioList"
}

func (l *ioList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *ioList) beforeSave() {}

// +checklocksignore
func (l *ioList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *ioList) afterLoad() {}

// +checklocksignore
func (l *ioList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *ioEntry) StateTypeName() string {
	return "pkg/sentry/mm.ioEntry"
}

func (e *ioEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *ioEntry) beforeSave() {}

// +checklocksignore
func (e *ioEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *ioEntry) afterLoad() {}

// +checklocksignore
func (e *ioEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func (mm *MemoryManager) StateTypeName() string {
	return "pkg/sentry/mm.MemoryManager"
}

func (mm *MemoryManager) StateFields() []string {
	return []string{
		"p",
		"mfp",
		"layout",
		"users",
		"vmas",
		"brk",
		"usageAS",
		"lockedAS",
		"dataAS",
		"defMLockMode",
		"pmas",
		"curRSS",
		"maxRSS",
		"dumpability",
		"argv",
		"envv",
		"auxv",
		"executable",
		"aioManager",
		"sleepForActivation",
		"vdsoSigReturnAddr",
		"membarrierPrivateEnabled",
		"membarrierRSeqEnabled",
	}
}

// +checklocksignore
func (mm *MemoryManager) StateSave(stateSinkObject state.Sink) {
	mm.beforeSave()
	if !state.IsZeroValue(&mm.active) {
		state.Failf("active is %#v, expected zero", &mm.active)
	}
	if !state.IsZeroValue(&mm.captureInvalidations) {
		state.Failf("captureInvalidations is %#v, expected zero", &mm.captureInvalidations)
	}
	stateSinkObject.Save(0, &mm.p)
	stateSinkObject.Save(1, &mm.mfp)
	stateSinkObject.Save(2, &mm.layout)
	stateSinkObject.Save(3, &mm.users)
	stateSinkObject.Save(4, &mm.vmas)
	stateSinkObject.Save(5, &mm.brk)
	stateSinkObject.Save(6, &mm.usageAS)
	stateSinkObject.Save(7, &mm.lockedAS)
	stateSinkObject.Save(8, &mm.dataAS)
	stateSinkObject.Save(9, &mm.defMLockMode)
	stateSinkObject.Save(10, &mm.pmas)
	stateSinkObject.Save(11, &mm.curRSS)
	stateSinkObject.Save(12, &mm.maxRSS)
	stateSinkObject.Save(13, &mm.dumpability)
	stateSinkObject.Save(14, &mm.argv)
	stateSinkObject.Save(15, &mm.envv)
	stateSinkObject.Save(16, &mm.auxv)
	stateSinkObject.Save(17, &mm.executable)
	stateSinkObject.Save(18, &mm.aioManager)
	stateSinkObject.Save(19, &mm.sleepForActivation)
	stateSinkObject.Save(20, &mm.vdsoSigReturnAddr)
	stateSinkObject.Save(21, &mm.membarrierPrivateEnabled)
	stateSinkObject.Save(22, &mm.membarrierRSeqEnabled)
}

// +checklocksignore
func (mm *MemoryManager) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &mm.p)
	stateSourceObject.Load(1, &mm.mfp)
	stateSourceObject.Load(2, &mm.layout)
	stateSourceObject.Load(3, &mm.users)
	stateSourceObject.Load(4, &mm.vmas)
	stateSourceObject.Load(5, &mm.brk)
	stateSourceObject.Load(6, &mm.usageAS)
	stateSourceObject.Load(7, &mm.lockedAS)
	stateSourceObject.Load(8, &mm.dataAS)
	stateSourceObject.Load(9, &mm.defMLockMode)
	stateSourceObject.Load(10, &mm.pmas)
	stateSourceObject.Load(11, &mm.curRSS)
	stateSourceObject.Load(12, &mm.maxRSS)
	stateSourceObject.Load(13, &mm.dumpability)
	stateSourceObject.Load(14, &mm.argv)
	stateSourceObject.Load(15, &mm.envv)
	stateSourceObject.Load(16, &mm.auxv)
	stateSourceObject.Load(17, &mm.executable)
	stateSourceObject.Load(18, &mm.aioManager)
	stateSourceObject.Load(19, &mm.sleepForActivation)
	stateSourceObject.Load(20, &mm.vdsoSigReturnAddr)
	stateSourceObject.Load(21, &mm.membarrierPrivateEnabled)
	stateSourceObject.Load(22, &mm.membarrierRSeqEnabled)
	stateSourceObject.AfterLoad(mm.afterLoad)
}

func (v *vma) StateTypeName() string {
	return "pkg/sentry/mm.vma"
}

func (v *vma) StateFields() []string {
	return []string{
		"mappable",
		"off",
		"realPerms",
		"dontfork",
		"mlockMode",
		"numaPolicy",
		"numaNodemask",
		"id",
		"hint",
		"lastFault",
	}
}

func (v *vma) beforeSave() {}

// +checklocksignore
func (v *vma) StateSave(stateSinkObject state.Sink) {
	v.beforeSave()
	var realPermsValue int
	realPermsValue = v.saveRealPerms()
	stateSinkObject.SaveValue(2, realPermsValue)
	stateSinkObject.Save(0, &v.mappable)
	stateSinkObject.Save(1, &v.off)
	stateSinkObject.Save(3, &v.dontfork)
	stateSinkObject.Save(4, &v.mlockMode)
	stateSinkObject.Save(5, &v.numaPolicy)
	stateSinkObject.Save(6, &v.numaNodemask)
	stateSinkObject.Save(7, &v.id)
	stateSinkObject.Save(8, &v.hint)
	stateSinkObject.Save(9, &v.lastFault)
}

func (v *vma) afterLoad() {}

// +checklocksignore
func (v *vma) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &v.mappable)
	stateSourceObject.Load(1, &v.off)
	stateSourceObject.Load(3, &v.dontfork)
	stateSourceObject.Load(4, &v.mlockMode)
	stateSourceObject.Load(5, &v.numaPolicy)
	stateSourceObject.Load(6, &v.numaNodemask)
	stateSourceObject.Load(7, &v.id)
	stateSourceObject.Load(8, &v.hint)
	stateSourceObject.Load(9, &v.lastFault)
	stateSourceObject.LoadValue(2, new(int), func(y any) { v.loadRealPerms(y.(int)) })
}

func (p *pma) StateTypeName() string {
	return "pkg/sentry/mm.pma"
}

func (p *pma) StateFields() []string {
	return []string{
		"off",
		"translatePerms",
		"effectivePerms",
		"maxPerms",
		"needCOW",
		"private",
	}
}

func (p *pma) beforeSave() {}

// +checklocksignore
func (p *pma) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.off)
	stateSinkObject.Save(1, &p.translatePerms)
	stateSinkObject.Save(2, &p.effectivePerms)
	stateSinkObject.Save(3, &p.maxPerms)
	stateSinkObject.Save(4, &p.needCOW)
	stateSinkObject.Save(5, &p.private)
}

func (p *pma) afterLoad() {}

// +checklocksignore
func (p *pma) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.off)
	stateSourceObject.Load(1, &p.translatePerms)
	stateSourceObject.Load(2, &p.effectivePerms)
	stateSourceObject.Load(3, &p.maxPerms)
	stateSourceObject.Load(4, &p.needCOW)
	stateSourceObject.Load(5, &p.private)
}

func (s *pmaSet) StateTypeName() string {
	return "pkg/sentry/mm.pmaSet"
}

func (s *pmaSet) StateFields() []string {
	return []string{
		"root",
	}
}

func (s *pmaSet) beforeSave() {}

// +checklocksignore
func (s *pmaSet) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var rootValue *pmaSegmentDataSlices
	rootValue = s.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (s *pmaSet) afterLoad() {}

// +checklocksignore
func (s *pmaSet) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new(*pmaSegmentDataSlices), func(y any) { s.loadRoot(y.(*pmaSegmentDataSlices)) })
}

func (n *pmanode) StateTypeName() string {
	return "pkg/sentry/mm.pmanode"
}

func (n *pmanode) StateFields() []string {
	return []string{
		"nrSegments",
		"parent",
		"parentIndex",
		"hasChildren",
		"maxGap",
		"keys",
		"values",
		"children",
	}
}

func (n *pmanode) beforeSave() {}

// +checklocksignore
func (n *pmanode) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.nrSegments)
	stateSinkObject.Save(1, &n.parent)
	stateSinkObject.Save(2, &n.parentIndex)
	stateSinkObject.Save(3, &n.hasChildren)
	stateSinkObject.Save(4, &n.maxGap)
	stateSinkObject.Save(5, &n.keys)
	stateSinkObject.Save(6, &n.values)
	stateSinkObject.Save(7, &n.children)
}

func (n *pmanode) afterLoad() {}

// +checklocksignore
func (n *pmanode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.nrSegments)
	stateSourceObject.Load(1, &n.parent)
	stateSourceObject.Load(2, &n.parentIndex)
	stateSourceObject.Load(3, &n.hasChildren)
	stateSourceObject.Load(4, &n.maxGap)
	stateSourceObject.Load(5, &n.keys)
	stateSourceObject.Load(6, &n.values)
	stateSourceObject.Load(7, &n.children)
}

func (p *pmaSegmentDataSlices) StateTypeName() string {
	return "pkg/sentry/mm.pmaSegmentDataSlices"
}

func (p *pmaSegmentDataSlices) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Values",
	}
}

func (p *pmaSegmentDataSlices) beforeSave() {}

// +checklocksignore
func (p *pmaSegmentDataSlices) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.Start)
	stateSinkObject.Save(1, &p.End)
	stateSinkObject.Save(2, &p.Values)
}

func (p *pmaSegmentDataSlices) afterLoad() {}

// +checklocksignore
func (p *pmaSegmentDataSlices) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.Start)
	stateSourceObject.Load(1, &p.End)
	stateSourceObject.Load(2, &p.Values)
}

func (m *SpecialMappable) StateTypeName() string {
	return "pkg/sentry/mm.SpecialMappable"
}

func (m *SpecialMappable) StateFields() []string {
	return []string{
		"SpecialMappableRefs",
		"mfp",
		"fr",
		"name",
	}
}

func (m *SpecialMappable) beforeSave() {}

// +checklocksignore
func (m *SpecialMappable) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.SpecialMappableRefs)
	stateSinkObject.Save(1, &m.mfp)
	stateSinkObject.Save(2, &m.fr)
	stateSinkObject.Save(3, &m.name)
}

func (m *SpecialMappable) afterLoad() {}

// +checklocksignore
func (m *SpecialMappable) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.SpecialMappableRefs)
	stateSourceObject.Load(1, &m.mfp)
	stateSourceObject.Load(2, &m.fr)
	stateSourceObject.Load(3, &m.name)
}

func (r *SpecialMappableRefs) StateTypeName() string {
	return "pkg/sentry/mm.SpecialMappableRefs"
}

func (r *SpecialMappableRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *SpecialMappableRefs) beforeSave() {}

// +checklocksignore
func (r *SpecialMappableRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *SpecialMappableRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(r.afterLoad)
}

func (s *vmaSet) StateTypeName() string {
	return "pkg/sentry/mm.vmaSet"
}

func (s *vmaSet) StateFields() []string {
	return []string{
		"root",
	}
}

func (s *vmaSet) beforeSave() {}

// +checklocksignore
func (s *vmaSet) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var rootValue *vmaSegmentDataSlices
	rootValue = s.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (s *vmaSet) afterLoad() {}

// +checklocksignore
func (s *vmaSet) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new(*vmaSegmentDataSlices), func(y any) { s.loadRoot(y.(*vmaSegmentDataSlices)) })
}

func (n *vmanode) StateTypeName() string {
	return "pkg/sentry/mm.vmanode"
}

func (n *vmanode) StateFields() []string {
	return []string{
		"nrSegments",
		"parent",
		"parentIndex",
		"hasChildren",
		"maxGap",
		"keys",
		"values",
		"children",
	}
}

func (n *vmanode) beforeSave() {}

// +checklocksignore
func (n *vmanode) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.nrSegments)
	stateSinkObject.Save(1, &n.parent)
	stateSinkObject.Save(2, &n.parentIndex)
	stateSinkObject.Save(3, &n.hasChildren)
	stateSinkObject.Save(4, &n.maxGap)
	stateSinkObject.Save(5, &n.keys)
	stateSinkObject.Save(6, &n.values)
	stateSinkObject.Save(7, &n.children)
}

func (n *vmanode) afterLoad() {}

// +checklocksignore
func (n *vmanode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.nrSegments)
	stateSourceObject.Load(1, &n.parent)
	stateSourceObject.Load(2, &n.parentIndex)
	stateSourceObject.Load(3, &n.hasChildren)
	stateSourceObject.Load(4, &n.maxGap)
	stateSourceObject.Load(5, &n.keys)
	stateSourceObject.Load(6, &n.values)
	stateSourceObject.Load(7, &n.children)
}

func (v *vmaSegmentDataSlices) StateTypeName() string {
	return "pkg/sentry/mm.vmaSegmentDataSlices"
}

func (v *vmaSegmentDataSlices) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Values",
	}
}

func (v *vmaSegmentDataSlices) beforeSave() {}

// +checklocksignore
func (v *vmaSegmentDataSlices) StateSave(stateSinkObject state.Sink) {
	v.beforeSave()
	stateSinkObject.Save(0, &v.Start)
	stateSinkObject.Save(1, &v.End)
	stateSinkObject.Save(2, &v.Values)
}

func (v *vmaSegmentDataSlices) afterLoad() {}

// +checklocksignore
func (v *vmaSegmentDataSlices) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &v.Start)
	stateSourceObject.Load(1, &v.End)
	stateSourceObject.Load(2, &v.Values)
}

func init() {
	state.Register((*aioManager)(nil))
	state.Register((*ioResult)(nil))
	state.Register((*AIOContext)(nil))
	state.Register((*aioMappable)(nil))
	state.Register((*aioMappableRefs)(nil))
	state.Register((*ioList)(nil))
	state.Register((*ioEntry)(nil))
	state.Register((*MemoryManager)(nil))
	state.Register((*vma)(nil))
	state.Register((*pma)(nil))
	state.Register((*pmaSet)(nil))
	state.Register((*pmanode)(nil))
	state.Register((*pmaSegmentDataSlices)(nil))
	state.Register((*SpecialMappable)(nil))
	state.Register((*SpecialMappableRefs)(nil))
	state.Register((*vmaSet)(nil))
	state.Register((*vmanode)(nil))
	state.Register((*vmaSegmentDataSlices)(nil))
}
