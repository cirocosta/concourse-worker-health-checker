SRC = $(shell find . -name "*.go" -and -not -name "*_test.go")
OUT = ./main.out

build: $(OUT)

image:
	docker build -t cirocosta/concourse-worker-health-checker .

fmt:
	go fmt ./...

test:
	go test -v ./...

run: $(OUT)
	./$^


$(OUT): $(SRC)
	go build -v -i -o $@ .
