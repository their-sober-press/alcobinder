package alcobinder

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/alcoano/alcobinder/pkg/htmlpagecombiner"
	"github.com/alcoano/alcobinder/pkg/paginator"
	"golang.org/x/net/html"
)

//BindMarkdownsToFile binds a directory of markdown files into a single PDF
func BindMarkdownsToFile(markdownDirectory string, outputPath string) error {
	outputFile, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	concatenatedMarkdown, err := concatenateMarkdown(markdownDirectory)
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

	htmlOutput, err := htmlpagecombiner.CombinePages(castedPages)
	if err != nil {
		return err
	}

	err = html.Render(outputFile, htmlOutput)
	if err != nil {
		return err
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
