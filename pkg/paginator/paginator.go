package paginator

import (
	"bufio"
	"fmt"
	"strings"
)

//Paginate returns an array of pages including blank pages, split by PAGE x (where x is a number)
func Paginate(text string) ([]Page, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))
	pages := []Page{}
	var cursor NumeralCursor = NewArabicNumeralCursor()
	var pageText string
	var err error

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PAGE ") {
			firstPageNumber := line[5:]
			cursor, err = NewNumeralCursorFromString(firstPageNumber)
			if err != nil {
				return nil, fmt.Errorf("first page has an invalid page number %#v", firstPageNumber)
			}
			break
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PAGE ") {
			pageText = formatWhiteSpace(pageText)
			page := NewPageFromMarkdown(cursor.Current(), pageText)
			pageText = ""

			pageNumber := line[5:]
			err := safelyIncrement(&cursor, pageNumber)
			if err != nil {
				return nil, err
			}
			pages = append(pages, page)
		} else {
			pageText += (line + "\n")
		}
	}
	pageText = formatWhiteSpace(pageText)
	pages = append(pages, NewPageFromMarkdown(cursor.Current(), pageText))
	return pages, nil
}

func formatWhiteSpace(pageText string) string {
	outText := pageText
	outText = strings.TrimLeft(outText, "\n")
	outText = strings.TrimRight(outText, "\n")
	return outText
}

func safelyIncrement(numeralCursorPtr *NumeralCursor, nextPage string) (err error) {
	numeralCursor := *numeralCursorPtr
	isNextPage, isSameType := numeralCursor.IsNextValue(nextPage)
	if !isSameType {
		newNumeralCursor, err := NewNumeralCursorFromString(nextPage)
		if err != nil {
			return fmt.Errorf("invalid page number %#v after page %#v", nextPage, numeralCursor.Current())
		}
		*numeralCursorPtr = newNumeralCursor
		return nil
	}
	if !isNextPage {
		return fmt.Errorf("PAGE %#v missing, %#v found", numeralCursor.PeekNext(), nextPage)
	}
	numeralCursor.Increment()
	return nil
}
