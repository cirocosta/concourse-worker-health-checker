SRC = $(shell find . -name "*.go" -and -not -name "*_test.go")
OUT = ./health-checker.out

image:
	docker build -t a .

build: $(OUT)

fmt:
	go fmt ./...

test:
	go test -v ./...

run: $(OUT)
	./$^


$(OUT): $(SRC)
	go build -v -i -o $@ .
