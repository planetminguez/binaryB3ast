# Telemetry Variables
BINARY_NAME=goforge
SRC=main.go
BUILD_DIR=build
LDFLAGS=-ldflags="-s -w" # Strips debug symbols to drastically reduce binary size

.PHONY: all help mac linux windows build-all test clean

all: help

help:
	@echo "💎 GOFORGE MULTI-OS COMPILATION CONTROLLER 💎"
	@echo "Individual OS Compilation Targets:"
	@echo "  make mac         - Compile optimized binary for macOS (Universal Universal/Fat Architecture)"
	@echo "  make linux       - Compile optimized binary for Linux (x86_64)"
	@echo "  make windows     - Compile optimized binary for Windows (x86_64 .exe)"
	@echo ""
	@echo "Automation & Diagnostic Tools:"
	@echo "  make build-all   - Simultaneously build assets for all three target systems"
	@echo "  make test        - Automatically deploy an absolute string test sandbox locally"
	@echo "  make clean       - Purge structural binary builds and clean cache matrices"

mac:
	@echo "🍏 [MAC OS] Synthesizing Mach-O structural framework layers..."
	@mkdir -p $(BUILD_DIR)
	# Compile separate architecture slices
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)_mac_x86 $(SRC)
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)_mac_arm $(SRC)
	# Fuse slices together natively using lipo if running on a Mac, otherwise keep both
	@if command -v lipo >/dev/null 2>&1; then \
		lipo -create -output $(BUILD_DIR)/$(BINARY_NAME)_mac $(BUILD_DIR)/$(BINARY_NAME)_mac_x86 $(BUILD_DIR)/$(BINARY_NAME)_mac_arm; \
		rm $(BUILD_DIR)/$(BINARY_NAME)_mac_x86 $(BUILD_DIR)/$(BINARY_NAME)_mac_arm; \
		echo "🍎 [SUCCESS] Mach-O Universal/Fat binary generated at: $(BUILD_DIR)/$(BINARY_NAME)_mac"; \
	else \
		echo "🍎 [SUCCESS] Slices generated: $(BUILD_DIR)/$(BINARY_NAME)_mac_x86 & $(BUILD_DIR)/$(BINARY_NAME)_mac_arm"; \
	fi

linux:
	@echo "🐧 [LINUX OS] Fabricating ELF standard executable slice..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)_linux_x64 $(SRC)
	@echo "🐧 [SUCCESS] ELF x64 framework binary generated at: $(BUILD_DIR)/$(BINARY_NAME)_linux_x64"

windows:
	@echo "🪟 [WINDOWS OS] Compiling Portable Executable (PE) runtime engine..."
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)_win_x64.exe $(SRC)
	@echo "🪟 [SUCCESS] PE x64 execution structure generated at: $(BUILD_DIR)/$(BINARY_NAME)_win_x64.exe"

build-all: mac linux windows
	@echo "\n💎 [MATRIX CLEAN] Structural assets compiled across all designated targets."

test:
	@echo "🧪 Synthesizing safe dummy data test string sandbox..."
	@echo "Made by: argp, axi0mx, danyl931, jaywalker, kirb, littlelailo, ni" > sandbox
	@echo "⚙️  Compiling localized framework context..."
	@go build -o $(BINARY_NAME) $(SRC)
	@echo "✅ Setup complete. Booting GoForge local sandbox environment..."
	./$(BINARY_NAME) sandbox

clean:
	@echo "🧹 Purging workspace artifacts and build cache frameworks..."
	rm -f $(BINARY_NAME) sandbox sandbox_patched sandbox_indexed
	rm -rf $(BUILD_DIR)
	@go clean
	@echo "✨ Workspace environment normalized."
