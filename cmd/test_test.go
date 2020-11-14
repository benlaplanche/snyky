package cmd_test

import (
	"bytes"
	"io/ioutil"
	"time"

	. "github.com/benlaplanche/snyky/cmd"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Snyky Test Command", func() {
	var Timestamp time.Time

	BeforeEach(func() {
		// gotta be a better way of forcing a timestamp
		timelayout := "2020-11-14T19:48:15.978286Z"
		timestring := "2020-11-14T19:48:15.978286Z"
		Timestamp, _ = time.Parse(timelayout, timestring)
	})

	Context("Runs successfully when", func() {
		It("Is not passed any flags", func() {
			cmd := NewTestCmd(Timestamp)
			b := bytes.NewBufferString("")
			cmd.SetOut(b)
			cmd.SetArgs([]string{"test"})

			cmd.Execute()
			_, err := ioutil.ReadAll(b)

			Expect(err).To(BeNil())
		})

		It("Discovers a file and policy pack", func() {
			cmd := NewTestCmd(Timestamp)
			b := bytes.NewBufferString("")
			cmd.SetOut(b)
			cmd.SetArgs([]string{"--source", "../examples/terraform.tf", "--packs", "../packs/terraform"})

			cmd.Execute()

			out, err := ioutil.ReadAll(b)

			expected_result, _ := ioutil.ReadFile("test_output.json")

			Expect(string(out)).To(MatchJSON(expected_result))
			Expect(err).To(BeNil())

		})
	})

})
