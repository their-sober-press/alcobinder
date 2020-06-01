package paginator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/their-sober-press/alcobinder/pkg/paginator"
)

const pageText = `
**First** line

_Second_ line
`

var _ = Describe("Page", func() {

	Describe(".NewPageFromMarkdown", func() {
		var page Page
		BeforeEach(func() {
			page = NewPageFromMarkdown("1", pageText)
		})

		It("adds the page number", func() {
			Expect(page.PageNumber).To(Equal("1"))
		})

		It("adds the markdown", func() {
			Expect(page.Markdown).To(Equal(pageText))
		})

	})
})
