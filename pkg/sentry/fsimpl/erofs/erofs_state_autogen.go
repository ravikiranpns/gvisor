// automatically generated by stateify.

package erofs

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (r *dentryRefs) StateTypeName() string {
	return "pkg/sentry/fsimpl/erofs.dentryRefs"
}

func (r *dentryRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *dentryRefs) beforeSave() {}

// +checklocksignore
func (r *dentryRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *dentryRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(r.afterLoad)
}

func (fd *directoryFD) StateTypeName() string {
	return "pkg/sentry/fsimpl/erofs.directoryFD"
}

func (fd *directoryFD) StateFields() []string {
	return []string{
		"fileDescription",
		"DirectoryFileDescriptionDefaultImpl",
		"off",
	}
}

func (fd *directoryFD) beforeSave() {}

// +checklocksignore
func (fd *directoryFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.fileDescription)
	stateSinkObject.Save(1, &fd.DirectoryFileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &fd.off)
}

func (fd *directoryFD) afterLoad() {}

// +checklocksignore
func (fd *directoryFD) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.fileDescription)
	stateSourceObject.Load(1, &fd.DirectoryFileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &fd.off)
}

func (fstype *FilesystemType) StateTypeName() string {
	return "pkg/sentry/fsimpl/erofs.FilesystemType"
}

func (fstype *FilesystemType) StateFields() []string {
	return []string{}
}

func (fstype *FilesystemType) beforeSave() {}

// +checklocksignore
func (fstype *FilesystemType) StateSave(stateSinkObject state.Sink) {
	fstype.beforeSave()
}

func (fstype *FilesystemType) afterLoad() {}

// +checklocksignore
func (fstype *FilesystemType) StateLoad(stateSourceObject state.Source) {
}

func (fs *filesystem) StateTypeName() string {
	return "pkg/sentry/fsimpl/erofs.filesystem"
}

func (fs *filesystem) StateFields() []string {
	return []string{
		"vfsfs",
		"mopts",
		"iopts",
		"devMinor",
		"root",
		"image",
		"mf",
		"inodeBuckets",
	}
}

func (fs *filesystem) beforeSave() {}

// +checklocksignore
func (fs *filesystem) StateSave(stateSinkObject state.Sink) {
	fs.beforeSave()
	stateSinkObject.Save(0, &fs.vfsfs)
	stateSinkObject.Save(1, &fs.mopts)
	stateSinkObject.Save(2, &fs.iopts)
	stateSinkObject.Save(3, &fs.devMinor)
	stateSinkObject.Save(4, &fs.root)
	stateSinkObject.Save(5, &fs.image)
	stateSinkObject.Save(6, &fs.mf)
	stateSinkObject.Save(7, &fs.inodeBuckets)
}

func (fs *filesystem) afterLoad() {}

// +checklocksignore
func (fs *filesystem) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fs.vfsfs)
	stateSourceObject.Load(1, &fs.mopts)
	stateSourceObject.Load(2, &fs.iopts)
	stateSourceObject.Load(3, &fs.devMinor)
	stateSourceObject.Load(4, &fs.root)
	stateSourceObject.Load(5, &fs.image)
	stateSourceObject.Load(6, &fs.mf)
	stateSourceObject.Load(7, &fs.inodeBuckets)
}

func (i *InternalFilesystemOptions) StateTypeName() string {
	return "pkg/sentry/fsimpl/erofs.InternalFilesystemOptions"
}

func (i *InternalFilesystemOptions) StateFields() []string {
	return []string{
		"UniqueID",
	}
}

func (i *InternalFilesystemOptions) beforeSave() {}

// +checklocksignore
func (i *InternalFilesystemOptions) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.UniqueID)
}

func (i *InternalFilesystemOptions) afterLoad() {}

// +checklocksignore
func (i *InternalFilesystemOptions) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.UniqueID)
}

func (ib *inodeBucket) StateTypeName() string {
	return "pkg/sentry/fsimpl/erofs.inodeBucket"
}

func (ib *inodeBucket) StateFields() []string {
	return []string{
		"inodeMap",
	}
}

func (ib *inodeBucket) beforeSave() {}

// +checklocksignore
func (ib *inodeBucket) StateSave(stateSinkObject state.Sink) {
	ib.beforeSave()
	stateSinkObject.Save(0, &ib.inodeMap)
}

func (ib *inodeBucket) afterLoad() {}

// +checklocksignore
func (ib *inodeBucket) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &ib.inodeMap)
}

func (i *inode) StateTypeName() string {
	return "pkg/sentry/fsimpl/erofs.inode"
}

func (i *inode) StateFields() []string {
	return []string{
		"Inode",
		"inodeRefs",
		"fs",
		"mappings",
		"locks",
		"watches",
	}
}

