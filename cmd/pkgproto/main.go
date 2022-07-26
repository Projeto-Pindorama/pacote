package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"internal/archivum"
	"log"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func main() {
	readstdin := bufio.NewScanner(os.Stdin)
	for readstdin.Scan() {
		archivum.Scan(readstdin.Text())
		/* I don't know exactly how to print the struct that
		* archivum.Scan() returns now. I need to do this fast as
		* possible.
		* fmt.Printf("%c %s %s %s %s %s %s\n", )
		*/
	}
}
