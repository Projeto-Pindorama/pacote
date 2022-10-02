package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Projeto-Pindorama/motoko/internal/archivum"
)

func main() {
	fs := archivum.NewUnixFS("/")
	readstdin := bufio.NewScanner(os.Stdin)
	for readstdin.Scan() {
		metadata, err := archivum.Scan(fs, readstdin.Text())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(archivum.MetadataToString(metadata))
	}
}
