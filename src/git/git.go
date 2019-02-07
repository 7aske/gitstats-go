package git

import (
	"../date"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"time"
)

func ProcessRepo(p string, email string, limit time.Time, m map[int]int) {
	repo, err := git.PlainOpen(p)
	if err != nil {
		panic(err)
	}
	ref, err := repo.Head()
	if err != nil {
		panic(err)
	}

	iterator, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		panic(err)
	}

	err = iterator.ForEach(func(c *object.Commit) error {
		if c.Author.Email != email {
			return nil
		}
		if limit.Before(c.Author.When) {
			//fmt.Println(limit, c.Author.When)

			daysAgo := date.DaysSince(c.Author.When)
			m[daysAgo]++
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
