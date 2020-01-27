tools: export GOBIN := $(PWD)/bin
tools:
	go install honnef.co/go/tools/cmd/staticcheck

check: export PATH := $(PWD)/bin:$(PATH)
check:
	staticcheck ./...
