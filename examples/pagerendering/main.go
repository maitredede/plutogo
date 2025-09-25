package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"

	"github.com/maitredede/plutogo"
)

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

	// Load the HTML content from file
	// book.LoadURL("Alice’s Adventures in Wonderland.html")
	if err := book.LoadURL("https://www.gutenberg.org/cache/epub/11/pg11-images.html", "", ""); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pageSize := book.GetPageSize()
	pageWidth := int(math.Ceil(pageSize.Width / plutogo.UnitsPX))
	pageHeight := int(math.Ceil(pageSize.Height / plutogo.UnitsPX))

	canvas, err := plutogo.NewImageCanvas(pageWidth, pageHeight, plutogo.ImageFormatARGB32)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer canvas.Close()

	// Render the first 3 pages to PNG files
	for pageIndex := 0; pageIndex < 3; pageIndex++ {
		filename := fmt.Sprintf("page-%d.png", pageIndex+1)

		// Clear the canvas to white before rendering each page
		if err := canvas.ClearSurface(1, 1, 1, 1); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Render the page onto the canvas
		if err := book.RenderPage(canvas, pageIndex); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Save the canvas to a PNG file
		if err := canvas.WriteToPNG(filename); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// Export pages 1 to 3 (inclusive) to PDF with step=1 (every page in order)
	if err := book.WriteToPDFRange("Wonderland.pdf", 1, 3, 1); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
