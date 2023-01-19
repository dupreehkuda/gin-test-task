.PHONY: build
build:
	docker build -t gin-task --build-arg addr=:8083 .

.PHONY: run
run:
	docker run -p 8080:8080 gin-task
