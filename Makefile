GOPATH = $(shell go env GOPATH)

run: build
	$(addprefix $(GOPATH)/, bin/server) -log=$(log)

build:
	go install github.com/martins0n/openvpn-traffic-viewer/server

test:
	go test -v ./...