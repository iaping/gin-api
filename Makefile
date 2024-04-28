default: dev

pro:
	@go build -trimpath -ldflags '-w -s' -v -o ./bin/gin-api .

dev: doc
	@go build -v -o ./bin/gin-api .

doc:
	@swag init