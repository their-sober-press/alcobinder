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
	var cursor pageCursor = NewPageCursor()
	var pageText string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PAGE ") {
			pageNumber, err := strconv.Atoi(line[5:])
			if err != nil {
				return nil, err
			}
			for i := 1; i < pageNumber; i++ {
				pages = append(pages, NewPageFromMarkdown(cursor.CurrentPage(), ""))
				cursor.Increment()
			}
			break
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PAGE ") {
			pageNumber := line[5:]
			if pageNumber != cursor.NextPage() {
				return nil, fmt.Errorf("PAGE %s missing", cursor.NextPage())
			}
			pageText = strings.TrimSpace(pageText)
			pages = append(pages, NewPageFromMarkdown(cursor.CurrentPage(), pageText))
			pageText = ""
			cursor.Increment()
		} else {
			pageText += (line + "\n")
		}
	}
	pageText = strings.TrimSpace(pageText)
	pages = append(pages, NewPageFromMarkdown(cursor.CurrentPage(), pageText))
	return pages, nil
}

type pageCursor struct {
	currentPage string
}

func NewPageCursor() pageCursor {
	return pageCursor{
		currentPage: "1",
	}
}

func (pc pageCursor) NextPage() string {
	pageInt, err := strconv.Atoi(pc.currentPage)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(pageInt + 1)
}

func (pc pageCursor) CurrentPage() string {
	return pc.currentPage
}

func (pc *pageCursor) Increment() {
	pageInt, err := strconv.Atoi(pc.currentPage)
	if err != nil {
		panic(err)
	}
	pc.currentPage = strconv.Itoa(pageInt + 1)
	fmt.Println(pc.currentPage)
}
