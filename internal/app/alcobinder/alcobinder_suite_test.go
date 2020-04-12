package alcobinder_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAlcobinder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Alcobinder Suite")
}
