CGO_ENABLED=1

.PHONY: test

generate-svg:
	scripts/puml2svg.sh

test:
	go test -cover ./...
