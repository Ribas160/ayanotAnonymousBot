.SILENT:
build:
	mkdir -p bin
	mkdir -p run
	go build -o ./bin/ayanotAnonymousBot ./cmd/main/
	go build -o bot ./cmd/remote