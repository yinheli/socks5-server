CMD = socks5-server
TARGET = dist
GIT_HASH := $(shell (git show-ref --head --hash=8 2> /dev/null || echo 00000000) | head -n1)

all: pack

deps:
	# install deps
	@hash dep > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go get -u github.com/golang/dep/cmd/dep; \
	fi
	@dep ensure -v

fmt:
	# gofmt code
	gofmt -s -l -w cmd util

$(CMD):
	# go build
	go build -o $(TARGET)/$@ -ldflags="-X main.build=$(GIT_HASH)" ./cmd/main.go

pack: $(CMD)
	@echo "done"

clean:
	@rm -rf ./$(TARGET)

.PHONY: $(CMD) all deps clean pack fmt