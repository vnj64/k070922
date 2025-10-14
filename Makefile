ifneq (,$(wildcard .env))
    include .env
    export
endif

swagger:
	swag init --parseDependency -g cmd/main.go --output=./docs