# Директории
PROTO_SRC_DIR := api
OUT_DIR := pkg
GOOGLE_API_DIR := $(PROTO_SRC_DIR)/google

# Список всех proto файлов
PROTO_FILES := $(shell find $(PROTO_SRC_DIR) -name "*.proto" | grep -vE '^$(PROTO_SRC_DIR)/google')

# Создание директорий для вывода
.PHONY: dirs
dirs:
	@mkdir -p $(OUT_DIR)/cat_v1
	@mkdir -p $(OUT_DIR)/cat_admin_v1
	@echo "📁 Output directories created: $(OUT_DIR)/cat_v1 and $(OUT_DIR)/cat_admin_v1"

# Основная команда для генерации gRPC и REST кода
.PHONY: all
all: dirs
	@for proto_file in $(PROTO_FILES); do \
		protoc \
			-I=$(PROTO_SRC_DIR) \
			-I=$(GOOGLE_API_DIR) \
			--go_out=paths=source_relative:$(OUT_DIR) \
			--go-grpc_out=paths=source_relative:$(OUT_DIR) \
			--grpc-gateway_out=paths=source_relative:$(OUT_DIR) \
			$$proto_file; \
	done
	@echo "✅ All proto files processed successfully in $(OUT_DIR)"

# Цель для скачивания необходимых Google API файлов
ANNOTATIONS_PROTO := $(GOOGLE_API_DIR)/api/annotations.proto
HTTP_PROTO := $(GOOGLE_API_DIR)/api/http.proto
DESCRIPTOR_PROTO := $(GOOGLE_API_DIR)/protobuf/descriptor.proto

.PHONY: download
download: dirs
	@mkdir -p $(GOOGLE_API_DIR)/api $(GOOGLE_API_DIR)/protobuf
	@curl -sSL https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto -o $(ANNOTATIONS_PROTO)
	@curl -sSL https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto -o $(HTTP_PROTO)
	@curl -sSL https://raw.githubusercontent.com/protocolbuffers/protobuf/main/src/google/protobuf/descriptor.proto -o $(DESCRIPTOR_PROTO)
	@echo "📥 Downloaded necessary Google API proto files"

# Очистка
.PHONY: clean
clean:
	rm -rf $(OUT_DIR) $(GOOGLE_API_DIR)
	@echo "🧹 Cleaned generated files and Google API files"
