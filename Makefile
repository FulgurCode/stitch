build:
	@templ generate
	@go build -o ./bin/benster-website ./cmd/main.go

run: build
	@./bin/benster-website
