package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// check for correct format type
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER(optional)]\n EX: go run . --output=<fileName.txt> something standard")
		os.Exit(1)
	}

	banner := "banner/standard.txt"
	// banner selection
	if len(os.Args) == 4 {
		banner = os.Args[3]
		switch banner {
		case "standard":
			banner = "banner/standard.txt"
		case "shadow":
			banner = "banner/shadow.txt"
		case "thinkertoy":
			banner = "banner/thinkertoy.txt"
		default:
			fmt.Println("Use a valid banner template")
			os.Exit(2)
		}
	}

	// read file
	file, err := os.ReadFile(banner)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	if len(os.Args) == 2 { // if only string argument
		fmt.Println(GenerateArt(file, os.Args[1]))
		os.Exit(0)
	}

	// write to output file
	output := GenerateArt(file, os.Args[2])
	flag := os.Args[1]
	if strings.Contains(flag, "--output=") {
		name := strings.Split(flag, "=")[1]
		os.WriteFile(name, []byte(output), 0o644)
		fmt.Println("Writing to file completed")
	}
}
