package paginator

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

//Paginate returns an array of pages including blank pages, split by PAGE x (where x is a number)
func Paginate(text string) ([]string, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))
	page := ""
	pages := []string{}
	var nextPageNumber int

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PAGE ") {
			pageNumber, err := strconv.Atoi(line[5:])
			if err != nil {
				return nil, err
			}
			for i := 1; i < pageNumber; i++ {
				pages = append(pages, "")
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
			pages = append(pages, strings.TrimSpace(page))
			page = ""
			nextPageNumber++
		} else {
			page += (line + "\n")
		}
	}
	pages = append(pages, strings.TrimSpace(page))
	return pages, nil
}
