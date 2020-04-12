package alcobinder

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/thecodingmachine/gotenberg-go-client/v7"
)

//BindMarkdownsToPdf binds a directory of markdown files into a single PDF
func BindMarkdownsToPdf(markdownDirectory string, outputPath string) error {
	concatenatedMarkdown, err := concatenateMarkdown(markdownDirectory)
	if err != nil {
		return err
	}

	client := &gotenberg.Client{Hostname: "http://localhost:3000"}
	html := `
		<!doctype html>
			<html lang="en">
			<head>
				<meta charset="utf-8">
				<title>My PDF</title>
			</head>
			<body>
				{{ toHTML .DirPath "file.md" }}
			</body>
		</html>
	`
	index, err := gotenberg.NewDocumentFromString("index.html", html)
	if err != nil {
		return err
	}

	markdownAsset, err := gotenberg.NewDocumentFromString("file.md", concatenatedMarkdown)

	req := gotenberg.NewMarkdownRequest(index)
	req.Assets(markdownAsset)

	client.Store(req, outputPath)

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
