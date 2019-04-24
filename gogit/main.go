package main

import (
	"fmt"
	"os"

	git "gopkg.in/src-d/go-git.v4"
)

func main() {
	_, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL:      "https://github.com/BurntSushi/toml.git",
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Println(err)
	}
	/*
		re, err := git.PlainOpen("/tmp/foo")
		if err != nil {
			fmt.Println(err)
		}
		tr, err := re.Worktree()
		if err != nil {
			fmt.Println(err)
		}
		//err = tr.Pull(new(git.PullOptions))
		err = tr.Reset(&git.ResetOptions{Mode: 1})
		if err != nil {
			fmt.Println("------", err)
		}
	*/
}
