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
				{
					PageNumber: "1",
				},
				{
					PageNumber: "2",
				},
				{
					PageNumber: "3",
					Markdown:   "This is page three.",
					HTML:       "<p>This is page three.</p>\n",
				},
				{
					PageNumber: "4",
					Markdown:   "This is page four.",
					HTML:       "<p>This is page four.</p>\n",
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
})
