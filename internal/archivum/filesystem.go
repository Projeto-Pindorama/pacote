package archivum

import (
	"io/fs"
	"os/user"
)

type OwnershipFS interface {
	fs.FS
	StatOwned(name string) OwnedFileInfo // avoid conflicts when embedding StatFS
}

type OwnedFileInfo interface {
	fs.FileInfo
	Group() *user.Group
	User() *user.User
}