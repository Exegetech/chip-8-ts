.PHONY: fmt test build

fmt:
	gofmt -w .

build:
	tinygo build -o ../web/public/cpu.wasm -target wasm ./main.go

test:
	go test test/instructions/*.go


