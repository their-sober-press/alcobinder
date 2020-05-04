package htmlpagecombiner_test

import (
	"bytes"
	"io"

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
		var err error
		JustBeforeEach(func() {
			output, err = CombinePages(pages)
			Expect(err).NotTo(HaveOccurred())
			document = goquery.NewDocumentFromNode(output)
		})
		It("creates an HTML document", func() {
			reader, writer := io.Pipe()
			go func() {
				err := html.Render(writer, output)
				Expect(err).NotTo(HaveOccurred())
				writer.Close()
			}()
			buf := bytes.Buffer{}
			_, err = buf.ReadFrom(reader)
			Expect(err).NotTo(HaveOccurred())
			outputString := buf.String()
			Expect(outputString).To(HavePrefix("<html>"))
		})
		It("adds the Pages.js script", func() {
			src, exists := document.Find("script").Attr("src")
			Expect(exists).To(BeTrue())
			Expect(src).To(Equal("https://unpkg.com/pagedjs/dist/paged.polyfill.js"))
		})

		Context("when there are pages", func() {
			BeforeEach(func() {
				pages = []Page{
					paginator.NewPageFromMarkdown("1", "**Hello** world!"),
					paginator.NewPageFromMarkdown("2", "_Good-bye_ world!"),
				}
			})
			It("Adds the pages as sections", func() {
				sections := document.Find("section")
				Expect(sections.Length()).To(Equal(2))
				Expect(sections.First().Text()).To(MatchRegexp("Hello world!"))
				Expect(sections.Last().Text()).To(MatchRegexp("Good-bye world!"))
			})
			It("Adds the page as numbers", func() {
				sections := document.Find("section")
				Expect(sections.First().Find("page-number").Text()).To(Equal("1"))
				Expect(sections.Last().Find("page-number").Text()).To(Equal("2"))
			})
			FIt("Formats the markdown", func() {
				sections := document.Find("section")
				Expect(sections.First().Html()).To(Equal("<strong>Hello</strong> world!"))
				Expect(sections.Last().Html()).To(Equal("<em>Good-bye</em> world!"))

			})

		})
	})
})
