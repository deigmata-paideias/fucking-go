.PHONY: help lint lint-fix

help: ## 显示帮助信息
	@echo "可用的 make 命令："
	@echo "  make lint      - 运行 golangci-lint 检查代码"
	@echo "  make lint-fix  - 运行 golangci-lint 并自动修复问题"

lint: ## 运行 golangci-lint
	@for dir in $$(find . -name "*.go" -type f -exec dirname {} \; | sort -u); do \
		echo "Linting $$dir..."; \
		(cd $$dir && golangci-lint run --timeout=5m) || true; \
	done

lint-fix: ## 运行 golangci-lint 并自动修复
	@for dir in $$(find . -name "*.go" -type f -exec dirname {} \; | sort -u); do \
		echo "Linting and fixing $$dir..."; \
		(cd $$dir && golangci-lint run --fix --timeout=5m) || true; \
	done
