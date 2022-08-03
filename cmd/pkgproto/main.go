package main

import (
	"bufio"
	"os"
	"fmt"

	"github.com/Projeto-Pindorama/motoko/internal/archivum"
)

func main() {
	readstdin := bufio.NewScanner(os.Stdin)
	for readstdin.Scan() {
		Metadata := archivum.Scan(readstdin.Text())
		fmt.Printf("%c %s %s %s %s %s %s\n",
		Metadata.FType,
		Metadata.Path,
		Metadata.Major,
		Metadata.Minor,
		Metadata.OctalMod,
		Metadata.Owner,
		Metadata.Group)
	}
}
