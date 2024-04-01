client_image = go-cobra-quiz-cli
server_image = go-cobra-quiz-server
user = a11199
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

run-local:
	go run main.go

requirements:
	make clean-packages
	go mod tidy

clean-packages:
	go clean -modcache

test: clean lint
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

go-test:
	go test -v ./...
bench:
	go test -bench .

up:
	docker compose up --build

down:
	docker compose down
	docker image rm ${client_image}
	docker image rm ${server_image}

push-client:
	docker push ${user}/${client_image}:latest

pull-client:
	docker push ${user}/${client_image}:latest

run-client:
	docker run ${client_image}
