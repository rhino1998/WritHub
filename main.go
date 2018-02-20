package main

import (
	"fmt"
	"log"
	"time"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func main() {

	repo, err := git.PlainInit("tmp", false)
	if err != nil {
		repo, err = git.PlainOpen("tmp")
		if err != nil {
			log.Fatalf("%v\n", err)
		}
	}

	wt, err := repo.Worktree()
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	start := time.Date(2017, time.September, 2, 12, 0, 0, 0, time.UTC)

	for i := 0; i < 10; i++ {
		wt.Commit("WritHub Pixel ART", &git.CommitOptions{
			Author: &object.Signature{
				When: start.Add(time.Hour * 24 * time.Duration(i)),
			},
		})
	}

	ref, err := repo.Head()
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	cIter, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	// ... just iterates over the commits, printing it
	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c)
		return nil
	})

	if err != nil {
		log.Fatalf("%v\n", err)
	}

}
