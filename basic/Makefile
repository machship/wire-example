.PHONY: install

build:
	wire

install:
	go install github.com/google/wire/cmd/wire@latest

run:
	go run main.go wire_gen.go event.go message.go greeter.go