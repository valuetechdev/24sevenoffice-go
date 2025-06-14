.PHONY: tidy
tidy:
	go mod tidy
	go fmt ./...

.PHONY: generate
generate:
	go generate ./...
	go fmt ./...
	sed -i '' 's/int \[\]/Int \[\]/g' ./soap24/**/*.gen.go

.PHONY: api-payroll24
api-payroll24:
	@echo "openapi: getting latest Swagger from Payroll24 and converting to OpenAPI 3"
	@curl https://converter.swagger.io/api/convert?url=https://me.24sevenoffice.com/swagger.json > ./api/openapi/payroll.json

.PHONY: api-rest24
api-rest24:
	@echo "openapi: getting latest Rest24"
	@curl https://rest-api.developer.24sevenoffice.com/doc/v1.json > ./api/openapi/rest.json

.PHONY: api
api:
	make -j2 api-payroll24 api-rest24	
	
.PHONY: check
check:
	go vet ./...
	go tool github.com/golangci/golangci-lint/cmd/golangci-lint run ./...
	go tool honnef.co/go/tools/cmd/staticcheck ./...

