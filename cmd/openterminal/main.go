package main

import (
	"fmt"
	"os"
	"strings"

	"OpenTerminal/pkg/runtime"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: openterminal run <command>")
		os.Exit(1)
	}

	if os.Args[1] != "run" {
		fmt.Println("unknown command")
		os.Exit(1)
	}

	command := strings.Join(os.Args[2:], " ")

	runner := runtime.NewRunner()

	result := runner.Run(command)

	fmt.Println("Status:", result.Status)

	if result.Output != "" {
		fmt.Println("Output:")
		fmt.Println(result.Output)
	}

	if result.Error != "" {
		fmt.Println("Error:")
		fmt.Println(result.Error)
	}

	if result.Status != runtime.Success {
		os.Exit(1)
	}
}
