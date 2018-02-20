package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"time"

	"github.com/rhino1998/writhub/tiny"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func main() {

	repo, err := git.PlainInit(os.Args[1], false)
	if err != nil {
		repo, err = git.PlainOpen(os.Args[1])
		if err != nil {
			log.Fatalf("%v\n", err)
		}
	}

	wt, err := repo.Worktree()
	if err != nil {
		log.Fatalf("WorkTree: %v\n", err)
	}

	start, err := time.Parse("Jan 2 2006", os.Args[2])
	if err != nil {
		log.Fatalf("Invalid time: %v\n", err)
	}
	start = start.Add(time.Hour * 12)

	img := image.NewAlpha(image.Rect(0, 0, 42, 7))
	width, height := img.Rect.Dx(), img.Rect.Dy()

	tiny.Font.DrawString(img, 0, 1, os.Args[3], color.Black)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if img.Pix[y*width+x] == 255 {
				wt.Commit("WritHub Pixel ART", &git.CommitOptions{
					Author: &object.Signature{
						Name:  os.Args[4],
						Email: os.Args[5],
						When:  start.Add(time.Hour * 24 * time.Duration(x*height+y)),
					},
				})
			}
		}
	}

	ref, err := repo.Head()
	if err != nil {
		log.Fatalf("Head: %v\n", err)
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
