package main

import (
	"fmt"
	"log"
	"os/exec"
)

/*
Execute external processes from a Go program.
*/
func main() {
	// First, let's try executing a built-in command like 'ls'. A useful pre-check is to verify that the command exists
	// on the PATH. If it does, we'll execute it.
	execLs()

	// Next, let's do something more involved. Let's execute the 'docker' command (if it exists!). Specifically, we'll
	// check for the version of 'docker' that's installed on the system.
	execDocker()
}

func execLs() {
	// Check if the command exists on the system.
	_, err := exec.LookPath("ls")
	if err != nil {
		log.Fatal("The 'ls' command does not exist on the PATH")
	}

	// Execute the command.
	output, err := exec.Command("ls").Output()
	if err != nil {
		log.Fatal(err)
	}

	// Print the output.
	fmt.Println("The 'ls' command found the following files in the current working directory:")
	fmt.Println(string(output))
}

func execDocker() {
	_, err := exec.LookPath("docker")
	if err != nil {
		log.Fatal("The 'docker' command does not exist on the PATH")
	}

	output, err := exec.Command("docker", "version", "-f", "{{.Server.Version}}").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The 'docker' command reports the following version:")
	fmt.Println(string(output))
}
