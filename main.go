package main

import (
	"fmt"
	"flag"
	"os/exec"
)

func main() {
	fmt.Println("gutils")
	commitMessage := flag.String("commit", "initial commit", "A commit message - defaults to 'initial commit'")
	gitRemote := flag.String("remote", "origin", "The specified upstream git remote to push to - defaults to origin")
	gitBranch := flag.String("branch", "main", "The branch to push to - defaults to main")
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

