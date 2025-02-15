
# Variables
BINARY_NAME=pretty-chord-progression
SOURCE_FILES=main.go

# Default target
all: build run

# Build target
build:
	@echo "Building the binary..."
	@go build -o ./bin/${BINARY_NAME} $(SOURCE_FILES)

# Run target
run: 
	@echo "Running the program..."
	@./bin/${BINARY_NAME} "2" "./test/input.txt" "./test/output.txt"

# Clean target
clean:
	@echo "Cleaning up..."
	@rm -f ./bin/${BINARY_NAME}

# Help target
help:
	@echo "Usage:"
	@echo "  make         - Builds and runs the program"
	@echo "  make build   - Builds the program"
	@echo "  make run     - Runs the program"
	@echo "  make clean   - Removes the binary"
	@echo "  make help    - Displays this help message"