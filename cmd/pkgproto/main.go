package main

import (
	"fmt"
	"bufio"
	"os"
	"path/filepath"
	"log"
)

const (
	err_stat = "unable to stat <%s>"
	err_link = "symbolic links are not supported <%s>"
)

func main() {
	readstdin := bufio.NewScanner(os.Stdin)
	for readstdin.Scan() {
		scan(readstdin.Text())
	}

}

func scan(path string) {
	/* Get absolute path for the file. */
	filepath,_ := filepath.Abs(path)
	
	/* Does it even exists? */
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		/* First, format the error message, then pass it to panic() */
		/* I've really needing to create an error-handling library */
		errmsg := fmt.Sprintf(err_stat, filepath)
		log.Fatal(errmsg)
	}

/*	fi, err := os.Lstat(filepath)
	octal_permissions := fi.Mode().Perm()
	switch type := fi.Mode(); {
		case type.IsRegular():

*/	
	fmt.Println(filepath)
}
