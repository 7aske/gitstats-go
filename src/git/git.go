package git

import (
	"../date"
	"../utils"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"io/ioutil"
	"os"
	"path"
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
			daysAgo := date.DaysSince(c.Author.When)
			m[daysAgo]++
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

var excluded = []string{"_test"}

func recursiveGetGit(p string, r *[]string, user string) {
	var dir []os.FileInfo
	var err error
	if dir, err = ioutil.ReadDir(p); err != nil {
		panic("path: cannot read " + p)
	}
	if _, err = os.Stat(p); err == nil {
		for _, f := range dir {
			absPath := path.Join(p, f.Name())
			if f.IsDir() && !utils.Contains(f.Name(), excluded) {
				if f.Name() == ".git" {
					if fi, err := os.Open(path.Join(absPath, "config")); err == nil {
						if utils.HasUser(user, fi) {
							*r = append(*r, p)
						}
					} else {
						panic("path: error opening " + path.Join(absPath, ".git", "config"))
					}
					break
				} else {
					recursiveGetGit(absPath, r, user)
				}
			}
		}
	} else {
		panic("path: invalid path " + p)
	}
}

func GetGit(p string, user string) []string {
	var repos []string
	recursiveGetGit(p, &repos, user)
	return repos
}
