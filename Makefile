include .env
export

.PHONY: init
init:
	@./scripts/install-git-hooks.shg

# Formatter & Linter

.PHONY: prettify
prettify:
	@./scripts/prettify.sh