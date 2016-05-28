# Main Makefile for ripe-atlas
#
# Copyright 2016 © by Ollivier Robert
#

.PATH= cmd/atlas:.
GOBIN=   ${GOPATH}/bin

SRCS= common.go config.go measurements.go probes.go types.go \
	cmd/atlas/atlas.go cmd/atlas/cmd_probes.go

OPTS=	-ldflags="-s -w" -v

all: atlas

atlas: ${SRCS}
	go build ${OPTS}
	go test -v

install:
	go install -v

clean:
	go clean -v

push:
	git push --all
	git push --tags
	git push --all upstream
	git push --tags upstream
