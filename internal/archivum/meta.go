/* 
* meta.go - Get UNIX metadata for files and return in a struct 
*
*/

package archivum 

import (
	"bufio"
	"fmt"
	"io/fs"
	"internal/pfmt"
	"log"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

type Metadata struct {
	ftype rune
	path string
	major string /* I think these three will be int in the end */
	minor string
	octalmod string
	owner string
	group string
}

const (
	errStat = "unable to stat <%s>"
	errSLink = "symbolic links are not supported <%s>"
)

func Scan(path string) Metadata {
	fi, err := os.Lstat(path)
	if os.IsNotExist(err) {
		pfmt.pfmt(stderr, MM_ERROR, errStat, path)
	}

	ftype := determineFType(fi)
	fstat := fi.Sys().(*syscall.Stat_t)
	/* majorNumber, minorNumber := */
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
		pfmt.pfmt(stderr, MM_ERROR, errSLink, path)
	}

	Data := Metadata {
		ftype: ftype,
		path: path,
		major: "major",
		minor: "minor",
		octalmod: "octalPermissions",
		owner: ownerPermissions,
		group: groupPermissions,
	}
	return Data
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
