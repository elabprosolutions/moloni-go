clean-cache:
	go clean -testcache

test: clean-cache
	go test -cover ./...

generate-models:
	oapi-codegen -generate types -package models -o ./pkg/billingprosolutions/models/api.go ./spec/openapi.yaml
