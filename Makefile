SRC = $(shell find . -name "*.go" -and -not -name "*_test.go")
OUT = ./health-checker.out


build: $(OUT)

fmt:
	go fmt ./...

vet:
	go vet ./...

test: vet
	go test -v ./...

run: $(OUT)
	./$^


$(OUT): $(SRC)
	go build -v -i -o $@ .
