package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/alcoano/alcobinder/internal/app/alcobinder"
	"gopkg.in/yaml.v2"
)

type alcoBinderOptions struct {
	MarkdownsDirectory string `yaml:"markdowns_directory"`
	OutputFilePath     string `yaml:"output_file_path"`
	PageWidth          string `yaml:"page_width"`
	PageHeight         string `yaml:"page_height"`
	BaseFontSize       string `yaml:"base_font_size"`
}

func main() {
	config := &alcoBinderOptions{}

	if len(os.Args) != 2 || os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "help" {
		fmt.Println("\n" +
			"  alcobinder, a markdown(s) to HTML/PDF binder\n\n" +
			"  Usage:\n" +
			"    alcobinder PATH_TO_MANIFEST_FILE\n\n" +
			"    PATH_TO_MANIFEST_FILE points to a YAML file with the following values:\n\n" +
			"      markdowns_directory: directory containing markdown files.\n" +
			"      output_file_path:    path where printable output HTML file is to be written\n" +
			"      page_width:          width of the output pages\n" +
			"      page_height:         height of the output pages\n" +
			"      base_font_size:      font size of paragraph text\n")
		os.Exit(0)
	}

	manifestFile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:\n", err)
		os.Exit(66) // EX_NOINPUT
	}

	err = yaml.Unmarshal(manifestFile, config)
	if err != nil {
		fmt.Println("Error parsing manifest:\n", err)
		os.Exit(65) // EX_DATAERR
	}

	err = alcobinder.BindMarkdownsToFile(alcobinder.Options{
		OutputPath:         config.OutputFilePath,
		MarkdownsDirectory: config.MarkdownsDirectory,
		PageHeight:         config.PageHeight,
		PageWidth:          config.PageWidth,
		BaseFontSize:       config.BaseFontSize,
	})
	if err != nil {
		panic(err)
	}
}
