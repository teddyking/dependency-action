package acceptance_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("the dependency-action binary", func() {
	var (
		env         []string
		session     *Session
		testHomeDir string
	)

	BeforeEach(func() {
		var err error

		testHomeDir, err = ioutil.TempDir("", "dependency-action-test")
		Expect(err).NotTo(HaveOccurred())

		env = []string{
			fmt.Sprintf("HOME=%s", testHomeDir),
		}
	})

	JustBeforeEach(func() {
		command := exec.Command(pathToActionBinary)
		command.Env = env

		var err error
		session, err = Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(os.RemoveAll(testHomeDir)).To(Succeed())
	})

	It("exits with an exit code of 0", func() {
		Eventually(session).Should(Exit(0))
	})

	When("the INPUT_TGZDEPS env var is set", func() {
		BeforeEach(func() {
			env = append(env, []string{
				fmt.Sprintf("INPUT_TGZDEPS=%s/dep1.tgz", testAssetsURL),
			}...)
		})

		It("downloads and extracts the tgz dependencies to $HOME", func() {
			Eventually(session).Should(Exit(0))

			Expect(filepath.Join(testHomeDir, "dep1")).To(BeADirectory())
			Expect(filepath.Join(testHomeDir, "dep1", "bin", "cake")).To(BeAnExistingFile())
		})
	})
})
