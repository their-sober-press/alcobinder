package alcobinder_test

import (
	"fmt"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/their-sober-press/alcobinder/internal/app/alcobinder"
)

var _ = Describe("BindMarkdownsToFile", func() {
	var inputDirectory, inputCSSFile, outputFile string
	var err error

	BeforeEach(func() {
		inputDirectory = "../../../test/data/single_markdown_file/"
		inputCSSFile = "../../../test/data/test.css"
		outputFile = fmt.Sprintf("../../../test/output/out%d.html", time.Now().UnixNano())
	})

	JustBeforeEach(func() {
		err = BindMarkdownsToFile(inputDirectory, inputCSSFile, outputFile)
	})

	It("should output a file", func() {
		Expect(err).ToNot(HaveOccurred())
		info, err := os.Stat(outputFile)
		Expect(err).NotTo(HaveOccurred())
		Expect(info.IsDir()).To(BeFalse())
	})
})
