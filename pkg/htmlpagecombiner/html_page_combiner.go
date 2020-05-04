package htmlpagecombiner

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"golang.org/x/net/html"
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
		body.AppendHtml(fmt.Sprintf("<section><page-number>%s</page-number>%s</section>", page.GetPageNumber(), page.GetMarkdown()))
	}
	return document.Get(0), nil
}

func renderHTML(md string) string {
	opts := html.RendererOptions{}
	renderer := html.NewRenderer(opts)

	extensions := parser.Attributes | parser.Tables
	parser := parser.NewWithExtensions(extensions)

	html := markdown.ToHTML([]byte(md), parser, renderer)
	return string(html)
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
