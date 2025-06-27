# Common makefile commands & variables between projects
include .make/common.mk

# Common Golang makefile commands & variables between projects
include .make/go.mk

## Set default repository details if not provided
REPO_NAME ?= go-template
REPO_OWNER ?= mrz1836

.PHONY: install-all-contributors
install-all-contributors: ## Installs all contributors locally
	@echo "installing all-contributors cli tool..."
	@yarn global add all-contributors-cli

.PHONY: update-contributors
update-contributors: ## Regenerates the contributors html/list
	@echo "generating contributor html..."
	@all-contributors generate
