package acceptance_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var pathToActionBinary string

func TestAcceptance(t *testing.T) {
	BeforeSuite(func() {
		var err error

		pathToActionBinary, err = Build("../cmd/dependency-action/main.go")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterSuite(func() {
		CleanupBuildArtifacts()
	})

	RegisterFailHandler(Fail)
	RunSpecs(t, "Acceptance Suite")
}
