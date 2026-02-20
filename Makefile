.PHONY: help lint lint-fix install-tools

lint: ## 运行 golangci-lint
	@for dir in $$(find . -name "*.go" -type f -exec dirname {} \; | sort -u); do \
		echo "Linting $$dir..."; \
		(cd $$dir && golangci-lint run --timeout=5m *.go) || true; \
	done

lint-fix: ## 运行 golangci-lint 并自动修复
	@for dir in $$(find . -name "*.go" -type f -exec dirname {} \; | sort -u); do \
		echo "Linting and fixing $$dir..."; \
		(cd $$dir && golangci-lint run --fix --timeout=5m *.go) || true; \
	done

install-tools:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		echo "golangci-lint is already installed"; \
		golangci-lint --version; \
	else \
		echo "Installing golangci-lint..."; \
	@curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.10.1
	@golangci-lint --version
	@echo "golangci-lint installed successfully"
