package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// fmt.Println("SUCCESS")
	filename := "terraform.tf"
	args := []string{"test", string(filename), "--policy=packs/terraform"}
	out, _ := exec.Command("conftest", args...).Output()
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// }

	if out != nil {
		fmt.Print(string(out))
	}
}