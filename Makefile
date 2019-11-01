APP_BIN :=htpasswd-init
DATE_VERSION := $(shell date +%Y%m%d-%H%M)
GIT_VERSION := $(shell git rev-parse --short HEAD)
GIT_DATE_VERSION := $(GIT_VERSION)-$(DATE_VERSION)

all: clean release

release:
	CGO_ENABLED=0 go build -o $(APP_BIN) -ldflags "-s -w -X main.Version=$(GIT_DATE_VERSION)"

clean:
	-rm -f $(APP_BIN)