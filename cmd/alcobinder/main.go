package main

import (
	"os"

	"github.com/alcoano/alcobinder/internal/app/alcobinder"
)

func main() {
	markdownsDir := os.Args[1]
	outputFile := os.Args[2]
	err := alcobinder.BindMarkdownsToPdf(markdownsDir, outputFile)
	if err != nil {
		panic(err)
	}
}
