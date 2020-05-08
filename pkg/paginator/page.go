package paginator

//Page is a page
type Page struct {
	Markdown   string
	PageNumber string
}

//NewPageFromMarkdown makes a new page, generating the markdown into HTML
func NewPageFromMarkdown(pageNumber string, markdown string) Page {
	return Page{
		Markdown:   markdown,
		PageNumber: pageNumber,
	}
}

// GetMarkdown returns the markdown of the page
func (p Page) GetMarkdown() string {
	return p.Markdown
}

// GetPageNumber returns the page number of the page
func (p Page) GetPageNumber() string {
	return p.PageNumber
}
