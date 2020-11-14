build:
	docker build -t rufous .
run-docker:
	docker run -it -p 8081:8081 rufous

run:
	make build
	make run-docker