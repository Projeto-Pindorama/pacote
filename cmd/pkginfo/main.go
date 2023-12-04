package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

type DefaultOpts struct {
	Long bool `short:"l" long:"long" description:"#long listing"`
}

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
	var opts DefaultOpts

	parser := flags.NewParser(&opts, flags.Default)

	parser.Usage = "\n" +
		"pkginfo [-q] [-pi] [-x|l] [options] [pkg ...]\n" +
		"pkginfo -d device [-q] [-x|l] [options] [pkg ...]\n" +
		"where\n" +
		"  -q #quiet mode\n" +
		"  -p #select partially installed packages\n" +
		"  -i #select completely installed packages\n" +
		"  -x #extracted listing\n" +
		"  -l #long listing\n" +
		"  -r #relocation base\n" +
		"and options may include:\n" +
		"  -c category, [category...]\n" +
		"  -a architecture\n" +
		"  -v version\n"

  parser.ParseArgs(os.Args)

  entries, err := os.ReadDir(os.Getenv("pkgdir"))
  
  if err != nil {
    fmt.Println("An error: " + err.Error())
    return
  }
 
  pkgSlc := []*PackageInfo{}
  
  for _, e := range entries {
	  pkg, err := ReadInfo(os.Getenv("pkgdir") + "/" + e.Name())

	  if err != nil {
		  fmt.Println("An error: " + err.Error())
      return
	  }

    pkgSlc = append(pkgSlc, pkg)
  }

  for _, pkg := range pkgSlc {
	  if opts.Long {
	  	fmt.Printf("   PkgInst:\t%s\n", pkg.PkgInst)
	  	fmt.Printf("      Name:\t%s\n", pkg.Name)
	  	fmt.Printf("  Category:\t%s\n", pkg.Category)
	  	fmt.Printf("      Arch:\t%s\n", pkg.Arch)
	  	fmt.Printf("   Version:\t%s\n", pkg.Version)
	  	fmt.Printf("    Vendor:\t%s\n", pkg.Vendor)
		  fmt.Printf("      Desc:\t%s\n", pkg.Desc)
		  fmt.Printf("   Hotline:\t%s\n\n", pkg.Hotline)
	  } else {
		  fmt.Printf("%s\t%s\t\t%s\n", pkg.Category, pkg.Pkg, pkg.Desc)
	  }
  }
}
