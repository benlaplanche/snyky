package cmd_test

import (
	"bytes"
	"io/ioutil"

	. "github.com/benlaplanche/snyky/cmd"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Snyky Test Command", func() {
	Context("Runs successfully when", func() {
		It("Discovers a file and policy pack", func() {
			cmd := NewTestCmd()
			b := bytes.NewBufferString("")
			cmd.SetOut(b)
			cmd.SetArgs([]string{"test"})

			// Workaround from https://github.com/onsi/ginkgo/issues/285
			// origArgs := os.Args[:]
			// os.Args = os.Args[:1]
			// cmd.Execute()
			// os.Args = origArgs[:]
			cmd.Execute()

			out, err := ioutil.ReadAll(b)

			Expect(string(out)).To(Equal("test called"))
			Expect(err).To(BeNil())

		})
	})

})
