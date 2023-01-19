.PHONY: build
build:
	docker build -t gin-task .

.PHONY: run
run:
	docker run -p 8080:8080 gin-task
