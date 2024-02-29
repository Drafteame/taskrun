DEFAULT_GOAL := help

help:
	@echo "build    build the taskrun binary"
	@echo "lint     run linter"
	@echo "format   run goimports-reviser"
	@echo "test     run tests"

build:
	go build -ldflags="-s -w -X 'github.com/Drafteame/taskrun/internal.Version=0.0.0'" -o ./.bin/taskrun ./cmd


lint:
	go vet ./... && revive -config=revive.toml -formatter=friendly ./...


format:
	goimports-reviser -format -excludes=bin,node_modules,tmp,.git ./...


test:
	go test -v -race ./... -cover -coverprofile=coverage.out


precommit:
	pre-commit install --install-hooks && pre-commit install --install-hooks --hook-type commit-msg