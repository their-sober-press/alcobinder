package alcobinder_test

import (
	"fmt"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/alcoano/alcobinder/internal/app/alcobinder"
)

var _ = Describe("BindMarkdownsToFile", func() {
	var inputDirectory, outputFile, pageWidth, pageHeight, baseFontSize string
	var err error

	BeforeEach(func() {
		inputDirectory = "../../../test/data/single_markdown_file/"
		outputFile = fmt.Sprintf("../../../test/output/out%d.html", time.Now().UnixNano())
		pageHeight = "11in"
		pageWidth = "8.5in"
		baseFontSize = "10pt"
	})

	JustBeforeEach(func() {
		err = BindMarkdownsToFile(Options{
			MarkdownsDirectory: inputDirectory,
			OutputPath:         outputFile,
			PageHeight:         pageHeight,
			PageWidth:          pageWidth,
			BaseFontSize:       baseFontSize,
		})
	})

	It("should output a PDF file", func() {
		Expect(err).ToNot(HaveOccurred())
		info, err := os.Stat(outputFile)
		Expect(err).NotTo(HaveOccurred())
		Expect(info.IsDir()).To(BeFalse())
	})

	Describe("input validation", func() {
		Context("when MarkdownsDirectory is missing", func() {
			BeforeEach(func() {
				inputDirectory = ""
			})

			It("returns an error", func() {
				Expect(err).To(MatchError(MissingMarkdownsDirectory{}))
			})
		})

		Context("when OutputPath is missing", func() {
			BeforeEach(func() {
				outputFile = ""
			})

			It("returns an error", func() {
				Expect(err).To(MatchError(MissingOutputPath{}))
			})
		})

		Context("when PageWidth is missing", func() {
			BeforeEach(func() {
				pageWidth = ""
			})

			It("returns an error", func() {
				Expect(err).To(MatchError(MissingPageWidth{}))
			})
		})

		Context("when PageHeight is missing", func() {
			BeforeEach(func() {
				pageHeight = ""
			})

			It("returns an error", func() {
				Expect(err).To(MatchError(MissingPageHeight{}))
			})
		})

		Context("when BaseFontSize is missing", func() {
			BeforeEach(func() {
				baseFontSize = ""
			})

			It("returns an error", func() {
				Expect(err).To(MatchError(MissingBaseFontSize{}))
			})
		})
	})
})
