build:
	@templ generate
	@go build -o ./bin/stitch ./cmd/main.go

run: build
	@./bin/stitch
