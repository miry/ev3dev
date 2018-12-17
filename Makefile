BINARY := findpath
VERSION := v0.1.0

PLATFORMS := linux darwin
os = $(word 1, $@)

go.mod:
		go mod init github.com/miry/ev3dev

.PHONY: setup
setup: go.mod

.PHONY: $(PLATFORMS)
$(PLATFORMS):
		mkdir -p release
		GOOS=$(os) GOARCH=arm GOARM=5 go build -o release/$(BINARY)-$(VERSION)-$(os)-arm5 ./cmd/$(BINARY)/

.PHONY: deploy
deploy: linux
		scp release/$(BINARY)-$(VERSION)-linux-arm5 robot@ev3dev.local:~/bin/

.PHONY: release
release: linux deploy

.PHONY: run
run: release
		ssh -it robot@ev3dev.local ./bin/$(BINARY)-$(VERSION)-linux-arm5
