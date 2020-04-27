package paginator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/alcoano/alcobinder/pkg/paginator"
)

const noMissingPages = `

PAGE 1

This is page one.


PAGE 2
This is page two.

Still page two.
PAGE 3
This is page three.

`

const missingPagesAtBeginning = `
PAGE 3
This is page three.

PAGE 4
This is page four.
`

const missingMiddlePages = `
PAGE 1
This is page one.

PAGE 3
This is page three.
`

const withFormatting = `
PAGE 1
This is page _one_.

PAGE 2
This is page **two**.
`

const withIndents = `
PAGE 1
# Heading

  This is page _one_.

PAGE 2
This is page **two**.
`

const withTables = `
PAGE 1

| Header 1 | Header 2 |
|----------|----------|
| Cell 1   | Cell 2   |
`

var _ = Describe("Paginate", func() {
	var input string
	var output []Page
	var err error

	JustBeforeEach(func() {
		output, err = Paginate(input)
	})

	Context("when no pages are missing", func() {
		BeforeEach(func() {
			input = noMissingPages
		})

		It("paginates splitting on PAGE x", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(Equal([]Page{
				{
					Markdown:   "This is page one.",
					HTML:       "<p>This is page one.</p>\n",
					PageNumber: "1",
				},
				{
					Markdown:   "This is page two.\n\nStill page two.",
					HTML:       "<p>This is page two.</p>\n\n<p>Still page two.</p>\n",
					PageNumber: "2",
				},
				{
					Markdown:   "This is page three.",
					HTML:       "<p>This is page three.</p>\n",
					PageNumber: "3",
				},
			}))
		})
	})

	Context("when first pages are missing", func() {
		BeforeEach(func() {
			input = missingPagesAtBeginning
		})

		It("paginates fills in missing pages", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(Equal([]Page{
				{Markdown: "", HTML: "", PageNumber: "1"},
				{Markdown: "", HTML: "", PageNumber: "2"},
				{
					Markdown:   "This is page three.",
					HTML:       "<p>This is page three.</p>\n",
					PageNumber: "3",
				},
				{
					Markdown:   "This is page four.",
					HTML:       "<p>This is page four.</p>\n",
					PageNumber: "4",
				},
			}))
		})
	})

	Context("when pages are missing in the middle", func() {
		BeforeEach(func() {
			input = missingMiddlePages
		})

		It("paginates fills in missing pages", func() {
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError("PAGE 2 missing"))
		})
	})

	Context("when pages have formatting", func() {
		BeforeEach(func() {
			input = withFormatting
		})

		It("converts the markdown formatting into HTML", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(Equal([]Page{
				{
					Markdown:   "This is page _one_.",
					HTML:       "<p>This is page <em>one</em>.</p>\n",
					PageNumber: "1",
				},
				{
					Markdown:   "This is page **two**.",
					HTML:       "<p>This is page <strong>two</strong>.</p>\n",
					PageNumber: "2",
				},
			}))
		})
	})

	Context("when lines don't have indents", func() {
		BeforeEach(func() {
			input = withIndents
		})

		It("gives them the 'new-paragraph' class", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(Equal([]Page{
				{
					Markdown:   "# Heading\n\n{.new-paragraph}\n  This is page _one_.",
					HTML:       "<h1>Heading</h1>\n\n<p class=\"new-paragraph\">This is page <em>one</em>.</p>\n",
					PageNumber: "1",
				},
				{
					Markdown:   "This is page **two**.",
					HTML:       "<p>This is page <strong>two</strong>.</p>\n",
					PageNumber: "2",
				},
			}))
		})
	})

	Context("when there are tables", func() {
		BeforeEach(func() {
			input = withTables
		})

		It("gives them the 'new-paragraph' class", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(Equal([]Page{
				{
					Markdown:   "| Header 1 | Header 2 |\n|----------|----------|\n| Cell 1   | Cell 2   |",
					HTML:       "<table>\n<thead>\n<tr>\n<th>Header 1</th>\n<th>Header 2</th>\n</tr>\n</thead>\n\n<tbody>\n<tr>\n<td>Cell 1</td>\n<td>Cell 2</td>\n</tr>\n</tbody>\n</table>\n",
					PageNumber: "1",
				},
			}))
		})
	})
})
