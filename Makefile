.PHONY: help lint lint-fix

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
