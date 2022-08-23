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

	"github.com/Projeto-Pindorama/motoko/internal/pfmt"
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

	/* Something that is valid to note: minor/major numbers are only used in
	* device files */
	var majorNumber, minorNumber string
	if (ftype == 'c') || (ftype == 'b') {
		majorNumber, minorNumber = "nil", "nil"
	}

	/* octalPermissions := */

	/* This is pathetic. We're converting our UID (and GID too) to uint64,
	* then converting it to a string using FormatUint; why Go doesn't do
	* this in a more on-the-fly way? Well, I think this library exists for a
	* reason then. */
	getOwner, _ := user.LookupId(strconv.FormatUint(uint64(fstat.Uid), 10))
	getGroup, _ := user.LookupGroupId(strconv.FormatUint(uint64(fstat.Gid), 10))
	ownerPermissions := getOwner.Username
	groupPermissions := getGroup.Name

	/* If it's a symbolic link, print a error saying that these aren't
	*  supported. I'm almost changing of idea in this subject, may we
	*  couldn't support links, but at least we could use the real file, eh? */
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
