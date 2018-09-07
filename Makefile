PROJECT=beta-platform
OS_TYPE=$(shell uname -a)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GITCOMMIT=$(shell git rev-parse HEAD)
BUILDDATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
MAJOR=0
MINOR=0
PATCH=3
REVISION=alpha
VERSION=v$(MAJOR).$(MINOR).$(PATCH)
GOVERSION=1.11
LONGVERSION=v$(MAJOR).$(MINOR).$(PATCH)-$(REVISION)
CWD=$(shell pwd)
VPATH=github.com/medtune/capsul/pkg
PROJECTPATH=$(CWD)
AUTHORS=Hilaly.Mohammed-Amine/El.bouchti.Alaa
OWNERS=$(AUTHORS)
LICENSETYPE=Apache-v2.0
LICENSEURL=https://raw.githubusercontent.com/medtune/capsul/master/LICENSE.txt


CONTAINER_ENGINE=docker

GOCMD=go
GOVERSION=$(GOCMD) version
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

build: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v $(BINARY_FILE)

test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)


test-clean:
	rm -f test/testdata/*_cam.png


# build base container image 
# does not compile binaries
build-base:
	@echo building base capsul library ...
	docker build \
		-t medtune/capsul:base \
		-f build/base.Dockerfile \
		.


# build capsul package and tag version/latest
capsul: build-base
	docker tag medtune/capsul:base medtune/capsul:$(VERSION)
	docker tag medtune/capsul:base medtune/capsul:latest
	

# alias capsul
capsul-base: capsul


# build capsul command line tool
build-cmd:
	@echo building base capsul library ...
	docker build \
		-t medtune/capsul:cmd \
		-f build/cmd.Dockerfile \
		.


# build capsul cmd & tag version/latest
capsul-cmd: build-cmd
	docker tag medtune/capsul:cmd medtune/capsul:cmd-$(VERSION)
	docker tag medtune/capsul:cmd medtune/capsul:cmd-latest


# Build mnist tf server image
build-mnist:
	@echo building model capsul mnist ...
	docker build \
		-t medtune/capsul:mnist \
		-f build/capsules/mnist.Dockerfile \
		.


# Build inception tf server image
build-inception:
	@echo building model capsul inception ...
	docker build \
		-t medtune/capsul:inception \
		-f build/capsules/inception.Dockerfile \
		.


# Build mura inception restnet v2 tf server image
build-mura-irn-v2:
	@echo building model capsul mura inception resnet v2 ...
	docker build \
		-t medtune/capsul:mura-irn-v2 \
		-f build/capsules/mura_irn_v2.Dockerfile \
		.


# Build mura mobilenet v2 tf server image
build-mura-mn-v2:
	@echo building model capsul mura mobilenet v2 ...
	docker build \
		-t medtune/capsul:mura-mn-v2 \
		-f build/capsules/mura_mobilenet_v2.Dockerfile \
		.


# Build mura mobile net grad cam customized server
build-mura-mn-v2-cam:
	@echo building grad cam server for mura mobilenet v2
	docker build \
		-t medtune/capsul:mura-mn-v2-cam \
		-f build/csflask/mura_mn_v2_cam.Dockerfile \
		.


# build mura main image
build-mura: build-mura-mn-v2
	@echo taging mura-mn-v2 with mura
	docker tag medtune/capsul:mura-mn-v2 medtune/capsul:mura


# Build mura 
build-mura-cam: build-mura-mn-v2-cam
	@echo building model capsul mura-cam ...
	docker tag medtune/capsul:mura-mn-v2-cam medtune/capsul:mura-cam


# Build chexray mobilenet v2
build-chexray-mn-v2:
	docker build \
		-t medtune/capsul:chexray-mn-v2 \
		-f build/capsules/chexray_mobilenet_v2.Dockerfile \
		.


# Build chexray mobilenet v2 grad cam
build-chexray-mn-v2-cam:
	docker build \
		-t medtune/capsul:chexray-mn-v2 \
		-f build/csflask/chexray_mobilenet_v2_cam.Dockerfile \
		.


# Build chexray densenet 121
build-chexray-dn-121:
	@echo building model capsul chexray densenet 121 ...
	docker build \
		-t medtune/capsul:chexray-dn-121 \
		-f build/capsules/chexray_densenet_121.Dockerfile \
		.


# Build chexray
build-chexray-pp:
	docker build \
		-t medtune/capsul:chexray-pp-helper \
		-f build/csflask/chexray_pp.Dockerfile \
		.


# Build csflask
build-csflask: build-mura-mn-v2-cam \
	build-chexray-pp \
	build-chexray-mn-v2-cam


# Build capsules
build-capsules: build-mnist \
	build-inception \
	build-mura-mn-v2 \
	build-mura-mn-v2-cam \
	build-mura-irn-v2 \
	build-chexray-dn-121 \
	build-chexray-mn-v2

