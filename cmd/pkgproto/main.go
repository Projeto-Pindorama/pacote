package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Projeto-Pindorama/motoko/internal/archivum"
)


func main() {
	class := *flag.String("c", "none", "Class")
	flag.Parse()
	fs := archivum.NewUnixFS("/")
	readstdin := bufio.NewScanner(os.Stdin)
	for readstdin.Scan() {
		metadata, err := archivum.Scan(fs, readstdin.Text())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(MetadataToString(metadata, class))
	}
}

func MetadataToString(m *archivum.Metadata, class string) string {
	if m.DeviceInfo == nil {
		return fmt.Sprintf(
			"%c %s %s %s %s %s",
			m.FType,
			class,
			m.Path,
			m.OctalMod,
			m.Owner,
			m.Group,
		)
	} else if m.FType == 's' {
		return fmt.Sprintf(
			"%c %s %s=%s",
			m.FType,
			class,
			m.Path,
			m.RealPath,
		)
	} else {
		return fmt.Sprintf(
			"%c %s %s %d %d %s %s %s",
			m.FType,
			class,
			m.Path,
			m.DeviceInfo.Major,
			m.DeviceInfo.Minor,
			m.OctalMod,
			m.Owner,
			m.Group,
		)
	}

}
