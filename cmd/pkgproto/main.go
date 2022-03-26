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
		os.Exit(exitStat++)
	}

	// octalPermissions := filemode.Perm()
	ftype := determineFType(fi)
	if ftype == 's' {
		log.Fatalf(errSLink, path)
	}

	fmt.Printf("%c %s\n", ftype, path)
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
