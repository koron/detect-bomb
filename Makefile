PROJECT=detectbomb
VERSION=v0.1.1

GOVERSION=$(shell go version)
GOOS=$(word 1,$(subst /, ,$(lastword $(GOVERSION))))
GOARCH=$(word 2,$(subst /, ,$(lastword $(GOVERSION))))

ARCNAME=$(PROJECT)-$(VERSION)-$(GOOS)-$(GOARCH)
RELDIR=$(PROJECT)-$(GOOS)-$(GOARCH)

release:
	rm -rf tmp/$(RELDIR)
	mkdir -p tmp/$(RELDIR)
	go clean
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags='-X main.version=$(VERSION)'
	cp $(PROJECT)$(SUFFIX_EXE) tmp/$(RELDIR)/
	tar czf tmp/$(ARCNAME).tar.gz -C tmp $(RELDIR)
	go clean

release-all:
	@$(MAKE) release GOOS=windows GOARCH=amd64 SUFFIX_EXE=.exe
	@$(MAKE) release GOOS=windows GOARCH=386   SUFFIX_EXE=.exe
	@$(MAKE) release GOOS=linux   GOARCH=amd64
	@$(MAKE) release GOOS=linux   GOARCH=386
	@$(MAKE) release GOOS=darwin  GOARCH=amd64

.PHONY: release