func (i *inode) beforeSave() {}

// +checklocksignore
func (i *inode) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.Inode)
	stateSinkObject.Save(1, &i.inodeRefs)
	stateSinkObject.Save(2, &i.fs)
	stateSinkObject.Save(3, &i.mappings)
	stateSinkObject.Save(4, &i.locks)
	stateSinkObject.Save(5, &i.watches)
}

func (i *inode) afterLoad() {}

// +checklocksignore
func (i *inode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.Inode)
	stateSourceObject.Load(1, &i.inodeRefs)
	stateSourceObject.Load(2, &i.fs)
	stateSourceObject.Load(3, &i.mappings)
	stateSourceObject.Load(4, &i.locks)
	stateSourceObject.Load(5, &i.watches)
}

func (d *dentry) StateTypeName() string {
	return "pkg/sentry/fsimpl/erofs.dentry"
}

func (d *dentry) StateFields() []string {
	return []string{
		"vfsd",
		"dentryRefs",
		"parent",
		"name",
		"inode",
		"childMap",
	}
}

func (d *dentry) beforeSave() {}

// +checklocksignore
func (d *dentry) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	var parentValue *dentry
	parentValue = d.saveParent()
	stateSinkObject.SaveValue(2, parentValue)
	stateSinkObject.Save(0, &d.vfsd)
	stateSinkObject.Save(1, &d.dentryRefs)
	stateSinkObject.Save(3, &d.name)
	stateSinkObject.Save(4, &d.inode)
	stateSinkObject.Save(5, &d.childMap)
}

func (d *dentry) afterLoad() {}

// +checklocksignore
func (d *dentry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.vfsd)
	stateSourceObject.Load(1, &d.dentryRefs)
	stateSourceObject.Load(3, &d.name)
	stateSourceObject.Load(4, &d.inode)
	stateSourceObject.Load(5, &d.childMap)
	stateSourceObject.LoadValue(2, new(*dentry), func(y any) { d.loadParent(y.(*dentry)) })
}

func (fd *fileDescription) StateTypeName() string {
	return "pkg/sentry/fsimpl/erofs.fileDescription"
}

func (fd *fileDescription) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"LockFD",
	}
}

func (fd *fileDescription) beforeSave() {}

// +checklocksignore
func (fd *fileDescription) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.vfsfd)
	stateSinkObject.Save(1, &fd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &fd.LockFD)
}

func (fd *fileDescription) afterLoad() {}

// +checklocksignore
func (fd *fileDescription) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.vfsfd)
	stateSourceObject.Load(1, &fd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &fd.LockFD)
}

func (r *inodeRefs) StateTypeName() string {
	return "pkg/sentry/fsimpl/erofs.inodeRefs"
}

func (r *inodeRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *inodeRefs) beforeSave() {}

// +checklocksignore
func (r *inodeRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *inodeRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(r.afterLoad)
}

func (fd *regularFileFD) StateTypeName() string {
	return "pkg/sentry/fsimpl/erofs.regularFileFD"
}

func (fd *regularFileFD) StateFields() []string {
	return []string{
		"fileDescription",
		"off",
	}
}

func (fd *regularFileFD) beforeSave() {}

// +checklocksignore
func (fd *regularFileFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.fileDescription)
	stateSinkObject.Save(1, &fd.off)
}

func (fd *regularFileFD) afterLoad() {}

// +checklocksignore
func (fd *regularFileFD) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.fileDescription)
	stateSourceObject.Load(1, &fd.off)
}

func (mf *imageMemmapFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/erofs.imageMemmapFile"
}

func (mf *imageMemmapFile) StateFields() []string {
	return []string{
		"image",
	}
}

func (mf *imageMemmapFile) beforeSave() {}

// +checklocksignore
func (mf *imageMemmapFile) StateSave(stateSinkObject state.Sink) {
	mf.beforeSave()
	stateSinkObject.Save(0, &mf.image)
}

func (mf *imageMemmapFile) afterLoad() {}

// +checklocksignore
func (mf *imageMemmapFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &mf.image)
}

func init() {
	state.Register((*dentryRefs)(nil))
	state.Register((*directoryFD)(nil))
	state.Register((*FilesystemType)(nil))
	state.Register((*filesystem)(nil))
	state.Register((*InternalFilesystemOptions)(nil))
	state.Register((*inodeBucket)(nil))
	state.Register((*inode)(nil))
	state.Register((*dentry)(nil))
	state.Register((*fileDescription)(nil))
	state.Register((*inodeRefs)(nil))
	state.Register((*regularFileFD)(nil))
	state.Register((*imageMemmapFile)(nil))
}
