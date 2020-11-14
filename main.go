package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// fmt.Println("SUCCESS")
	filename := "terraform.tf"
	args := []string{"test", string(filename), "--policy=packs/terraform"}
	out, err := exec.Command("conftest", args...).Output()
	// fmt.Printf("Test results are...", out)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if out != nil {
		// fmt.Fprintln(os.Stdout, out)
		// fmt.Printf("Test results are...\n %v", out)
		fmt.Print(string(out))
	}
}
