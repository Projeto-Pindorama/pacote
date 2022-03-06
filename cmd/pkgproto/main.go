package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
)

const (
	errStat  = "unable to stat <%s>"
	errSLink = "symbolic links are not supported <%s>"
)

func main() {
	readstdin := bufio.NewScanner(os.Stdin)
	for readstdin.Scan() {
		scan(readstdin.Text())
	}
}

func scan(path string) {
	/* Does it even exists? */
	fi, err := os.Lstat(path) // Lstat already does the work
	if os.IsNotExist(err) {
		/* First, format the error message, then pass it to panic() */
		/* I've really needing to create an error-handling library */
		log.Fatalf(errStat, path)
	}

	// octalPermissions := filemode.Perm()
	ftype := determineFType(fi)
	if ftype == 's' {
		log.Fatalf(errSLink, path)
	}

	fmt.Printf("%c %s", ftype, path)
}

func determineFType(fi os.FileInfo) rune {
	switch mode := fi.Mode(); {
	case mode.IsRegular():
		return 'f'
	case mode.IsDir():
		return 'd'
	case mode&fs.ModeSymlink != 0:
		return 's'
	default:
		return '?'
	}
}
