# ---------------------------------------------------------------------------
# .make/temp.mk  –  One-shot “install-template” helper
# ---------------------------------------------------------------------------
# This file removes itself after first use and erases its own include line
# from the project-root Makefile, so you never see it again. 🎩✨
# ---------------------------------------------------------------------------

# Detect GNU sed (gsed) on macOS; fall back to system sed elsewhere
SED          := $(shell command -v gsed 2>/dev/null || command -v sed)

ifeq ($(shell $(SED) --version 2>/dev/null | grep -c 'GNU'),1)
  SED_INPLACE := -i       # GNU sed: -i[EXT]
else
  SED_INPLACE := -i ''    # BSD sed: -i '' …
endif

.PHONY: install-template
install-template: ## Kick-start a fresh copy of go-template (run once!)
	@set -e; \
	# --------------------------------------------------------------------
	# 0. Self-destruct prerequisites
	# --------------------------------------------------------------------
	echo "🧹 Cleaning up temp boot-strapper…"; \
	if [ -f .make/temp.mk ]; then \
	  rm .make/temp.mk; \
	  echo "   • Removed .make/temp.mk"; \
	fi; \
	if [ -f Makefile ]; then \
	  env LC_ALL=C $(SED) $(SED_INPLACE) '/include[[:space:]].*\.make\/temp\.mk/d' Makefile; \
	  echo "   • Stripped 'include .make/temp.mk' from Makefile"; \
	fi; \
	\
	# --------------------------------------------------------------------
	# 1. Validate CLI input
	# --------------------------------------------------------------------
	if [ -z "$(owner)" ]; then \
	  echo '✋ Supply \"owner\" (e.g. owner=my_org)'; exit 1; \
	fi; \
	if [ -z "$(repo)" ]; then \
	  echo '✋ Supply \"repo\"  (e.g. repo=my-lib)'; exit 1; \
	fi; \
	\
	# --------------------------------------------------------------------
	# 2. Global search-and-replace
	# --------------------------------------------------------------------
	echo "🔧 Initializing template for $(owner)/$(repo)…"; \
	find . \
	  \( -path './.git' -o -path './.idea' -o -path './.make' \) -prune -o \
	  -type f \
	  \( -not -name '*.png' -and -not -name '*.jpg' -and -not -name '*.jpeg' \
	     -and -not -name '*.gif' -and -not -name '*.ico' \) \
	  -print0 | \
	  env LC_ALL=C xargs -0 $(SED) $(SED_INPLACE) \
	    -e 's#mrz1836/go-template#$(owner)/$(repo)#g' \
	    -e 's#go-template#$(repo)#g' \
	    -e 's#mrz1836#$(owner)#g'; \
	\
	# --------------------------------------------------------------------
	# 3. Remove default share image (optional)
	# --------------------------------------------------------------------
	if [ -f .github/IMAGES/go-share-image.png ]; then \
	  rm .github/IMAGES/go-share-image.png; \
	  echo '🗑️  Removed default social-share image (.github/IMAGES/go-share-image.png)'; \
	fi; \
	\
	# --------------------------------------------------------------------
	# 4. Success message
	# --------------------------------------------------------------------
	echo '✅ Done! Review the changes, review all files, and commit the initial version.'
