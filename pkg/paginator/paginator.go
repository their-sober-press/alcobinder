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
			cursor.setPage(strconv.Itoa(pageNumber))
			break
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PAGE ") {
			pageNumber := line[5:]
			if pageNumber != cursor.nextPage() {
				fmt.Println("36" != "36")
				return nil, fmt.Errorf("PAGE %#v missing, %#v found", cursor.nextPage(), pageNumber)
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
	return outText
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

func (pc *pageCursor) setPage(pageNumber string) {
	pc.page = pageNumber
}
