package paginator

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

//Page is a page
type Page struct {
	Markdown   string
	HTML       string
	PageNumber string
}

//Paginate returns an array of pages including blank pages, split by PAGE x (where x is a number)
func Paginate(text string) ([]Page, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))
	pages := []Page{}
	page := Page{}
	var nextPageNumber int

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PAGE ") {
			pageNumber, err := strconv.Atoi(line[5:])
			if err != nil {
				return nil, err
			}
			page.PageNumber = strconv.Itoa(pageNumber)
			for i := 1; i < pageNumber; i++ {
				pages = append(pages, Page{
					Markdown:   "",
					PageNumber: strconv.Itoa(i),
				})
			}
			nextPageNumber = pageNumber + 1
			break
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PAGE ") {
			pageNumber, err := strconv.Atoi(line[5:])
			if err != nil {
				return nil, err
			}
			if pageNumber != nextPageNumber {
				return nil, fmt.Errorf("PAGE %d missing", nextPageNumber)
			}
			page.Markdown = strings.TrimSpace(page.Markdown)
			page.HTML = renderHTML(page.Markdown)
			pages = append(pages, page)
			page = Page{
				PageNumber: strconv.Itoa(nextPageNumber),
			}
			nextPageNumber++
		} else {
			page.Markdown += (line + "\n")
		}
	}
	page.Markdown = strings.TrimSpace(page.Markdown)
	page.HTML = renderHTML(page.Markdown)
	pages = append(pages, page)
	return pages, nil
}

func renderHTML(md string) string {
	opts := html.RendererOptions{}
	renderer := html.NewRenderer(opts)

	html := markdown.ToHTML([]byte(md), nil, renderer)
	return string(html)
}
