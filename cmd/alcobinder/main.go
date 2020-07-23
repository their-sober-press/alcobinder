package main

import (
	"fmt"
	"os"

	"github.com/their-sober-press/alcobinder/pkg/alcobinder"
)

func main() {
	if len(os.Args) != 4 || os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "help" {
		fmt.Println("\n" +
			"  alcobinder, a markdown(s) to HTML/PDF binder\n\n" +
			"  Usage:\n" +
			"    alcobinder INPUT_MARKDOWNS_DIR INPUT_CSS_FILE OUTPUT_FILE_PATH\n\n" +
			"    INPUT_MARKDOWNS_DIR    directory containing markdown files\n" +
			"    INPUT_CSS_FILE         path to file used to style the book\n" +
			"    OUTPUT_FILE_PATH       path where printable output HTML file is to be written\n")
		os.Exit(0)
	}

	err := alcobinder.BindMarkdownsToFile(os.Args[1], os.Args[2], os.Args[3])
	if err != nil {
		panic(err)
	}
}
