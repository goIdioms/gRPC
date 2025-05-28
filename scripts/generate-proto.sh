#!/bin/bash

# Установка необходимых инструментов
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Создание директорий для сгенерированных файлов
mkdir -p api/proto/user/v1
mkdir -p api/proto/order/v1

# Генерация Go кода из proto файлов
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       api/proto/user/v1/user.proto

protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       api/proto/order/v1/order.proto

echo "Proto files generated successfully!"
