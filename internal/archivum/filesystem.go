package archivum

import (
	"io/fs"
	"os"
	"os/user"
	"syscall"
)

type OperatingSystemFS interface {
	fs.FS

	Stat(name string) (OSFileInfo, error)
}

type OSFileInfo interface {
	fs.FileInfo

	Group() *user.Group
	Owner() *user.User
	DeviceInfo() *DeviceInfo
}

type UnixFS struct {
	fs.FS
}

type UnixFileInfo struct {
	fs.FileInfo
	group *user.Group
	owner *user.User
}

func NewUnixFS(path string) UnixFS {
	return UnixFS{os.DirFS(path)}
}

func (fs UnixFS) Stat(name string) (OSFileInfo, error) {
	fileInfo, err := os.Lstat(name)
	if err != nil {
		return nil, err
	}

	sysInfo := fileInfo.Sys().(*syscall.Stat_t)

	owner, err := userByUID(sysInfo.Uid)
	if err != nil {
		return nil, err
	}

	group, err := groupByGID(sysInfo.Gid)
	if err != nil {
		return nil, err
	}

	unixFileInfo := UnixFileInfo{
		FileInfo: fileInfo,
		group:    group,
		owner:    owner,
	}

	return unixFileInfo, nil
}

func (fi UnixFileInfo) Owner() *user.User {
	return fi.owner
}

func (fi UnixFileInfo) Group() *user.Group {
	return fi.group
}

func (fi UnixFileInfo) DeviceInfo() *DeviceInfo {
	sysInfo := fi.Sys().(*syscall.Stat_t)
	return deviceInfoByDev(sysInfo.Rdev)
}
