package main

import (
	"fmt"
	"os"
	"syscall"
)

/*
This is a true 'exec' example which replaces the process image (a Go program) with the process image of another program
(like 'ls'). This is a very neat feature of Unix-like operating systems. The 'exec' family of system calls is really
useful for creating things like launcher programs.
*/
func main() {
	fmt.Println("Hello from a Go program. This process will be 'morphed' from a Go program into an execution of the 'ls' program. The process stays the same, but we can consider that 'image' (memory) is replaced.")

	command := "/bin/ls"
	// Oddly, the first argument is the command itself. Why is that?
	args := []string{"ls", "-lh"}

	// Replace the current process with the new process
	err := syscall.Exec(command, args, os.Environ())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// This line will not be executed if syscall.Exec was successful. Pretty neat, but also pretty strange. This means
	// that things like deferred functions will not be executed. Be careful!
	fmt.Println("This will not be printed.")
}
