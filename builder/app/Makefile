DOCKER_TAG	:= build_builder

build: Dockerfile
	docker build -t build_builder .

run: build
	docker run -it --rm -p 1234:1234 build_builder
