build-base:
	@echo building base capsul library ...
	docker build \
		-t medtune/capsul:latest \
		-f Dockerfile \
		.

	docker tag medtune/capsul:latest medtune/capsul:v0.0.2

build-mnist:
	@echo building model capsul mnist ...
	docker build \
		-t medtune/capsul:mnist \
		-f build/capsules/mnist.Dockerfile \
		.

build-inception:
	@echo building model capsul inception ...
	docker build \
		-t medtune/capsul:inception \
		-f build/capsules/inception.Dockerfile \
		.

build-mura:
	@echo building model capsul mura ...
	docker build \
		-t medtune/capsul:mura \
		-f build/capsules/mura.Dockerfile \
		.

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