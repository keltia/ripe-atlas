# Main Makefile for ripe-atlas
#
# Copyright 2016 © by Ollivier Robert
#

.PATH= cmd/atlas:.
GOBIN=   ${GOPATH}/bin

SRCS= common.go keys.go measurements.go probes.go types.go \
    measurement_subr.go \
	cmd/atlas/atlas.go cmd/atlas/cmd_probes.go cmd/atlas/cmd_measurements.go \
	cmd/atlas/cmd_dns.go cmd/atlas/cmd_http.go cmd/atlas/cmd_ip.go \
	cmd/atlas/cmd_ntp.go cmd/atlas/cmd_ping.go cmd/atlas/cmd_sslcert.go \
	cmd/atlas/cmd_traceroute.go cmd/atlas/cmd_keys.go cmd/atlas/cmd_results.go \
	cmd/atlas/config.go

USRC=	 cmd/atlas/config_unix.go
WSRC=	cmd/atlas/config_windows.go

BIN=	atlas
EXE=	${BIN}.exe

OPTS=	-ldflags="-s -w" -v

all: ${BIN} ${EXE}

${BIN}: ${SRCS} ${USRC}
	go build ${OPTS} ./cmd/...

${EXE}: ${SRCS} ${WSRC}
	GOOS=windows go build ${OPTS} ./cmd/...

test:
	go test -v ./...

install: ${BIN}
	go install -v ./cmd/...

clean:
	go clean -v ./...
	rm -f ${BIN} ${EXE}

push:
	git push --all
	git push --tags
	git push --all upstream
	git push --tags upstream
