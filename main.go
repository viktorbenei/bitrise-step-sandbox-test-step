package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("All env vars:")
	fmt.Println("=========================")
	fmt.Println(strings.Join(os.Environ(), "\n---\n"))
	fmt.Println("=========================")

	cmdLog, err := exec.Command("pwd").CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to run command, error: %#v | output: %s", err, cmdLog)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Printf("'pwd' command output: (%s)", cmdLog)
	fmt.Println()

	os.Exit(0)
}
