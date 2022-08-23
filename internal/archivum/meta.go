/*
* meta.go - Get UNIX metadata for files and return in a struct
*
 */

package archivum

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

type Metadata struct {
	FType    rune
	Path     string
	Major    string
	Minor    string
	OctalMod string
	Owner    string
	Group    string
}

var (
	errSLink = errors.New("symbolic links are unsupported")
)

func Scan(path string) (*Metadata, error) {
	fi, err := os.Lstat(path)
	if err != nil {
		return nil, err
	}

	ftype := determineFType(fi)
	fstat := fi.Sys().(*syscall.Stat_t)

	// minor/major numbers are only used in device files */
	var majorNumber, minorNumber string
	if (ftype == 'c') || (ftype == 'b') {
		majorNumber, minorNumber = "nil", "nil"
	}

	/* octalPermissions := */

	getOwner, _ := user.LookupId(strconv.FormatUint(uint64(fstat.Uid), 10))
	getGroup, _ := user.LookupGroupId(strconv.FormatUint(uint64(fstat.Gid), 10))
	ownerPermissions := getOwner.Username
	groupPermissions := getGroup.Name

	if ftype == 's' {
		return nil, fmt.Errorf("%w <%s>", errSLink, path)
	}

	Data := &Metadata{
		FType:    ftype,
		Path:     path,
		Major:    majorNumber,
		Minor:    minorNumber,
		OctalMod: "octalPermissions",
		Owner:    ownerPermissions,
		Group:    groupPermissions,
	}

	return Data, nil
}

func determineFType(fi os.FileInfo) rune {
	switch mode := fi.Mode(); {
	case mode.IsRegular():
		return 'f'
	case mode.IsDir():
		return 'd'
	case mode&fs.ModeSymlink != 0:
		return 's'
	case mode&fs.ModeDevice != 0:
		return 'b'
	case mode&fs.ModeCharDevice != 0:
		return 'c'
	case mode&fs.ModeNamedPipe != 0:
		return 'p'
	default:
		return '?'
	}
}
