package paginator_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPaginator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Paginator Suite")
}
