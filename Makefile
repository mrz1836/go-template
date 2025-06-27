# Common makefile commands & variables between projects
include .make/common.mk

# Common Golang makefile commands & variables between projects
include .make/go.mk

## Set default repository details if not provided
REPO_NAME ?= go-template
REPO_OWNER ?= mrz1836

# ------------------------------------------------------------------------------
# install-template — Kick-start a fresh copy of go-template
# ------------------------------------------------------------------------------
# Usage:
#   make install-template owner=myorg repo=my-lib
#
# What it does:
#   1. Replaces every occurrence of:
#        - "mrz1836/go-template" → "${owner}/${repo}"
#        - "go-template"        → "${repo}"
#        - "mrz1836"            → "${owner}"
#   2. Removes the default social-share image so you can drop in your own.
#   3. Prints a friendly reminder of the next steps.
#
# Notes:
#   * Works on macOS (BSD make + BSD/GNU sed) and Linux (GNU make/sed).
#   * Skips the .git directory to avoid mangling Git internals.
# ------------------------------------------------------------------------------

# Detect GNU sed (gsed) on macOS, fall back to the system sed otherwise
SED := $(shell command -v gsed 2>/dev/null || command -v sed)

.PHONY: install-template
install-template: ## Kick-start a fresh copy of go-template
	@set -e; \
	if [ -z "$(owner)" ]; then \
	  echo "✋ Please supply \"owner\" (e.g. owner=myorg)"; \
	  exit 1; \
	fi; \
	if [ -z "$(repo)" ]; then \
	  echo "✋ Please supply \"repo\" (e.g. repo=my-lib)"; \
	  exit 1; \
	fi; \
	echo "🔧 Initializing template for $(owner)/$(repo)..."; \
	find . -type f \( -name "*" \) -not -path "./.git/*" -print0 | \
	xargs -0 $(SED) -i '' \\
	  -e 's#mrz1836/go-template#$(owner)/$(repo)#g' \\
	  -e 's#go-template#$(repo)#g' \\
	  -e 's#mrz1836#$(owner)#g'; \\
	if [ -f .github/IMAGES/go-share-image.png ]; then \\
	  rm .github/IMAGES/go-share-image.png; \\
	  echo "🗑️ Removed default social-share image (.github/IMAGES/go-share-image.png)"; \\
	fi; \\
	echo "✅ Done! Review the changes, add your own share image, and commit the initial version.";
