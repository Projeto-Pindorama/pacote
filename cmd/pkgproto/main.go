package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Projeto-Pindorama/motoko/internal/archivum"
)

var class string
var ignoreLinks bool

func main() {
	flag.StringVar(&class, "c", "none", "Maps the class of all paths to class.")
	flag.BoolVar(&ignoreLinks, "i", false, "Ignores symbolic links and records the paths as ftype=f (a file) versus ftype=s (symbolic link).")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: pkgproto [-i] [-c class] [path ...]\n")
	}

	flag.Parse()

	remainingArgs := flag.Args()

	unixFS := archivum.NewUnixFS("/")

	for _, v := range remainingArgs {
		dir := strings.SplitN(v, "=", 2)

		filepath.WalkDir(dir[0], func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				log.Fatal(err)
				return err
			}

			metadata, err := archivum.Scan(unixFS, path)
			if err != nil {
				log.Fatal(err)
				return err
			}

			fmt.Println(strings.Replace(MetadataToString(metadata), dir[0], dir[1], 1))

			return nil
		})
	}

	if len(remainingArgs) > 0 {
		return
	}

	readstdin := bufio.NewScanner(os.Stdin)

	for readstdin.Scan() {
		metadata, err := archivum.Scan(unixFS, readstdin.Text())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(MetadataToString(metadata))
	}
}

func MetadataToString(m *archivum.Metadata) string {
	var ret string
	var fileType rune = m.FType

	if m.FType == 's' && ignoreLinks {
		fileType = 'f'
	}

	ret = fmt.Sprintf("%c %s %s", fileType, class, m.Path)

	if fileType == 's' {
		ret = ret + fmt.Sprintf("=%s", m.RealPath)
	} else if m.DeviceInfo == nil {
		ret = ret + fmt.Sprintf(" %s %s %s", m.OctalMod, m.Owner, m.Group)
	} else {
		ret = ret + fmt.Sprintf(" %d %d %s %s %s", m.DeviceInfo.Major, m.DeviceInfo.Minor, m.OctalMod, m.Owner, m.Group)
	}

	return ret
}
