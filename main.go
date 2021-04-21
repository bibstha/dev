package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Things to implement
// dev clone project_name
// dev cd
func main() {
	cmd := os.Args[1]
	cmdArgs := os.Args[2:]
	if cmd == "clone" {
		cmdClone(cmdArgs)
	} else if cmd == "cd" {
		cmdCd(cmdArgs)
	} else if cmd == "echo" {
		fmt.Printf("Dev binary echoing: %v\n", cmdArgs)
	}
}

func cmdClone(args []string) {
	repoBase := "github.com/" + args[0]
	repoUrl := "https://" + repoBase
	userHomeDir, _ := os.UserHomeDir()
	destFolder := userHomeDir + "/src/" + repoBase
	cmd := exec.Command("git", "clone", repoUrl, destFolder)
	output, err := cmd.CombinedOutput()
	fmt.Printf("%s\n", output)
	if err != nil {
		fmt.Printf("Failed to clone: %v", err)
		// os.Exit(0)
	}
}

func cmdCd(args []string) {
	repoBase := "github.com/" + args[0]
	userHomeDir, _ := os.UserHomeDir()
	destFolder := userHomeDir + "/src/" + repoBase
	fmt.Println(destFolder)
}
