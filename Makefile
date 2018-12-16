BINARY := findpath

PLATFORMS := linux darwin
os = $(word 1, $@)

go.mod:
		go mod init github.com/miry/ev3dev

.PHONY: setup
setup: go.mod

.PHONY: $(PLATFORMS)
$(PLATFORMS):
		mkdir -p release
		GOOS=$(os) GOARCH=arm GOARM=5 go build -o release/$(BINARY)-v0.1.0-$(os)-arm5 ./cmd/$(BINARY)/

.PHONY: release
release: linux
