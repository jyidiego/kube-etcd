package prettyprint_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGomegaMatchers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "internal/prettyprint")
}
