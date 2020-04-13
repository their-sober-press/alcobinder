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
PAGE 2
This is page two.

PAGE 3
This is page three.
`

const missingMiddlePages = `
PAGE 1
This is page one.

PAGE 3
This is page three.
`

var _ = Describe("Paginator", func() {
	var input string
	var output []string
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
			Expect(output).To(Equal([]string{
				"This is page one.",
				"This is page two.\n\nStill page two.",
				"This is page three.",
			}))
		})
	})

	Context("when first pages are missing", func() {
		BeforeEach(func() {
			input = missingPagesAtBeginning
		})

		It("paginates fills in missing pages", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(Equal([]string{
				"",
				"This is page two.",
				"This is page three.",
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
})
