#Makefile
API_NAME=bmw-data-repoitory

build:
	@echo "Creando Binario ..."
	@go build -o build/bin/${API_NAME} cmd/main.go
	@echo "Binario generado en build/bin/${API_NAME}"

test:
	@echo "Ejecutando tests..."
	@go test ./...

run:
	@echo "Creando Binario ..."
	@go build -o build/bin/${API_NAME} cmd/main.go
	@echo "Binario generado en build/bin/${API_NAME}"
	@echo "Ejecutando Binario ..."
	@build/bin/${API_NAME}