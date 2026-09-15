// Command create_mod produces a Go module source zip using golang.org/x/mod/zip,
// the same implementation the module proxy uses, so the member set and layout
// match proxy.golang.org by construction.
//
// It lives under .github/ so that it is excluded from the module zip (the
// staging step drops .github/ before zipping) and ignored by `go build ./...`,
// which skips directories whose name begins with a dot.
package main

import (
	"log"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatal("usage: create_mod <module-path> <version> <source-dir> <output-zip>")
	}
	m := module.Version{Path: os.Args[1], Version: os.Args[2]}
	f, err := os.Create(os.Args[4])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := zip.CreateFromDir(f, m, os.Args[3]); err != nil {
		log.Fatal(err)
	}
	log.Printf("created module zip: %s", os.Args[4])
}
