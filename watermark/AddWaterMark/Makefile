# this is commnet
CROSS=

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BUILD_TIME=$(shell date +"%Y%m%d%H%M%S")

DEBUGFLAGS=--race
RELEASEFLAGS=

# when you run make, it means make all, and all depends on test and build.
.PHONY: all
all: debug

.PHONY: debug
debug: clean
debug: FLAG=$(DEBUGFLAGS)
debug: AddWaterMark

.PHONY: release
release: clean
release: FLAG=$(RELEASEFLAGS)
release: AddWaterMark

GLOBAL_VERSION = 0.0.4
GLOBAL_MOD = github.com/MDGSF/watermark

GLOBAL_FLAGS = "-X $(GLOBAL_MOD).Version=$(GLOBAL_VERSION) \
	-X $(GLOBAL_MOD).BuildTime=$(BUILD_TIME)"

AddWaterMark:
	CC=gcc \
	GOOS=linux \
	GOARCH=amd64 \
	$(GOBUILD) \
		-ldflags $(GLOBAL_FLAGS) \
		-o ./bin/$@ ./main/main.go
	strip ./bin/$@
	cp ./bin/$@ ./$@

AddWaterMark_win:
	GOOS=windows \
	GOARCH=amd64 \
	$(GOBUILD) \
		-ldflags $(GLOBAL_FLAGS) \
		-o ./bin/AddWaterMark_$(GLOBAL_VERSION).exe ./main/main.go

.PHONY: clean
clean:
	$(GOCLEAN)
	-rm -rf AddWaterMark
	-rm -rf *.jpg
	-rm -rf ./bin
	-find . -name "*_marked.jpg" -type f -print -exec rm -rf {} \; 
	-find . -name "*_marked.png" -type f -print -exec rm -rf {} \; 
