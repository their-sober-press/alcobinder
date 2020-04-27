package paginator

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

//Paginate returns an array of pages including blank pages, split by PAGE x (where x is a number)
func Paginate(text string) ([]Page, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))
	pages := []Page{}
	var cursor pageCursor = newPageCursor()
	var pageText string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PAGE ") {
			pageNumber, err := strconv.Atoi(line[5:])
			if err != nil {
				return nil, err
			}
			for i := 1; i < pageNumber; i++ {
				pages = append(pages, NewPageFromMarkdown(cursor.currentPage(), ""))
				cursor.increment()
			}
			break
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PAGE ") {
			pageNumber := line[5:]
			if pageNumber != cursor.nextPage() {
				return nil, fmt.Errorf("PAGE %s missing", cursor.nextPage())
			}
			pageText = formatWhiteSpace(pageText)
			pages = append(pages, NewPageFromMarkdown(cursor.currentPage(), pageText))
			pageText = ""
			cursor.increment()
		} else {
			pageText += (line + "\n")
		}
	}
	pageText = formatWhiteSpace(pageText)
	pages = append(pages, NewPageFromMarkdown(cursor.currentPage(), pageText))
	return pages, nil
}

func formatWhiteSpace(pageText string) string {
	outText := pageText
	outText = strings.TrimLeft(outText, "\n")
	outText = strings.TrimRight(outText, "\n")
	paragraphs := strings.Split(outText, "\n\n")
	var newParagraphs []string
	for _, paragraph := range paragraphs {
		if strings.HasPrefix(paragraph, " ") {
			paragraph = "{.new-paragraph}\n" + paragraph
		}
		newParagraphs = append(newParagraphs, paragraph)
	}
	return strings.Join(newParagraphs, "\n\n")
}

type pageCursor struct {
	page string
}

func newPageCursor() pageCursor {
	return pageCursor{
		page: "1",
	}
}

func (pc pageCursor) nextPage() string {
	pageInt, err := strconv.Atoi(pc.page)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(pageInt + 1)
}

func (pc pageCursor) currentPage() string {
	return pc.page
}

func (pc *pageCursor) increment() {
	pageInt, err := strconv.Atoi(pc.page)
	if err != nil {
		panic(err)
	}
	pc.page = strconv.Itoa(pageInt + 1)
}
