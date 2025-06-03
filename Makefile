.PHONY: build install clean uninstall check-path

BINARY_NAME = archi
INSTALL_PATH = $(HOME)/bin

build:
	go build -o $(BINARY_NAME)

install: build
	mkdir -p $(INSTALL_PATH)
	cp $(BINARY_NAME) $(INSTALL_PATH)/
	rm -f $(BINARY_NAME)
	@echo "✅ $(BINARY_NAME) installed to $(INSTALL_PATH)"
	@echo ""
	@if ! echo "$$PATH" | grep -q "$(INSTALL_PATH)"; then \
		echo "⚠️  $(INSTALL_PATH) is not in your PATH"; \
		echo "Add this line to your ~/.bashrc:"; \
		echo "export PATH=\"\$$PATH:$(INSTALL_PATH)\""; \
		echo ""; \
		echo "Then run: source ~/.bashrc"; \
	else \
		echo "✅ $(INSTALL_PATH) is already in PATH"; \
		echo "You can now use: $(BINARY_NAME) create"; \
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