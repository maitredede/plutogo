package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/maitredede/plutogo"
)

//go:embed content.html
var kHTMLContent string

func main() {
	version := plutogo.Version()
	buildinfo := plutogo.BuildInfo()
	fmt.Printf("plutobook version: %s\n%s\n", version, buildinfo)

	book, err := plutogo.NewBook(plutogo.PageSizeA4, plutogo.PageMarginsNarrow, plutogo.MediaTypePrint)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer book.Close()

	if err := book.LoadHTML(kHTMLContent, "", "", ""); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := book.WriteToPDF("hello.pdf"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
