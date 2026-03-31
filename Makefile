GOBIN := "${GOPATH}/bin"

.PHONY: help
help:
	@echo "Targets:"
	@echo "    tidy:            tidy go mod"
	@echo "    lint:            run go-lint"

.PHONY: tidy
tidy:
	git ls-files go.mod '**/*go.mod' -z | xargs -0 -I{} bash -xc 'cd $$(dirname {}) && go mod tidy -v'

.PHONY: go-lint
go-lint:

.PHONY: generate
generate:
	git ls-files go.mod '**/*go.mod' -z | xargs -0 -I{} bash -xc 'cd $$(dirname {}) && go generate ./...'
