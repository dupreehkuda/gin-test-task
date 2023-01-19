.PHONY: build
build:
	docker build

.PHONY: run
run:
	docker run -e SERVICE_ADDRESS='8080'
