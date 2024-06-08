build:# Define the name of the binary
BINARY_NAME=./bin/magic

# Default target
all: build

# Build the binary
build:
	go build -o $(BINARY_NAME) ./cmd/main.go
run: build
	$(BINARY_NAME)

clean:
	rm -rf $(BINARY_NAME)
