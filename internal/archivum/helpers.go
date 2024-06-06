package archivum

import (
	"io/fs"
	"os/user"
	"strconv"
)

func groupByGID(gid uint32) (*user.Group, error) {
	return user.LookupGroupId(strconv.FormatUint(uint64(gid), 10))
}

func userByUID(uid uint32) (*user.User, error) {
	return user.LookupId(strconv.FormatUint(uint64(uid), 10))
}

func permsOctal(fm fs.FileMode) string {
	return "0" + strconv.FormatUint(uint64(fm), 8)
}
