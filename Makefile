lint: ## Runs linter for .go files
	@golangci-lint run --config .config/go.yml
	@echo "Go lint passed successfully"

clean:
	rm -rf ./build

#critic:
#	gocritic check -enableAll ./...
#
#security:
#	gosec ./...

test: clean lint
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

bench:
	go test -bench .
