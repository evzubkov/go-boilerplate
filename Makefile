.SILENT:
.PHONY: help

## Список команд
help:
	printf "Available targets\n\n"
	awk '/^[a-zA-Z\-_0-9]+:/ { \
			helpMessage = match(lastLine, /^## (.*)/); \
			if (helpMessage) { \
					helpCommand = substr($$1, 0, index($$1, ":")-1); \
					helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
					printf "%-25s %s\n", helpCommand, helpMessage; \
			} \
	} \
	{ lastLine = $$0 }' ${MAKEFILE_LIST}


.PHONY: build-rest-api
## Собрать rest-api
build-rest-api: build-doc 
	go build -o build/rest-api ./cmd/rest-api

.PHONY: build-doc
## Собрать документацию
build-doc:
	cd ./cmd/rest-api && ~/work/bin/swag init

.PHONY: build-all
## Собрать все
build-all: build-doc build-rest-api

.PHONY: test
## Запустить тесты
test:
	go test -v -race -timeout 30s ./...