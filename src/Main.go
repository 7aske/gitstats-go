package main

import (
	"./date"
	"./git"
	"./print"
	"./scan"
	"time"
)

var rootDir = "/home/nik/Documents/CODE"
var gitRepos []string
var user = "7aske"
var email = "ntasic7@gmail.com"
func main() {
	limit := time.Now().AddDate(0,-4,-3)
	history := date.GenerateHistory(limit)
	gitRepos = scan.GetGit(rootDir)
	for _, repo := range gitRepos {
		git.ProcessRepo(repo, email, limit, history)
	}
	keys := date.GetSortedHistoryKeys(&history)
	total := 0
	for i := range keys {
		total += history[i]
	}
	print.PrintMap(history, keys)
}



