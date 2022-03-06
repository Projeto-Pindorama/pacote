package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

const (
	err_stat = "unable to stat <%s>"
	err_slink = "symbolic links are not supported <%s>"
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
		errmsg := fmt.Sprintf(err_stat, path)
		log.Fatal(errmsg)
	}

	fi, err := os.Lstat(path)
	octal_permissions := fi.Mode().Perm()
	switch filetype := fi.Mode(); {
		case filetype.IsRegular():
			type := f
		case filetupe.IsDir():
			type := d
		case filetype&fs.ModeNamedPipe != 0:
			type := p
		case filetype&fs.ModeSymlink !=0:
			errmsg := fmt.Sprintf(err_slink, path)
			log.Fatal(errmsg)
	}
	fmt.Printf('%c %s', type, path)
}
