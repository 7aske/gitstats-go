NAME=gitstats
OUT=out
MAIN=src/Main.go

default: build

dep:
	go get gopkg.in/src-d/go-git.v4
	go get github.com/go-ini/ini

.PHONY: build

build: $(MAIN)
	go build -o $(OUT)/$(NAME) $(MAIN)

run:
	go run $(MAIN)