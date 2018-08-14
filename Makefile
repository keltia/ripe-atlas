# Main Makefile for ripe-atlas
#
# Copyright 2016 © by Ollivier Robert
#

.PATH= cmd/atlas:.
GO=		go
GOBIN=  ${GOPATH}/bin

SRCS= common.go credits.go keys.go measurements.go probes.go types.go \
    measurement_subr.go errors.go client.go version.go \
	cmd/atlas/atlas.go cmd/atlas/cmd_probes.go cmd/atlas/cmd_measurements.go \
	cmd/atlas/cmd_dns.go cmd/atlas/cmd_http.go cmd/atlas/cmd_ip.go \
	cmd/atlas/cmd_ntp.go cmd/atlas/cmd_ping.go cmd/atlas/cmd_sslcert.go \
	cmd/atlas/cmd_traceroute.go cmd/atlas/cmd_keys.go cmd/atlas/cmd_results.go \
	cmd/atlas/config.go cmd/atlas/cmd_credits.go cmd/atlas/flags.go cmd/atlas/utils.go

USRC=	 cmd/atlas/config_unix.go
WSRC=	cmd/atlas/config_windows.go

BIN=	atlas
EXE=	${BIN}.exe
XTRAS=	contrib/* README.md

OPTS=	-ldflags="-s -w" -v

all: ${BIN}

windows:  ${EXE}

${BIN}: ${SRCS} ${USRC}
	${GO} build ${OPTS} ./cmd/...

${EXE}: ${SRCS} ${WSRC}
	GOOS=windows ${GO} build ${OPTS} ./cmd/...

test:
	${GO} test -v ./...

bench:
	${GO} test -bench=. -benchmem ./...

lint:
	gometalinter ./...

install: ${BIN}
	${GO} install ${OPT} -v ./cmd/...

pkg: ${BIN} ${EXE}
	-/bin/mkdir pkg
	tar cvf - ${BIN} ${XTRAS} | xz > pkg/${BIN}.tar.xz
	zip -r pkg/${BIN}.zip ${EXE} ${XTRAS}

clean:
	${GO} clean -v ./...
	rm -f ${BIN} ${EXE} pkg/*

push:
	git push --all
	git push --tags
