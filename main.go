package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// fmt.Println("SUCCESS")
	filename := "terraform.tf"
	args := []string{"test", string(filename), "--policy=packs/terraform"}
	out, err := exec.Command("conftest", args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}
