#!make

.DEFAULT_GOAL := build

XLB_BIN=fsm-xlb
CCM_BIN=fsm-xlb-ccm
CLI_BIN=fsmxlbc

nproc ?= 2

clang ?= $(shell lsb_release -r | cut -f2 | sed s/22.04/clang-13/ | sed s/20.04/clang-10/)

TARGETS := linux/amd64 linux/arm64
DIST_DIRS := find * -type d -exec
GOPATH = $(shell go env GOPATH)
GOBIN  = $(GOPATH)/bin
GOX    = go run github.com/mitchellh/gox
SHA256 = sha256sum
ifeq ($(shell uname),Darwin)
	SHA256 = shasum -a 256
endif

VERSION ?= dev
BUILD_DATE ?= $$(date +%Y-%m-%d-%H:%M)
GIT_SHA=$$(git rev-parse HEAD)
BUILD_DATE_VAR := github.com/flomesh-io/fsmxlb/pkg/version.BuildDate
BUILD_VERSION_VAR := github.com/flomesh-io/fsmxlb/pkg/version.Version
BUILD_GITCOMMIT_VAR := github.com/flomesh-io/fsmxlb/pkg/version.GitCommit

GO_LDFLAGS ?= "-X $(BUILD_DATE_VAR)=$(BUILD_DATE) -X $(BUILD_VERSION_VAR)=$(VERSION) -X $(BUILD_GITCOMMIT_VAR)=$(GIT_SHA) -s -w"

