.PHONY: build install clean

BINARY_NAME = archi
INSTALL_PATH = $(HOME)/bin

build:
	go build -o $(BINARY_NAME)

install: build
	mkdir -p $(INSTALL_PATH)
	mv $(BINARY_NAME) $(INSTALL_PATH)
	@echo "✅ Installation completed in $(INSTALL_PATH)"
	@if ! grep -q "$(INSTALL_PATH)" ~/.bashrc; then \
		echo "Adding $(INSTALL_PATH) to PATH..."; \
		echo 'export PATH=$PATH:$(INSTALL_PATH)' >> ~/.bashrc; \
		echo "✅ PATH updated in ~/.bashrc"; \
	else \
		echo "$(INSTALL_PATH) already in PATH"; \
	fi
	@echo ""
	@echo "🎉 Setup complete! Restart your terminal or run:"
	@echo "source ~/.bashrc"
	@echo ""
	@echo "Then you can use: $(BINARY_NAME) create"

clean:
	rm -f $(BINARY_NAME)