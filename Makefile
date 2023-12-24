build:
	@go build -o ./bin/dexcli *.go

run: build
	./bin/dexcli