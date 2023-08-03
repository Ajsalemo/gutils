package main

import (
	"fmt"
	"flag"
	"os/exec"
)

func main() {
	fmt.Println("gutils")
	// Define a flag for the commit message to git
	commitMessage := flag.String("c", "initial commit", "A commit message - defaults to 'initial commit'")
	// Define a flag for the git remote
	gitRemote := flag.String("r", "origin", "The specified upstream git remote to push to - defaults to origin")
	// Define a flag for the branch to be pushed to
	gitBranch := flag.String("b", "main", "The branch to push to - defaults to main")
	// Parse all input
	flag.Parse()

	// flags package expects pointers back to the original values
	fmt.Println(*commitMessage)
	fmt.Println(*gitRemote)
	fmt.Println(*gitBranch)

	// git add .
	gitAdd := exec.Command("git", "add", ".")
	// git commit [commit message] 
	gitCommit := exec.Command("git", "commit", "-m", *commitMessage)
	// git push [remote] [branch]
	gitPush := exec.Command("git", "push", *gitRemote, *gitBranch)

	gitAddStdOut, err := gitAdd.Output()

	if err != nil {
        fmt.Println(err.Error())
        return
    }

	fmt.Println(string(gitAddStdOut))

	gitCommitStdOut, err := gitCommit.Output()

	if err != nil {
        fmt.Println(err.Error())
        return
    }

	fmt.Println(string(gitCommitStdOut))

	gitPushStdOut, err := gitPush.Output()

	if err != nil {
        fmt.Println(err.Error())
        return
    }

	fmt.Println(string(gitPushStdOut))
}

