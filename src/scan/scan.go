package scan

import (
	"io/ioutil"
	"os"
	"path"
	"../utils"
)
var user = "7aske"
var email = "ntasic7@gmail.com"
var excluded = []string{"_test"}
func recursiveGetGit(p string, r *[]string) {
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
					recursiveGetGit(absPath, r)
				}
			}
		}
	} else {
		panic("path: invalid path " + p)
	}
}

func GetGit(p string) []string {
	var repos []string
	recursiveGetGit(p, &repos)
	return repos
}