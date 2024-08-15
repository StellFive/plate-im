package main

import (
	"fmt"
	"os"
	"plato/cmd"
)

func main() {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		fmt.Println("GOPATH environment variable is not set")
	} else {
			fmt.Println("GOPATH is:", gopath)
	}

	cmd.Execute()
}