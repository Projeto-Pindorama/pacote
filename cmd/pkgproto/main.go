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
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		/* First, format the error message, then pass it to panic() */
		/* I've really needing to create an error-handling library */
		log.Fatalf(errStat, path)
	}

	fi, err := os.Lstat(path)
	filemode := fi.Mode()
	// octalPermissions := filemode.Perm()

	var ftype rune
	switch {
	case filemode.IsRegular():
		ftype = 'f'
	case filemode.IsDir():
		ftype = 'd'
	case filemode&fs.ModeSymlink != 0:
		log.Fatalf(errSLink, path)
	}
	fmt.Printf("%c %s", ftype, path)
}
