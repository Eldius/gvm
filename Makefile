
test:
	go test ./... -cover -coverprofile=coverage.out

lint:
	# revive -formatter stylish ./...
	golangci-lint run

releaser:
	goreleaser build

vuln:
	govulncheck ./...

run:
	go run ./cmd/cli/ install 1.21.0


ls-remote:
	go run ./cmd/cli/ ls-remote

release-build:
	@make release bump=build

release-minor:
	@make release bump=minor

release-major:
	@make release bump=major

release:
	$(eval v0 := $(shell git describe --tags --abbrev=0 | sed -Ee 's/^v|-.*//'))
ifeq ($(bump), major)
	$(eval f := 1)
else ifeq ($(bump), minor)
	$(eval f := 2)
else
	$(eval f := 3)
endif
	$(eval v := $(shell echo $(v0) | awk -F. -v OFS=. -v f=$(f) '{ $$f++ } 1'))
	@echo "current version: $(v0)"
	@echo "next version: $(v)"
	git tag v$(v)
