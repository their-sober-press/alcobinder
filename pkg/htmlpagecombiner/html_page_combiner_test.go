package htmlpagecombiner_test

import (
	"github.com/PuerkitoBio/goquery"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/html"

	. "github.com/alcoano/alcobinder/pkg/htmlpagecombiner"
	"github.com/alcoano/alcobinder/pkg/paginator"
)

var _ = Describe("HtmlPageCombiner", func() {
	Describe("CombinePages", func() {
		var pages []Page
		var output *html.Node
		var document *goquery.Document
		var sections *goquery.Selection
		var err error
		const css = "h1 { color: red }"

		BeforeEach(func() {
			pages = []Page{}
		})

		JustBeforeEach(func() {
			output, err = CombinePages(pages, css)
			Expect(err).NotTo(HaveOccurred())
			document = goquery.NewDocumentFromNode(output)
			sections = document.Find("section")
		})

		It("adds the CSS", func() {
			Expect(document.Find("style").Text()).To(Equal(css))

		})

		Context("when there are pages", func() {
			BeforeEach(func() {
				pages = []Page{
					paginator.NewPageFromMarkdown("1", "**Hello** world!"),
					paginator.NewPageFromMarkdown("2", "_Good-bye_ world!"),
				}
			})
			It("Adds the pages as sections", func() {
				Expect(sections.Length()).To(Equal(2))
				Expect(sections.First().Text()).To(MatchRegexp("Hello world!"))
				Expect(sections.Last().Text()).To(MatchRegexp("Good-bye world!"))
			})
			It("Adds the page as numbers", func() {
				Expect(sections.First().AttrOr("data-page-number", "")).To(Equal("1"))
				Expect(sections.Last().AttrOr("data-page-number", "")).To(Equal("2"))
			})
			It("Formats the markdown", func() {
				Expect(sections.First().Html()).To(Equal("<p><strong>Hello</strong> world!</p>\n"))
				Expect(sections.Last().Html()).To(Equal("<p><em>Good-bye</em> world!</p>\n"))
			})
		})

		Context("when some paragraphs are indented", func() {
			BeforeEach(func() {
				pages = []Page{
					paginator.NewPageFromMarkdown("1", "  Indented paragraph."),
					paginator.NewPageFromMarkdown("2", "Nonindented paragraph."),
				}
			})
			It("Adds the class for indented paragraphs", func() {
				Expect(sections.First().Html()).To(Equal("<p class=\"indented\">Indented paragraph.</p>\n"))
				Expect(sections.Last().Html()).To(Equal("<p>Nonindented paragraph.</p>\n"))
			})
		})

		Context("when a paragraph starts with an asterisk", func() {
			BeforeEach(func() {
				pages = []Page{
					paginator.NewPageFromMarkdown("1", "I have a footnote.\\*\n\n\\*I am a footnote."),
				}
			})
			It("Adds the class for the footnote", func() {
				Expect(sections.First().Html()).To(Equal("<p>I have a footnote.*</p>\n\n<p class=\"footnote\">*I am a footnote.</p>\n"))
			})
		})

		Context("when there are some extra new lines", func() {
			BeforeEach(func() {
				pages = []Page{
					paginator.NewPageFromMarkdown("1", "  Indented paragraph.\n\n\n  Another indented paragraph."),
					paginator.NewPageFromMarkdown("2", "\n\n\nNonindented paragraph."),
				}
			})
			It("removes the extraneous new lines", func() {
				Expect(sections.First().Html()).To(Equal("<p class=\"indented\">Indented paragraph.</p>\n\n<p class=\"indented\">Another indented paragraph.</p>\n"))
				Expect(sections.Last().Html()).To(Equal("<p>Nonindented paragraph.</p>\n"))
			})
		})

		Context("when there are tables", func() {
			BeforeEach(func() {
				pages = []Page{
					paginator.NewPageFromMarkdown("1", tableMarkdown),
				}
			})
			It("renders the table", func() {
				Expect(sections.First().Html()).To(Equal("<table>\n<thead>\n<tr>\n<th>Header 1</th>\n<th>Header 2</th>\n</tr>\n</thead>\n\n<tbody>\n<tr>\n<td>Cell 1</td>\n<td>Cell 2</td>\n</tr>\n</tbody>\n</table>\n"))
			})
		})

	})
})

const tableMarkdown = `
| Header 1 | Header 2 |
|----------|----------|
| Cell 1   | Cell 2   |
`
