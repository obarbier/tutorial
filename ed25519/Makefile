GO_BIN=/usr/local/go/bin/go

build:
	rm -rf output
	mkdir -p output
	$(GO_BIN) build -race main.go
	mv main output/