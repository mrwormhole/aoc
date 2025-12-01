CXX = clang++
CXXFLAGS = -std=c++17 -Wall -Wextra -O2

# Default day
DAY ?= 1

# Directories
DAY_DIR = day$(DAY)
SRC = $(DAY_DIR)/main.cc
TARGET = $(DAY_DIR)/main

.PHONY: all build clean run run-sample fmt help

# Default target
all: build

# Build specific day
build:
	@if [ ! -f "$(SRC)" ]; then \
		echo "Error: $(SRC) not found"; \
		exit 1; \
	fi
	$(CXX) $(CXXFLAGS) $(SRC) -o $(TARGET)
	@echo "Built day$(DAY)"

# Run with input.txt
run: build
	@cd $(DAY_DIR) && ./main input.txt

# Run with sample.txt
run-sample: build
	@cd $(DAY_DIR) && ./main sample.txt

# Clean all days
clean:
	@find . -name "main" -type f -delete
	@echo "Cleaned all days"

# Format all days
fmt:
	@find . -name "main.cc" -type f -exec clang-format -i {} \;
	@echo "Formatted all days"

# Help
help:
	@echo "Usage:"
	@echo "  make [DAY=N]             - Build day N (default: 1)"
	@echo "  make run [DAY=N]         - Build and run day N with input.txt"
	@echo "  make run-sample [DAY=N]  - Build and run day N with sample.txt"
	@echo "  make fmt                 - Format all days with clang-format"
	@echo "  make clean       	      - Clean all days"
	@echo ""
	@echo "Examples:"
	@echo "  make DAY=5               - Build day 5"
	@echo "  make run DAY=5           - Run day 5 with input.txt"
	@echo "  make run-sample DAY=3    - Run day 3 with sample.txt"
