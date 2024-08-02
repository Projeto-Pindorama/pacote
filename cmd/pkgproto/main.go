/*
 * pkgproto - generate a prototype file
 */
/*
 * Copyright (C) 2022 - 2024: Luiz Antônio Rangel (takusuman)
 * 				Samuel Brederodes (callsamu)
 *				João Pedro Vieira (JoaoP-Vieira)
 *
 *
 * SPDX-Licence-Identifier: NCSA 
 */

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

	"github.com/Projeto-Pindorama/pacote/internal/archivum"
)

var progname string
var class string
var fIgnoreLinks bool

func main() {
	progname = os.Args[0]
	var remainingArgs []string
	var unixFS archivum.UnixFS

	/* Argument parser. */
	flag.StringVar(&class, "c", "none",
		"Maps the class of all paths to class.")
	flag.BoolVar(&fIgnoreLinks, "i", false, `Ignores symbolic links and records
the paths as ftype=f (a file) versus
ftype=s (symbolic link).`)
	flag.Usage = usage
	flag.Parse()
	remainingArgs = flag.Args()

	unixFS = archivum.NewUnixFS("/")
	for _, v := range remainingArgs {
		dir := strings.SplitN(v, "=", 2)

		filepath.WalkDir(dir[0],
			func(path string, d fs.DirEntry, err error) error {
			var metadata *archivum.Metadata

			if err != nil {
				log.Fatal(err)
				return err
			}

			metadata, err = archivum.Scan(unixFS, path)
			if err != nil {
				log.Fatal(err)
				return err
			}

			fmt.Println(strings.Replace(MetadataToString(metadata),
								dir[0], dir[1], 1))

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
	var fileType rune
	fileType = m.FType

	if fileType == 's' && fIgnoreLinks {
		fileType = 'f'
	}

	ret = fmt.Sprintf("%c %s %s",
			fileType, class, m.Path)

	if fileType == 's' {
		ret = ret + fmt.Sprintf("=%s", m.RealPath)
	} else if m.DeviceInfo == nil {
		ret = ret + fmt.Sprintf(" %s %s %s",
					m.OctalMod, m.Owner, m.Group)
	} else {
		ret = ret + fmt.Sprintf(" %d %d %s %s %s", 
				m.DeviceInfo.Major, m.DeviceInfo.Minor, m.OctalMod,
				m.Owner, m.Group)
	}

	return ret
}

func usage() {
	fmt.Fprintf(os.Stderr,
		"usage: %s [-i] [-c class] [path ...]\n",
			progname)
	os.Exit(1)
}
