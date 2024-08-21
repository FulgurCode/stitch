build:
	go build -o ./bin/cloth-shop-website ./cmd/main.go

run: build
	./bin/cloth-shop-website
