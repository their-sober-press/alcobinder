package htmlpagecombiner

import (
	"fmt"
	"html/template"
	"io"
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
func CombinePages(pages []Page, options Options) (*html.Node, error) {
	reader, writer := io.Pipe()
	emptyBookTemplate, err := template.New("EmptyBook").Parse(emptyBook)
	if err != nil {
		panic(err)
	}

	go func() {
		err = emptyBookTemplate.Execute(writer, options)
		if err != nil {
			panic(err)
		}
		writer.Close()
	}()

	document, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}
	body := document.Find("body").First()
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

	extensions := parser.Attributes | parser.Tables
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
	<script src="https://unpkg.com/pagedjs/dist/paged.polyfill.js"></script>
	<style>
		html {
			font-family: "Times New Roman MT Std", Times, serif;
			hyphens: auto;
		}
		section {
			page-break-after: always;
			string-set: pageNumber attr(data-page-number)
		}
		p {
			font-size: {{.BaseFontSize}};
			margin: 0;
			text-align: justify;
			font-weight: 200;
			line-height: 1.35;
			/* letter-spacing: 0.05em; */
		}
		p.indented {
			text-indent: 2ch;
		}
		blockquote + p::first-letter {
			line-height: 1;
			font-size: 1.5em;
			font-weight: bold;
		}
		blockquote {
			font-style: italic;
			margin-bottom: 2em;
		}
		h1 {
			string-set: chapterTitle content(text);
			page: firstPageInChapter;
			text-align: center;
			font-weight: lighter;
			font-size: 20pt;
		}
		@page {
			size: {{.PageWidth}} {{.PageHeight}};
			margin-bottom: 0.65in;
			margin-left: 0.65in;
			margin-right: 0.65in;
			word-break: break-word;
			@top-center {
				text-transform: uppercase;	
				font-size: 0.65em;
				letter-spacing: 0.25em;
				content: string(chapterTitle);
			}
		}

		@page firstPageInChapter {
			@top-center {
				content: "";
			}
			@top-right {
				content: "";
			}
			@top-left {
				content: "";
			}
			@bottom-center {
				font-size: 0.75em;
				content: string(pageNumber);
			}
		}

		@page :left {
			@top-left {
				font-size: 0.75em;
				content: string(pageNumber);
			}
		}

		@page :right {
			@top-right {
				font-size: 0.75em;
				content: string(pageNumber);
			}
		}

		p.footnote {
			margin-top: 1em;
			font-size: 0.5em;
		}

	</style>
</head>
<body>
</body>
</html>
`
