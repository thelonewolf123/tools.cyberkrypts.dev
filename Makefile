build:
	@templ generate
	@go build -o bin/api .

live:
	@air --build.cmd "make build" --build.bin "./bin/api"

run:
	@templ generate
	@go run main.go