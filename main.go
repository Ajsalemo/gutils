package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

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
	// git commit -m [commit message]
	gitCommit := exec.Command("git", "commit", "-m", commitMessage)
	gitCommit.Stdout = os.Stdout
	gitCommit.Stderr = os.Stderr

	err := gitCommit.Run()
	// Check if we push a commit where nothing is going to be commited to the branch
	// This is returned as an error - but we want to clean up the output of this and present this in less of a critical manner
	// Therefor we return from those messages as exit(0) - which is fine and mimics normal Git behavior
	// if strings.Contains(string(err.Error()), "nothing to commit, working tree clean") {
	// 	// If the bool is true, log out the message in full to stdout still
	// 	fmt.Println(string(err.Error()))
	// 	// Call exit(0)
	// 	os.Exit(0)
	// }

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
