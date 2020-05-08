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
	const inputDirectory = "../../../test/data/single_markdown_file/"
	outputFile := fmt.Sprintf("../../../test/output/out%d.html", time.Now().UnixNano())
	BeforeEach(func() {
		err := BindMarkdownsToFile(inputDirectory, outputFile)
		Expect(err).ToNot(HaveOccurred())
	})

	It("should output a PDF file", func() {
		info, err := os.Stat(outputFile)
		Expect(err).NotTo(HaveOccurred())
		Expect(info.IsDir()).To(BeFalse())
	})
})
