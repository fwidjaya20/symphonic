include .env
export

.PHONY: init
init:
	@./scripts/install-git-hooks.shg