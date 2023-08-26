package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

func executeGitAdd() {
	fmt.Println("\n")
	// git add .
	gitAdd := exec.Command("git", "add", ".")
	gitAdd.Stdout = os.Stdout
	gitAdd.Stderr = os.Stderr

	err := gitAdd.Run()
	if err != nil {
		log.Fatalf("executeGitAdd() failed: %s", err)
	}
}

func executeGitCommit(commitMessage string) {
	// git commit [commit message]
	gitCommit := exec.Command("git", "commit", "-m", commitMessage)
	out, e := gitCommit.Output()
	// Check if gitCommit.Output() returns an error - return types from this function is []byte and error
	if e != nil {
		log.Fatalf("executeGitCommit() [gitCommit.Output()] failed: %s", e)
	}

	if strings.Contains(string(out), "nothing to commit, working tree clean") {
		return
	}

	err := gitCommit.Run()
	if err != nil {
		log.Fatalf("executeGitCommit() failed: %s", err)
	}
}

func executeGitPush(gitRemote string, gitBranch string) {
	// git push [remote] [branch]
	gitPush := exec.Command("git", "push", gitRemote, gitBranch)
	gitPush.Stdout = os.Stdout
	gitPush.Stderr = os.Stderr

	err := gitPush.Run()
	if err != nil {
		log.Fatalf("executeGitPush() failed: %s", err)
	}
}

func main() {
	// Documentation reference: https://github.com/common-nighthawk/go-figure
	myFigure := figure.NewColorFigure("gutils", "", "blue", true)
	myFigure.Print()
	// Define a flag for the commit message to git
	commitMessage := flag.String("c", "initial commit", "A commit message - defaults to 'initial commit'")
	// Define a flag for the git remote
	gitRemote := flag.String("r", "origin", "The specified upstream git remote to push to - defaults to origin")
	// Define a flag for the branch to be pushed to
	gitBranch := flag.String("b", "main", "The branch to push to - defaults to main")
	// Parse all input
	flag.Parse()

	// flags package expects pointers back to the original values
	// git add .
	executeGitAdd()
	// git commit [commit message]
	executeGitCommit(*commitMessage)
	// git push [remote] [branch]
	executeGitPush(*gitRemote, *gitBranch)
}
