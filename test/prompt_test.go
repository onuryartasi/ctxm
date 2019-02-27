package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"testing"

	expect "github.com/Netflix/go-expect"
	"github.com/onuryartasi/context-manager/util"
)

func TestChangeContext(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	os.Setenv("KUBECONFIG", fmt.Sprintf("%s/mocks/config1:%s/mocks/config2", wd, wd))
	config := util.GetRawConfig()
	contexts := util.GetContexts(config)

	sort.Strings(contexts)
	fmt.Println(contexts)
	c, err := expect.NewConsole(expect.WithStdout(os.Stdout))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	input := contexts[0]
	os.Chdir(fmt.Sprintf("%s/..", wd))
	cmd := exec.Command("./context-manager")
	cmd.Stdin = c.Tty()
	cmd.Stdout = c.Tty()
	cmd.Stderr = c.Tty()

	go func() {
		c.ExpectString("Choose a context:")
		c.SendLine(input)
		c.ExpectEOF()
		config2 := util.GetRawConfig()
		if input != config2.CurrentContext {
			t.Errorf("Doesnt match %s, %s", input, config2.CurrentContext)
		}
	}()

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
