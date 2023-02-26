
test:
	go test ./... -cover

lint:
	# revive -formatter stylish ./...
	golangci-lint run

release:
	goreleaser build
