package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/sahilm/fuzzy"
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
		os.Exit(1)
	}
}

func cmdCd(args []string) {
	bestProj, found := fuzzyFindProjectBase(args[0])
	if !found {
		os.Exit(1)
	}

	userHomeDir, _ := os.UserHomeDir()
	destFolder := userHomeDir + "/src/github.com/" + bestProj
	fmt.Println(destFolder)
}

func fuzzyFindProjectBase(projName string) (bestDir string, found bool) {
	userHomeDir, _ := os.UserHomeDir()
	baseDir := userHomeDir + "/src/github.com/"
	dirEntries, err := os.ReadDir(baseDir)
	var allDirs []string
	if err != nil {
		return "", false
	}

	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() {
			continue
		}
		subDirEntries, _ := os.ReadDir(baseDir + dirEntry.Name())
		for _, subDirEntry := range subDirEntries {
			allDirs = append(allDirs, dirEntry.Name()+"/"+subDirEntry.Name())
		}
	}

	results := fuzzy.Find(projName, allDirs)
	if results.Len() < 1 {
		return "", false
	}
	return results[0].Str, true
}
