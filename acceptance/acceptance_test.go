package acceptance_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
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

	When("the INPUT_DEPS env var is set and contains 1 tgz file url", func() {
		BeforeEach(func() {
			env = append(env, []string{
				fmt.Sprintf("INPUT_DEPS=%s/dep1.tgz", testAssetsURL),
			}...)
		})

		It("downloads and extracts the tgz dependency to $HOME", func() {
			Eventually(session).Should(Exit(0))

			Expect(filepath.Join(testHomeDir, "dep1")).To(BeADirectory())
			Expect(filepath.Join(testHomeDir, "dep1", "bin", "cake")).To(BeAnExistingFile())
		})
	})

	When("the INPUT_DEPS env var is set and contains 2 tgz file urls", func() {
		BeforeEach(func() {
			env = append(env, []string{
				fmt.Sprintf("INPUT_DEPS=%s/dep1.tgz,%s/dep2.tar.gz", testAssetsURL, testAssetsURL),
			}...)
		})

		It("downloads and extracts the tgz dependencies to $HOME", func() {
			Eventually(session).Should(Exit(0))

			Expect(filepath.Join(testHomeDir, "dep1")).To(BeADirectory())
			Expect(filepath.Join(testHomeDir, "dep1", "bin", "cake")).To(BeAnExistingFile())

			Expect(filepath.Join(testHomeDir, "dep2")).To(BeADirectory())
			Expect(filepath.Join(testHomeDir, "dep2", "bin", "cake")).To(BeAnExistingFile())
		})
	})

	When("the INPUT_DEPS env var is set and contains 1 txz file url", func() {
		BeforeEach(func() {
			env = append(env, []string{
				fmt.Sprintf("INPUT_DEPS=%s/dep3.txz", testAssetsURL),
			}...)
		})

		It("downloads and extracts the txz dependency to $HOME", func() {
			Eventually(session).Should(Exit(0))

			Expect(filepath.Join(testHomeDir, "dep3")).To(BeADirectory())
			Expect(filepath.Join(testHomeDir, "dep3", "bin", "cake")).To(BeAnExistingFile())
		})
	})

	When("the INPUT_DEPS env var is set and contains a mixture of tgz and txz file urls", func() {
		BeforeEach(func() {
			env = append(env, []string{
				fmt.Sprintf("INPUT_DEPS=%s/dep1.tgz,%s/dep2.tar.gz,%s/dep3.txz", testAssetsURL, testAssetsURL, testAssetsURL),
			}...)
		})

		It("downloads and extracts all dependencies to $HOME", func() {
			Eventually(session).Should(Exit(0))

			Expect(filepath.Join(testHomeDir, "dep1")).To(BeADirectory())
			Expect(filepath.Join(testHomeDir, "dep1", "bin", "cake")).To(BeAnExistingFile())

			Expect(filepath.Join(testHomeDir, "dep2")).To(BeADirectory())
			Expect(filepath.Join(testHomeDir, "dep2", "bin", "cake")).To(BeAnExistingFile())

			Expect(filepath.Join(testHomeDir, "dep3")).To(BeADirectory())
			Expect(filepath.Join(testHomeDir, "dep3", "bin", "cake")).To(BeAnExistingFile())
		})
	})

	When("the INPUT_DEPS env var is set and contains a url to an unsupported filetype", func() {
		BeforeEach(func() {
			env = append(env, []string{
				"INPUT_DEPS=https://github.com",
			}...)
		})

		It("exits with an exit code of 1", func() {
			Eventually(session).Should(Exit(1))
		})

		It("prints an informative error message to stdout", func() {
			Eventually(session.Err).Should(Say("ERROR: unable to unarchive file at 'https://github.com', ensure it is a supported filetype"))
		})
	})
})
