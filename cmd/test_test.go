package cmd_test

import (
	"bytes"
	"fmt"
	"io/ioutil"

	. "github.com/benlaplanche/snyky/cmd"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Snyky Test Command", func() {
	Context("Runs successfully when", func() {
		It("Is not passed any flags", func() {
			cmd := NewTestCmd()
			b := bytes.NewBufferString("")
			cmd.SetOut(b)
			cmd.SetArgs([]string{"test"})

			cmd.Execute()
			out, err := ioutil.ReadAll(b)

			Expect(string(out)).To(Equal("Error executing conftest"))
			Expect(err).To(BeNil())
		})

		FIt("Discovers a file and policy pack", func() {
			cmd := NewTestCmd()
			b := bytes.NewBufferString("")
			cmd.SetOut(b)
			cmd.SetArgs([]string{"--source", "../terraform.tf", "--packs", "../packs/terraform"})

			cmd.Execute()

			out, err := ioutil.ReadAll(b)
			fmt.Println(string(out))

			expected_result, _ := ioutil.ReadFile("test_output.json")

			Expect(string(out)).To(MatchJSON(expected_result))
			Expect(err).To(BeNil())

		})
	})

})
