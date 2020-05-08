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

// CombinePages combines pages into a printer friendly HTML document
func CombinePages(pages []Page) (*html.Node, error) {
	reader := strings.NewReader(emptyBook)
	document, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}
	body := document.Find("body").First()
	for _, page := range pages {
		pageHtml := renderHTML(page.GetMarkdown())
		body.AppendHtml(fmt.Sprintf("<section><page-number>%s</page-number>%s</section>", page.GetPageNumber(), pageHtml))
	}
	return document.Get(0), nil
}

func renderHTML(md string) string {
	markdownToRender := markIndentation(md)

	opts := htmlrenderer.RendererOptions{}
	renderer := htmlrenderer.NewRenderer(opts)

	extensions := parser.Attributes | parser.Tables
	parser := parser.NewWithExtensions(extensions)

	html := markdown.ToHTML([]byte(markdownToRender), parser, renderer)
	return string(html)
}

func markIndentation(pageText string) string {
	outText := pageText
	paragraphs := strings.Split(outText, "\n\n")
	var newParagraphs []string
	for _, paragraph := range paragraphs {
		paragraph = strings.TrimLeft(paragraph, "\n")
		paragraph = strings.TrimRight(paragraph, "\n")
		if strings.HasPrefix(paragraph, " ") {
			paragraph = "{.indented}\n" + paragraph
		}
		newParagraphs = append(newParagraphs, paragraph)
	}
	return strings.Join(newParagraphs, "\n\n")
}

const emptyBook = `
<html>
<head>
	<script src="https://unpkg.com/pagedjs/dist/paged.polyfill.js"></script>
</head>
<body>
</body>
</html>
`
