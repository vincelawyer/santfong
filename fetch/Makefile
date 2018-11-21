# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)

ALL_GO_SOURCES=$(shell /bin/sh -c "find *.go | grep -v _test.go")

crawl: fmt
	@echo "\033[92mCrawling page...\033[0m"
	go run $(ALL_GO_SOURCES)

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	go fmt *.go

install:
	go get -u github.com/PuerkitoBio/goquery
	go get -u github.com/djimenez/iconv-go
