package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInfo(filePath string) (*PackageInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var pkg PackageInfo
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.SplitN(line, "=", 2)
		if len(fields) != 2 {
			continue
		}

		key := fields[0]
		value := fields[1]

		switch key {
		case "VENDOR":
			pkg.Vendor = value
		case "PKGINST":
			pkg.PkgInst = value
		case "HOTLINE":
			pkg.Hotline = value
		case "PKG":
			pkg.Pkg = value
		case "ARCH":
			pkg.Arch = value
		case "DESC":
			pkg.Desc = value
		case "CATEGORY":
			pkg.Category = value
		case "NAME":
			pkg.Name = value
		case "VERSION":
			pkg.Version = value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return nil, err
	}

	return &pkg, nil
}
