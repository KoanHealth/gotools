setup: ## Install all the build and lint dependencies
	go install golang.org/x/tools/cmd/goimports
	go install github.com/golang/lint/golint
	go install github.com/onsi/ginkgo/ginkgo/v2@latest

lint:
	golint ./...

format:
	go fmt ./...

build: $(shell find . -name "*.go")
	go install -v ./...

clean:
	go clean

test: build
	ginkgo -r

version:
	git tag $(shell git tag --list --sort=-v:refname v* | head -n 1 | awk -F. -v OFS=. '{$$NF++;print}')
	git push origin master --tag

ci:
	docker-compose -f .github/compose/docker-compose.yml build
	docker-compose -f .github/compose/docker-compose.yml run test
