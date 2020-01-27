VERSION := $(shell git describe --always --dirty --tags)
LDFLAGS := -ldflags '-X main.VersionString=$(VERSION)'

build-%:
	go build $(LDFLAGS) -o build/$* ./cmd/$*

# Enable auto-completion for all available main packages.
COMMANDS := $(patsubst cmd/%/.,%,$(wildcard cmd/*/.))
$(addprefix build-,$(COMMANDS)):
