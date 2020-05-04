package htmlpagecombiner_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHtmlpagecombiner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Htmlpagecombiner Suite")
}
