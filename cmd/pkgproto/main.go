package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

const (
	errStat  = "unable to stat <%s>"
	errSLink = "symbolic links are not supported <%s>"
	exitStat = 0
)

func main() {
	readstdin := bufio.NewScanner(os.Stdin)
	for readstdin.Scan() {
		scan(readstdin.Text())
	}
}

func scan(path string) {
	fi, err := os.Lstat(path) // Lstat already does the work of checking if exists
	if os.IsNotExist(err) {
		log.Fatalf(errStat, path) // log.Fatalf already already formats error messages
		os.Exit(exitStat + 1)
	}

	ftype := determineFType(fi)
	fstat := fi.Sys().(*syscall.Stat_t)

	/* I do not have an idea of how we will be getting these for now...
	* majorNumber, minorNumber :=
	 */

	/* Later on this shall be moved to a library (libperms maybe?), since we
	* probably will be manipulating permissions in the future. */

	/* octalPermissions := */

	/* This is pathetic. We're converting our UID (and GID too) to uint64,
	* then converting it to a string using FormatUint; why Go doesn't do
	* this in a more on-the-fly way? */
	getOwner, _ := user.LookupId(strconv.FormatUint(uint64(fstat.Uid), 10))
	getGroup, _ := user.LookupGroupId(strconv.FormatUint(uint64(fstat.Gid), 10))
	ownerPermissions := getOwner.Username
	groupPermissions := getGroup.Name

	/* If it's a symbolic link, print a error saying that these aren't
	*  supported. I'm almost changing of idea in this subject, may we
	*  couldn't support links, but at least we could use the real file, eh? */
	if ftype == 's' {
		log.Fatalf(errSLink, path)
	}

	fmt.Printf("%c %s %s %s %s %s %s\n", ftype, path, "major", "minor", "octalPermissions", ownerPermissions, groupPermissions)
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
