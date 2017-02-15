# Main Makefile for ripe-atlas
#
# Copyright 2016 © by Ollivier Robert
#

.PATH= cmd/atlas:.
GOBIN=   ${GOPATH}/bin

SRCS= common.go config.go keys.go measurements.go probes.go types.go \
    measurement_subr.go \
	cmd/atlas/atlas.go cmd/atlas/cmd_probes.go cmd/atlas/cmd_measures.go \
	cmd/atlas/cmd_dns.go cmd/atlas/cmd_http.go cmd/atlas/cmd_ip.go \
	cmd/atlas/cmd_ntp.go cmd/atlas/cmd_ping.go cmd/atlas/cmd_sslcert.go \
	cmd/atlas/cmd_traceroute.go cmd/atlas/cmd_keys.go

OPTS=	-ldflags="-s -w" -v

all: atlas

atlas: ${SRCS}
	go build ${OPTS} ./cmd/...
	go test ./...

test:
	go test -v ./...

install:
	go install -v

clean:
	go clean -v

push:
	git push --all
	git push --tags
	git push --all upstream
	git push --tags upstream
