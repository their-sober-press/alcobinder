package htmlpagecombiner

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"golang.org/x/net/html"

	htmlrenderer "github.com/gomarkdown/markdown/html"
)

// Page any page that can give its Markdown and PageNumber
type Page interface {
	GetMarkdown() string
	GetPageNumber() string
}

// Options are arguments that are used to render the HTML document
type Options struct {
	PageWidth    string
	PageHeight   string
	BaseFontSize string
}

// CombinePages combines pages into a printer friendly HTML document
func CombinePages(pages []Page, css string) (*html.Node, error) {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(emptyBook))
	if err != nil {
		return nil, err
	}
	body := document.Find("body").First()
	document.Find("style").SetText(css)
	for _, page := range pages {
		pageHtml := renderHTML(page.GetMarkdown())
		body.AppendHtml(fmt.Sprintf(`<section data-page-number="%s">%s</section>`, page.GetPageNumber(), pageHtml))
	}
	return document.Get(0), nil
}

func renderHTML(md string) string {
	markdownToRender := decorateWithClasses(md)

	opts := htmlrenderer.RendererOptions{}
	renderer := htmlrenderer.NewRenderer(opts)

	extensions := parser.Attributes | parser.Tables //TODO: smart number start
	parser := parser.NewWithExtensions(extensions)

	html := markdown.ToHTML([]byte(markdownToRender), parser, renderer)
	return string(html)
}

func decorateWithClasses(pageText string) string {
	outText := pageText
	paragraphs := strings.Split(outText, "\n\n")
	var newParagraphs []string
	for _, paragraph := range paragraphs {
		paragraph = strings.TrimLeft(paragraph, "\n")
		paragraph = strings.TrimRight(paragraph, "\n")
		if strings.HasPrefix(paragraph, " ") {
			paragraph = "{.indented}\n" + paragraph
		} else if strings.HasPrefix(paragraph, "\\*") {
			paragraph = "{.footnote}\n" + paragraph
		}
		newParagraphs = append(newParagraphs, paragraph)
	}
	return strings.Join(newParagraphs, "\n\n")
}

const emptyBook = `
<!doctype html>
<html lang=en>
<head>
	<style></style>
</head>
<body>
</body>
</html>
`
