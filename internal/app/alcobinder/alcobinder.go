package alcobinder

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

//BindMarkdownsToPdf binds a directory of markdown files into a single PDF
func BindMarkdownsToPdf(markdownDirectory string, outputPath string) error {
	concatinatedMarkdown := ""
	err := filepath.Walk(markdownDirectory, func(path string, info os.FileInfo, err error) error {
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
		return err
	}

	fmt.Println(concatinatedMarkdown)

	return nil
}
