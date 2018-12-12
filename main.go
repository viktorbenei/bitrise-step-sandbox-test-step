package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("All env vars:")
	fmt.Println("=========================")
	fmt.Println(strings.Join(os.Environ(), "\n---\n"))
	fmt.Println("=========================")
	os.Exit(0)
}
