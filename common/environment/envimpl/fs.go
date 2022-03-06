package envimpl

import (
	"github.com/bmclab-git/v2ray-core/v5/common/environment"
	"github.com/bmclab-git/v2ray-core/v5/common/platform/filesystem"
	"github.com/bmclab-git/v2ray-core/v5/common/platform/filesystem/fsifce"
)

type fileSystemDefaultImpl struct{}

func (f fileSystemDefaultImpl) OpenFileForReadSeek() fsifce.FileSeekerFunc {
	return filesystem.NewFileSeeker
}

func (f fileSystemDefaultImpl) OpenFileForRead() fsifce.FileReaderFunc {
	return filesystem.NewFileReader
}

func (f fileSystemDefaultImpl) OpenFileForWrite() fsifce.FileWriterFunc {
	return filesystem.NewFileWriter
}

func NewDefaultFileSystemDefaultImpl() environment.FileSystemCapabilitySet {
	return fileSystemDefaultImpl{}
}
