# Altalune Justfile

# Variables
binary_name := "altalune"
cmd_dir := "."
build_dir := "./bin"

# Default recipe
default: help

# Run the frontend dev server
build-fe:
    @echo "Building Frontend..."
    @cd frontend && bun run build
    @cd ..

# Build the proxy server binary
build:
    @just build-fe
    @echo "Building {{binary_name}}..."
    @mkdir -p {{build_dir}}
    @go build -o {{build_dir}}/{{binary_name}} {{cmd_dir}}
    @echo "Binary built: {{build_dir}}/{{binary_name}}"

# Run the proxy server
run-be:
    @echo "Starting Backend Server..."
    @go run {{cmd_dir}} --verbose

# Run the frontend dev server
run-fe:
    @echo "Starting Frontend Dev Server..."
    @cd frontend && bun run dev

# Run both frontend and backend concurrently
run:
    @echo "Starting both Frontend and Backend servers..."
    just run-fe
    # @just run-be & just run-fe

# Clean build artifacts
clean:
    @echo "Cleaning build artifacts..."
    @rm -rf {{build_dir}}
    @echo "Clean complete"

# Run Go tests
test-go:
    @echo "Running Go tests..."
    @go test -v ./...

# Download dependencies
deps:
    @echo "Downloading Go dependencies..."
    @go mod tidy
    @echo "Installing frontend dependencies..."
    @cd frontend && bun install

# Show CLI help
help-cli:
    @echo "Showing CLI help..."
    @go run {{cmd_dir}} --help

# Show help
help:
    @echo "Altalune"
    @echo "=================="
    @echo ""
    @echo "Available commands:"
    @echo ""
    @echo "  build           Build the binary"
    @echo "  run             Run both frontend and backend concurrently"
    @echo "  run-be          Run only the backend server"
    @echo "  run-fe          Run only the frontend dev server"
    @echo "  deps            Download all dependencies (Go + Frontend)"
    @echo "  clean           Clean build artifacts"
    @echo "  help-cli        Show CLI help"
    @echo "  help            Show this help message"
    @echo ""
