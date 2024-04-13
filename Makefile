# Build the custom cas -o specifies the output folder
build:
	@go build -o bin/custom_cas

# Run the custom cas
run: build
	@./bin/custom_cas

# Run the tests
tests:
	@go test ./... -v
