
test:
	go test ./... -cover -coverprofile=coverage.out

lint:
	# revive -formatter stylish ./...
	golangci-lint run

release:
	goreleaser build

vuln:
	govulncheck ./...

run:
	go run ./cmd/cli/ install 1.19.1
