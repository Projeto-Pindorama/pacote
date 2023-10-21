package main

import (
	"fmt"
	"os"
)

type PackageInfo struct {
	Vendor   string
	PkgInst  string
	Hotline  string
	Pkg      string
	Arch     string
	Desc     string
	Category string
	Name     string
	Version  string
}

func main() {
	pkg, err := ReadInfo(os.Args[1])

	if err != nil {
		fmt.Println("An error: " + err.Error())
		return
	}

	fmt.Printf("Vendor: %s\n", pkg.Vendor)
	fmt.Printf("PkgInst: %s\n", pkg.PkgInst)
	fmt.Printf("Hotline: %s\n", pkg.Hotline)
	fmt.Printf("Pkg: %s\n", pkg.Pkg)
	fmt.Printf("Arch: %s\n", pkg.Arch)
	fmt.Printf("Desc: %s\n", pkg.Desc)
	fmt.Printf("Category: %s\n", pkg.Category)
	fmt.Printf("Name: %s\n", pkg.Name)
	fmt.Printf("Version: %s\n", pkg.Version)
}
