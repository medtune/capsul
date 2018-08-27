build-base:
	@echo building base capsul library ...
	docker build \
		-t medtune/capsul:latest \
		-f Dockerfile \
		.

	docker tag medtune/capsul:latest medtune/capsul:v0.0.2

capsul: build-base
capsul-base: build-base


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
		-t medtune/capsul:mura_irn_v2 \
		-f build/capsules/mura_irn_v2.Dockerfile \
		.


# build mura main image
build-mura: build-mura-irn-v2
	@echo taging mura-irn-v2 with mura
	docker tag medtune/capsul:mura_irn_v2 medtune/capsul:mura


# Build mura 
build-mura-cam:
	@echo building model capsul mura-cam ...
	docker build \
		-t medtune/capsul:mura-cam \
		-f build/capsules/mura-cam.Dockerfile \
		.


build-chexray:
	@echo building model capsul chexray ...
	docker build \
		-t medtune/capsul:chexray \
		-f build/capsules/chexray.Dockerfile \
		.
