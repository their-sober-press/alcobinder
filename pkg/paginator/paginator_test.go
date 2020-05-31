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

const nonsensicalPageNumber = `
PAGE 1
This is page one.

PAGE $
This is the dollar page.
`

const withRomanNumerals = `
PAGE i
This is preface page one.

PAGE ii
This is preface page two.

PAGE 1
This is page one.
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
					PageNumber: "1",
				},
				{
					Markdown:   "This is page two.\n\nStill page two.",
					PageNumber: "2",
				},
				{
					Markdown:   "This is page three.",
					PageNumber: "3",
				},
			}))
		})
	})

	Context("when first pages are missing", func() {
		BeforeEach(func() {
			input = missingPagesAtBeginning
		})

		It("ignores the missing pages", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(Equal([]Page{
				{
					Markdown:   "This is page three.",
					PageNumber: "3",
				},
				{
					Markdown:   "This is page four.",
					PageNumber: "4",
				},
			}))
		})
	})

	Context("when pages are missing in the middle", func() {
		BeforeEach(func() {
			input = missingMiddlePages
		})

		It("errors", func() {
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError("PAGE \"2\" missing, \"3\" found"))
		})
	})

	Context("when the first page has a nonsensical number", func() {
		BeforeEach(func() {
			input = "PAGE !"
		})

		It("errors", func() {
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError("first page has an invalid page number \"!\""))
		})
	})

	Context("when there is a nonsensical page number in the middle", func() {
		BeforeEach(func() {
			input = nonsensicalPageNumber
		})

		It("errors", func() {
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError("invalid page number \"$\" after page \"1\""))
		})
	})

	Context("when pages start with roman numerals", func() {
		BeforeEach(func() {
			input = withRomanNumerals
		})

		It("paginates roman numerals until it hits arabic numerals", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(Equal([]Page{
				{
					Markdown:   "This is preface page one.",
					PageNumber: "i",
				},
				{
					Markdown:   "This is preface page two.",
					PageNumber: "ii",
				},
				{
					Markdown:   "This is page one.",
					PageNumber: "1",
				},
			}))
		})
	})
})
