package main

import (
	"fmt"
	"flag"
)

func main() {
	fmt.Println("gutils")
	commitMessage := flag.String("commit", "initial commit", "A commit message - defaults to 'initial commit'")
	gitRemote := flag.String("remote", "origin", "The specified upstream git remote to push to - defaults to origin")
	gitBranch := flag.String("branch", "main", "The branch to push to - defaults to main")
	// flags package expects pointers back to the original values
	fmt.Println(*commitMessage)
	fmt.Println(*gitRemote)
	fmt.Println(*gitBranch)
}