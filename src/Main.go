package main

import (
	"./date"
	"./git"
	"./print"
	"bufio"
	"fmt"
	"github.com/go-ini/ini"
	"os"
	"os/user"
	"path"
	"strings"
	"time"
)

var rootDir = "/home/nik/Documents/CODE"
var gitRepos []string
var userName = ""
var email = ""

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	var configPath = path.Join(usr.HomeDir, ".gitstats-go")
	if _, err := os.Stat(configPath); err != nil {
		f, _ := os.Create(configPath)
		_, _ = f.Write([]byte("username=\nemail=\ndir=\n"))
		_ = f.Close()
	}
	config, err := ini.Load(configPath)
	userName = config.Section("").Key("username").String()
	email = config.Section("").Key("email").String()
	rootDir = config.Section("").Key("dir").String()
	if userName == "" {
		scan := bufio.NewReader(os.Stdin)
		fmt.Print("enter username: ")
		userName, _ = scan.ReadString('\n')
		userName = strings.Trim(userName, "\n")
		config.Section("").Key("username").SetValue(userName)
	}
	if email == "" {
		scan := bufio.NewReader(os.Stdin)
		fmt.Print("enter email: ")
		email, _ = scan.ReadString('\n')
		email = strings.Trim(email, "\n")
		config.Section("").Key("email").SetValue(email)
	}
	if rootDir == "" {
		scan := bufio.NewReader(os.Stdin)
		fmt.Print("enter dir: ")
		rootDir, _ = scan.ReadString('\n')
		rootDir = strings.Trim(rootDir, "\n")
		config.Section("").Key("dir").SetValue(rootDir)
	}
	if _, err := os.Stat(rootDir); err != nil {
		panic(err)
	}
	_ = config.SaveTo(configPath)

	limit := time.Now().AddDate(0, -4, - time.Now().Day())
	now := time.Now()
	history := date.GenerateHistory(limit)
	gitRepos = git.GetGit(rootDir, userName)
	for _, repo := range gitRepos {
		git.ProcessRepo(repo, email, limit, history)
	}
	keys := date.GetSortedHistoryKeys(&history)
	total := 0
	for i := range keys {
		total += history[i]
	}
	print.OutputHistory(history, keys, now)
}
