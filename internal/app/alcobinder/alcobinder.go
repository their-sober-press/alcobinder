package alcobinder

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/alcoano/alcobinder/pkg/htmlpagecombiner"
	"github.com/alcoano/alcobinder/pkg/paginator"
	"golang.org/x/net/html"
)

//Options for specifying where an how to bind the markdowns into a document
//all field are required
type Options struct {
	MarkdownsDirectory string
	OutputPath         string
	PageWidth          string
	PageHeight         string
	BaseFontSize       string
}

//BindMarkdownsToFile binds a directory of markdown files into a single PDF
func BindMarkdownsToFile(options Options) error {
	err := validateOptions(options)
	if err != nil {
		return err
	}

	outputFile, err := os.OpenFile(options.OutputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	concatenatedMarkdown, err := concatenateMarkdown(options.MarkdownsDirectory)
	if err != nil {
		return err
	}

	pages, err := paginator.Paginate(concatenatedMarkdown)
	if err != nil {
		return err
	}

	castedPages := make([]htmlpagecombiner.Page, len(pages))
	for i, v := range pages {
		castedPages[i] = v
	}

	htmlOutput, err := htmlpagecombiner.CombinePages(castedPages, htmlpagecombiner.Options{
		PageHeight:   options.PageHeight,
		PageWidth:    options.PageWidth,
		BaseFontSize: options.BaseFontSize,
	})
	if err != nil {
		return err
	}

	err = html.Render(outputFile, htmlOutput)
	if err != nil {
		return err
	}

	return nil
}

func validateOptions(options Options) error {
	if options.MarkdownsDirectory == "" {
		return MissingMarkdownsDirectory{}
	}
	if options.OutputPath == "" {
		return MissingOutputPath{}
	}
	if options.PageWidth == "" {
		return MissingPageWidth{}
	}
	if options.PageHeight == "" {
		return MissingPageHeight{}
	}
	if options.BaseFontSize == "" {
		return MissingBaseFontSize{}
	}
	return nil
}

func concatenateMarkdown(markdownDirectory string) (string, error) {
	concatinatedMarkdown := ""
	err := filepath.Walk(markdownDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".md" {
			fileBytes, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			concatinatedMarkdown += string(fileBytes)
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return concatinatedMarkdown, nil
}

// MissingMarkdownsDirectory error for when options do not specify a directory with markdown files for combining
type MissingMarkdownsDirectory struct{}

func (e MissingMarkdownsDirectory) Error() string { return "missing markdowns directory in config" }

// MissingOutputPath error for when options do not specify path to put the generated HTML
type MissingOutputPath struct{}

func (e MissingOutputPath) Error() string { return "missing output path in config" }

// MissingPageWidth error for when options do not specify output file page width
type MissingPageWidth struct{}

func (e MissingPageWidth) Error() string { return "missing page width in config" }

// MissingPageHeight error for when options do not specify output file page height
type MissingPageHeight struct{}

func (e MissingPageHeight) Error() string { return "missing page height in config" }

// MissingBaseFontSize error for when options do not specify the base font size for page rendering
type MissingBaseFontSize struct{}

func (e MissingBaseFontSize) Error() string { return "missing base font size in config" }
