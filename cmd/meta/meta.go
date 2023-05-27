package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func main() {
	fmt.Println(`
Here is another one of my earliest Go programs. It explores metadata stuff like the runtime. It's a work in progress.`, "\n")

	_, file, _, _ := runtime.Caller(0)
	rootDir := filepath.Dir(file)
	fmt.Println("This Go program is running in the directory: ", rootDir)
}
