# Main Makefile for ripe-atlas
#
# Copyright 2016 © by Ollivier Robert
#

.PATH= cmd/atlas:.
GOBIN=   ${GOPATH}/bin

SRCS= common.go credits.go keys.go measurements.go probes.go types.go \
    measurement_subr.go errors.go client.go version.go \
	cmd/atlas/atlas.go cmd/atlas/cmd_probes.go cmd/atlas/cmd_measurements.go \
	cmd/atlas/cmd_dns.go cmd/atlas/cmd_http.go cmd/atlas/cmd_ip.go \
	cmd/atlas/cmd_ntp.go cmd/atlas/cmd_ping.go cmd/atlas/cmd_sslcert.go \
	cmd/atlas/cmd_traceroute.go cmd/atlas/cmd_keys.go cmd/atlas/cmd_results.go \
	cmd/atlas/config.go cmd/atlas/cmd_credits.go cmd/atlas/utils.go

USRC=	 cmd/atlas/config_unix.go
WSRC=	cmd/atlas/config_windows.go

BIN=	atlas
EXE=	${BIN}.exe

OPTS=	-ldflags="-s -w" -v

all: checks ${BIN}

check:
	@V=`go version|cut -d' ' -f 3| sed 's/^go//'` && \
	if test "x$$V" \< "x1.8" ; then \
		echo "You must have go 1.8+"; \
		exit 1; \
	fi

windows:  ${EXE}

${BIN}: ${SRCS} ${USRC}
	go build ${OPTS} ./cmd/...

${EXE}: ${SRCS} ${WSRC}
	GOOS=windows go build ${OPTS} ./cmd/...

test: check
	go test -v ./...

install: check ${BIN}
	go install -v ./cmd/...

clean:
	go clean -v ./...
	rm -f ${BIN} ${EXE}

push:
	git push --all
	git push --tags
