simple_http/run:
	@go run ./simple_http

oas/validate:
	@openapi-generator-cli validate -i ${spec}

.PHONY: simple_http/run oas/validate
