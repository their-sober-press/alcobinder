package paginator

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

//Page is a page
type Page struct {
	Markdown   string
	HTML       string
	PageNumber string
}

//NewPageFromMarkdown makes a new page, generating the markdown into HTML
func NewPageFromMarkdown(pageNumber string, markdown string) Page {
	return Page{
		Markdown:   markdown,
		PageNumber: pageNumber,
		HTML:       renderHTML(markdown),
	}
}

func renderHTML(md string) string {
	opts := html.RendererOptions{}
	renderer := html.NewRenderer(opts)

	extensions := parser.Attributes | parser.Tables
	parser := parser.NewWithExtensions(extensions)

	html := markdown.ToHTML([]byte(md), parser, renderer)
	return string(html)
}

// GetMarkdown returns the markdown of the page
func (p Page) GetMarkdown() string {
	return p.Markdown
}

// GetPageNumber returns the page number of the page
func (p Page) GetPageNumber() string {
	return p.PageNumber
}
