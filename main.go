package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// fmt.Println("SUCCESS")
	filename := "terraform.tf"
	out, err := exec.Command("conftest", "test", filename, "--policy=packs/terraform").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}
