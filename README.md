# service_template

---

- [Подготовка](#подготовка)
  - [Установка protoc с плагинами на macos](#установка-protoc-с-плагинами-на-macos)
  - [Генерация контрактов](#генерация-контрактов)

---

# Подготовка
## Установка protoc с плагинами на macos

```shell
brew install protobuf
```
- Версия должна быть 3.0 или выше
```shell
protoc --version
```
- Плагин `protoc-gen-go` для генерации Go-кода
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
- Плагин `protoc-gen-go-grpc` для генерации gRPC-кода
```shell
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
- Убедись, что `protoc-gen-go` и `protoc-gen-go-grpc` находятся в вашем `$PATH`. Обычно Go устанавливает бинарные файлы в `$GOPATH/bin`, поэтому добавьте его в `$PATH`, если ещё не сделали этого
```shell
export PATH="$PATH:$(go env GOPATH)/bin"
```
- Установка `protoc-gen-grpc-gateway`
```shell
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
```
- Установка плагина для генерации OpenAPI (Swagger) документации, если потребуется
```shell
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
```

## Генерация контрактов
- Скачивание зависимостей для proto
```shell
make download
```
- Генерация контрактов сервиса
```shell
make all
```