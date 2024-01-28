build:
	@templ generate
	@go build -o tmp/main .

live:
	@air

run:
	@templ generate
	@go run main.go