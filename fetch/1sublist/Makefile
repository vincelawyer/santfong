# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
ifndef GOROOT
	export GOROOT=$(realpath ../../../go)
	export GOPATH=$(realpath .)
	export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)
endif

ALL_GO_SOURCES=$(shell /bin/sh -c "find *.go | grep -v _test.go")

sublist: fmt
	@echo "\033[92mFetching remote images...\033[0m"
	git checkout ../../content/pages/
	go run $(ALL_GO_SOURCES)

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	go fmt *.go

test: fmt
	go test -v

run:
	cd ../.. ; make

checkout_content:
	git checkout ../../content/

install:
	go get -u github.com/mozillazg/go-pinyin
	go get -u github.com/PuerkitoBio/goquery
	go get -u github.com/djimenez/iconv-go
