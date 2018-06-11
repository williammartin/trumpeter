package integration_test

import (
	"fmt"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Trumpeter", func() {

	var (
		trumpeterCmd *exec.Cmd

		session *gexec.Session
	)

	BeforeEach(func() {
		trumpeterCmd = exec.Command(trumpeterPath)
		trumpeterCmd.Env = append(trumpeterCmd.Env, twitterSecrets()...)
	})

	JustBeforeEach(func() {
		session = execBin(trumpeterCmd)
	})

	It("exits with a successful code", func() {
		Eventually(session).Should(gexec.Exit(0))
	})

	It("returns one of the best quotes, everybody knows", func() {
		Eventually(session).Should(gbytes.Say(".* - President Trump"))
	})
})

func twitterSecrets() []string {
	return []string{
		fmt.Sprintf("CONSUMER_KEY=%s", os.Getenv("CONSUMER_KEY")),
		fmt.Sprintf("CONSUMER_SECRET=%s", os.Getenv("CONSUMER_SECRET")),
		fmt.Sprintf("ACCESS_TOKEN=%s", os.Getenv("ACCESS_TOKEN")),
		fmt.Sprintf("ACCESS_TOKEN_SECRET=%s", os.Getenv("ACCESS_TOKEN_SECRET")),
	}
}
