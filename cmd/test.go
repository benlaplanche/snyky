/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// Source is the filesource
var Source string
var Packs string

// NewTestCmd exported for testing
func NewTestCmd() *cobra.Command {
	var testCmd = &cobra.Command{
		Use:   "test",
		Short: "Scan your configuration files for rules violations",
		Long: `An experimental Snyk IaC replacement that uses conftest locally to scan your configuration files for rules violations. 
		This uses a concept of packs, which are groups of policies that can be turned on and off.`,
		Run: func(cmd *cobra.Command, args []string) {

			// Setup the files to be scanned
			// Appending this here as it needs to be the 2nd argument
			filename, _ := cmd.Flags().GetString("source")
			args = append(args, filename)

			// we need to run `$conftest test` so need to ensure `test` is the 1st argument
			args = append([]string{"test"}, args...)

			// Setup the policy packs to run
			packs, _ := cmd.Flags().GetString("packs")
			if packs == "" {
				packs = "packs"
			}
			args = append(args, fmt.Sprintf("--policy=%s", packs))

			// add in our enforced output flag
			args = append(args, "--output=json")

			out, err := exec.Command("conftest", args...).Output()

			if string(out) == "" && err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), "Error executing conftest")
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), string(out))
			}
		},
	}
	// var Source string
	testCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
	testCmd.Flags().StringVarP(&Packs, "packs", "p", "", "Path to the policy packs to be run")
	testCmd.MarkFlagRequired("source")
	return testCmd
}

func init() {
	testCmd := NewTestCmd()

	// testCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
	// testCmd.MarkFlagRequired("source")

	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// testCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
}
