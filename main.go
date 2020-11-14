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
	fmt.Printf("Test results are...", out)
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	// log.Fatal(err)
	// }
	// fmt.Printf("The date is %s\n", out)
}
