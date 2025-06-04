.PHONY: build install clean uninstall check-path add-path

BINARY_NAME = archi
INSTALL_PATH = $(HOME)/bin
MODULE_NAME = github.com/thomas-bressel/archi-ts-cli

build:
	@if [ ! -f go.mod ]; then \
		echo "🔧 Initialisation of Go mudule ..."; \
		go mod init $(MODULE_NAME); \
	fi
	@echo "📦 Dependencies installation..."
	go mod tidy
	go build -o $(BINARY_NAME)

install: build
	mkdir -p $(INSTALL_PATH)
	cp $(BINARY_NAME) $(INSTALL_PATH)/
	rm -f $(BINARY_NAME)
	@echo "✅ $(BINARY_NAME) installed in $(INSTALL_PATH)"
	@echo ""
	@if ! echo "$$PATH" | grep -q "$(INSTALL_PATH)"; then \
		echo "⚠️  $(INSTALL_PATH) is not in PATH"; \
		echo "Add this line to your ~/.bashrc :"; \
		echo "export PATH=\"\$$PATH:$(INSTALL_PATH)\""; \
		echo ""; \
		echo "Then execute : source ~/.bashrc"; \
	else \
		echo "✅ $(INSTALL_PATH) is already in le PATH"; \
		echo "Now you can use : $(BINARY_NAME) create"; \
	fi

# Add to PATH safely (only if user explicitly asks)
add-path:
	@if ! grep -q "$(INSTALL_PATH)" ~/.bashrc; then \
		echo 'export PATH="$$PATH:$(INSTALL_PATH)"' >> ~/.bashrc; \
		echo "✅ PATH updated. Run: source ~/.bashrc"; \
	else \
		echo "$(INSTALL_PATH) already in ~/.bashrc"; \
	fi

check-path:
	@if echo "$$PATH" | grep -q "$(INSTALL_PATH)"; then \
		echo "✅ $(INSTALL_PATH) is in PATH"; \
	else \
		echo "❌ $(INSTALL_PATH) is NOT in PATH"; \
	fi

clean:
	rm -f $(BINARY_NAME)

uninstall:
	rm -f $(INSTALL_PATH)/$(BINARY_NAME)
	@echo "🗑️  $(BINARY_NAME) uninstalled"