.PHONY: depends
depends:
	@apt -y update
	@arch=$(arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/) && echo $arch && if [ "$arch" = "arm64" ] ; then apt install -y gcc-multilib-arm-linux-gnueabihf; else apt update && apt install -y gcc-multilib;fi
	@apt install -y $(clang) llvm libelf-dev libpcap-dev net-tools elfutils dwarves git libbsd-dev bridge-utils wget arping unzip build-essential bison flex sudo iproute2 pkg-config tcpdump iputils-ping keepalived curl bash-completion
	@apt -y autoremove
	@cp bpf/utils/mkxlb_bpffs.sh /usr/local/sbin/mkxlb_bpffs
	@mkdir -p /opt/fsmxlb/cert
	@cp api/certification/* /opt/fsmxlb/cert/
	@if [ ! -f /usr/local/sbin/bpftool ]; then git clone --recurse-submodules https://github.com/libbpf/bpftool.git && cd bpftool/src/ && make clean && make -j $(nproc) && cp -f ./bpftool /usr/local/sbin/bpftool && cd - && rm -fr bpftool; fi
	@if [ ! -f /usr/local/sbin/ntc ]; then wget https://github.com/cybwan/iproute2/archive/refs/heads/fsmxlb.zip && unzip fsmxlb.zip && cd iproute2-fsmxlb/libbpf/src/ && mkdir build && DESTDIR=build make install && cd - && cd iproute2-fsmxlb/ && export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:`pwd`/libbpf/src/ && LIBBPF_FORCE=on LIBBPF_DIR=`pwd`/libbpf/src/build ./configure && make && cp -f tc/tc /usr/local/sbin/ntc && cd - && cd iproute2-fsmxlb/libbpf/src/ && make install && cd - && rm -fr fsmxlb.zip iproute2-fsmxlb; fi
	@if [ ! -f /usr/lib64/libbpf.so.0.4.0 ]; then cd bpf && make && make install && cd -; fi
	@if [ ! -f /usr/sbin/gobgp ]; then arch=${shell arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/} && echo https://github.com/osrg/gobgp/releases/download/v3.5.0/gobgp_3.5.0_linux_$${arch}.tar.gz.tar.gz && wget https://github.com/osrg/gobgp/releases/download/v3.5.0/gobgp_3.5.0_linux_$${arch}.tar.gz && tar -xzf gobgp_3.5.0_linux_$${arch}.tar.gz && rm gobgp_3.5.0_linux_$${arch}.tar.gz LICENSE README.md && mv gobgp* /usr/sbin/; fi
	@if [ ! -f /usr/local/go/bin/go ]; then arch=${shell arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/} && echo https://go.dev/dl/go1.19.linux-$${arch}.tar.gz && wget https://go.dev/dl/go1.19.linux-$${arch}.tar.gz && tar -xzf go1.19.linux-$${arch}.tar.gz --directory /usr/local/ && rm go1.19.linux-$${arch}.tar.gz;echo please export PATH=\$${PATH}:/usr/local/go/bin; fi
	@ kver=${shell uname -r | cut -d"-" -f1} && echo $${kver} && if $$(dpkg --compare-versions $${kver}  "lt" "5.14"); then ukv=${shell sudo apt list linux-image-5.*-generic 2>&1 | grep ^linux | cut -d '-' -f 3,4 | sort -rV | head -n1} && sudo apt install -y linux-modules-$${ukv}-generic linux-headers-$${ukv}-generic linux-image-$${ukv}-generic; fi

subsys:
	mkdir -p /opt/fsmxlb/cert
	cp bpf/utils/mkxlb_bpffs.sh /usr/local/sbin/mkxlb_bpffs
	cp api/certification/* /opt/fsmxlb/cert/
	cd bpf && make

subsys-clean:
	cd bpf && make clean

build:
	@go build -o bin/${XLB_BIN} -ldflags ${GO_LDFLAGS} ./cmd/fsmxlbd/fsmxlbd.go

build-cli:
	@go build -o bin/${CLI_BIN} -ldflags ${GO_LDFLAGS} ./cmd/fsmxlbc/fsmxlbc.go

build-ccm:
	@go build -o bin/${CCM_BIN} -ldflags ${GO_LDFLAGS} ./cmd/fsmccm/fsmccm.go

clean:
	go clean
	rm -rf bin

run:
	@./bin/$(XLB_BIN)

stop:
	@sudo killall fsmxlb >> /dev/null 2>&1 || true
	@sudo ip link set dev flb0 xdp off >> /dev/null 2>&1 || true
	@sudo ip link set dev flb0 xdpgeneric off >> /dev/null 2>&1 || true
	@sudo ntc filter del dev flb0 ingress >> /dev/null 2>&1 || true
	@sudo umount /opt/fsmxlb/dp >> /dev/null 2>&1 || true
	@sudo rm -fr /opt/fsmxlb/dp/bpf >> /dev/null 2>&1 || true
	@sudo ip link del flb0 >> /dev/null 2>&1 || true

lint:
	golangci-lint run --enable-all

.PHONY: go-vet
go-vet:
	go vet ./...

.PHONY: go-fmt
go-fmt:
	go fmt ./...

.PHONY: go-mod-tidy
go-mod-tidy:
	./scripts/go-mod-tidy.sh

# -------------------------------------------
#  release targets below
# -------------------------------------------

.PHONY: build-cross
build-cross:
	GO111MODULE=on CGO_ENABLED=0 $(GOX) -ldflags $(GO_LDFLAGS) -parallel=5 -output="_dist/{{.OS}}-{{.Arch}}/$(CLI_BIN)" -osarch='$(TARGETS)' ./cmd/fsmxlbc

.PHONY: dist
dist:
	( \
		cd _dist && \
		$(DIST_DIRS) cp ../LICENSE {} \; && \
		$(DIST_DIRS) cp ../README.md {} \; && \
		$(DIST_DIRS) tar -zcf fsmxlb-cli-${VERSION}-{}.tar.gz {} \; && \
		$(DIST_DIRS) zip -r fsmxlb-cli-${VERSION}-{}.zip {} \; && \
		$(SHA256) fsmxlb-* > sha256sums.txt \
	)

.PHONY: release-artifacts
release-artifacts: build-cross dist