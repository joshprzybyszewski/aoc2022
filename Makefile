.DEFAULT_GOAL: help

.PHONY: help
help: ## Citation: https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: results
results: ## Regenerate the results shown in the README
	./scripts/populateResults.sh

.PHONY: start
start: ## Start today's puzzle! Cleans out last year's input and answers and fetches today's input
	./scripts/startDay.sh
