package paginator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/their-sober-press/alcobinder/pkg/paginator"
)

var _ = Describe("ParseNumeralCursor", func() {
	var numeral NumeralCursor
	var err error

	It("parses arabic numerals", func() {
		numeral, err := NewNumeralCursorFromString("2")
		Expect(err).ToNot(HaveOccurred())
		Expect(numeral.Current()).To(Equal("2"))
	})

	It("parses roman numerals", func() {
		numeral, err := NewNumeralCursorFromString("ii")
		Expect(err).ToNot(HaveOccurred())
		Expect(numeral.Current()).To(Equal("ii"))
	})

	It("errors on non-numeral strings", func() {
		numeral, err := NewNumeralCursorFromString("$")
		Expect(err).To(MatchError("invalid numeral \"$\""))
		Expect(numeral).To(BeNil())
	})

	Describe("ArabicNumeralCursor", func() {
		BeforeEach(func() {
			numeral = NewArabicNumeralCursor()
		})

		Describe("NewArabicNumeralCursor", func() {
			It("starts at 1", func() {
				Expect(numeral.Current()).To(Equal("1"))
			})
		})

		Describe("PeekNext", func() {
			It("returns the next numeral without incrementing the numeral", func() {
				Expect(numeral.PeekNext()).To(Equal("2"))
				Expect(numeral.Current()).To(Equal("1"))
			})
		})

		Describe("Increment", func() {
			It("increments the numeral", func() {
				numeral.Increment()
				Expect(numeral.Current()).To(Equal("2"))
			})
		})

		Describe("Set", func() {
			Context("when the value is the is a valid numeral", func() {
				BeforeEach(func() {
					err = numeral.Set("5")
				})

				It("does not error", func() {
					Expect(err).NotTo(HaveOccurred())
				})

				It("sets the numeral to value", func() {
					Expect(numeral.Current()).To(Equal("5"))
				})
			})

			Context("when the value is the is an invalid numeral", func() {
				BeforeEach(func() {
					err = numeral.Set("$")
				})

				It("errors", func() {
					Expect(err).To(MatchError("invalid arabic numeral \"$\""))
				})

				It("does not set the numeral", func() {
					Expect(numeral.Current()).To(Equal("1"))
				})
			})
		})

		Describe("IsNextNumeral", func() {
			var isNext, isSameType bool
			var candidate string

			JustBeforeEach(func() {
				isNext, isSameType = numeral.IsNextValue(candidate)
			})
			Context("when the value is not a valid arabic numeral", func() {
				BeforeEach(func() {
					candidate = "i"
				})

				It("returns false, false", func() {
					Expect(isNext).To(BeFalse())
					Expect(isSameType).To(BeFalse())
				})
			})

			Context("when the value is a valid arabic numeral that is not next", func() {
				BeforeEach(func() {
					candidate = "5"
				})

				It("returns false, true", func() {
					Expect(isNext).To(BeFalse())
					Expect(isSameType).To(BeTrue())
				})
			})

			Context("when the value is a valid arabic numeral that is next", func() {
				BeforeEach(func() {
					candidate = "2"
				})

				It("returns true, true", func() {
					Expect(isNext).To(BeTrue())
					Expect(isSameType).To(BeTrue())
				})
			})
		})
	})

	Describe("RomanNumeralCursor", func() {
		BeforeEach(func() {
			numeral = NewRomanNumeralCursor()
		})

		Describe("NewArabicNumeralCursor", func() {
			It("starts at 1", func() {
				Expect(numeral.Current()).To(Equal("i"))
			})
		})

		Describe("PeekNext", func() {
			It("returns the next numeral without incrementing the numeral", func() {
				Expect(numeral.PeekNext()).To(Equal("ii"))
				Expect(numeral.Current()).To(Equal("i"))
			})
		})

		Describe("Increment", func() {
			It("increments the numeral", func() {
				numeral.Increment()
				Expect(numeral.Current()).To(Equal("ii"))
			})
		})

		Describe("Set", func() {
			Context("when the value is the is a valid numeral", func() {
				BeforeEach(func() {
					err = numeral.Set("vi")
				})

				It("does not error", func() {
					Expect(err).NotTo(HaveOccurred())
				})

				It("sets the numeral to value", func() {
					Expect(numeral.Current()).To(Equal("vi"))
				})
			})

			Context("when the value is the is an invalid numeral", func() {
				BeforeEach(func() {
					err = numeral.Set("$")
				})

				It("errors", func() {
					Expect(err).To(MatchError("invalid roman numeral \"$\""))
				})

				It("does not set the numeral", func() {
					Expect(numeral.Current()).To(Equal("i"))
				})
			})
		})

		Describe("IsNextNumeral", func() {
			var isNext, isSameType bool
			var candidate string

			JustBeforeEach(func() {
				isNext, isSameType = numeral.IsNextValue(candidate)
			})
			Context("when the value is not a valid roman numeral", func() {
				BeforeEach(func() {
					candidate = "2"
				})

				It("returns false, false", func() {
					Expect(isNext).To(BeFalse())
					Expect(isSameType).To(BeFalse())
				})
			})

			Context("when the value is a valid roman numeral that is not next", func() {
				BeforeEach(func() {
					candidate = "iii"
				})

				It("returns false, true", func() {
					Expect(isNext).To(BeFalse())
					Expect(isSameType).To(BeTrue())
				})
			})

			Context("when the value is a valid roman numeral that is next", func() {
				BeforeEach(func() {
					candidate = "ii"
				})

				It("returns true, true", func() {
					Expect(isNext).To(BeTrue())
					Expect(isSameType).To(BeTrue())
				})
			})
		})
	})
})
