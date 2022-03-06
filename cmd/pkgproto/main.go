package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

const (
	errStat = "unable to stat <%s>"
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
		errMsg := fmt.Sprintf(errStat, path)
		log.Fatal(errMsg)
	}

	fi, err := os.Lstat(path)
	octalPermissions := fi.Mode().Perm()
	switch filetype := fi.Mode(); {
		case filetype.IsRegular():
			type = "f"
		case filetupe.IsDir():
			type = "d"
		case filetype&fs.ModeSymlink !=0:
			errMsg = fmt.Sprintf(errSLink, path)
			log.Fatal(errMsg)
	}
	fmt.Printf('%c %s', type, path)
}
