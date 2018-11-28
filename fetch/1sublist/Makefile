# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
ifndef GOROOT
	export GOROOT=$(realpath ../../../go)
	export GOPATH=$(realpath .)
	export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)
endif

ALL_GO_SOURCES=$(shell /bin/sh -c "find *.go | grep -v _test.go")

ulmark: fmt checkout_content
	go test -v $(ALL_GO_SOURCES) ulmark_test.go

image: fmt checkout_content
	go test -v $(ALL_GO_SOURCES) imagerst_test.go

ogimage: fmt checkout_content
	go test -v $(ALL_GO_SOURCES) ogimage_test.go

sublist: fmt checkout_content
	@echo "\033[92mCreate product list and final product rst files...\033[0m"
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
