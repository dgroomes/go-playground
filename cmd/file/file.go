package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("Let's read the files in the current working directory")

	currentWorkingDir, _ := os.Getwd()

	log.Println("The current working directory is:", currentWorkingDir)

	file, _ := os.Open(currentWorkingDir)
	names, _ := file.Readdirnames(0)

	log.Println("The directory contains these files:")
	fmt.Println(names)
}
