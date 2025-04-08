# Makefile para Treew

# Variables
BINARY_NAME=treew
VERSION=$(shell git describe --tags --always --dirty)
BUILD_DIR=build
LDFLAGS=-ldflags "-X main.Version=$(VERSION)"
GO_FILES=./cmd/treew/main.go

# Plataformas para compilación cruzada (simplificado)
PLATFORMS=linux/amd64 darwin/amd64 windows/amd64

# Colores para la salida
GREEN=\033[0;32m
BLUE=\033[0;34m
NC=\033[0m # No Color

.PHONY: all clean build-all build test install uninstall

all: clean build-all

# Limpiar archivos compilados
clean:
	@echo "$(BLUE)🧹 Limpiando archivos compilados...$(NC)"
	@rm -rf $(BUILD_DIR)
	@echo "$(GREEN)✅ Limpieza completada$(NC)"

# Compilar para todas las plataformas
build-all:
	@echo "$(BLUE)🌍 Compilando para múltiples plataformas...$(NC)"
	@mkdir -p $(BUILD_DIR)
	@for platform in $(PLATFORMS); do \
		os=$${platform%/*}; \
		arch=$${platform#*/}; \
		echo "$(BLUE)📦 Compilando para $$os/$$arch...$(NC)"; \
		if [ "$$os" = "windows" ]; then \
			GOOS=$$os GOARCH=$$arch go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-$$os-$$arch.exe $(GO_FILES); \
		else \
			GOOS=$$os GOARCH=$$arch go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-$$os-$$arch $(GO_FILES); \
		fi; \
	done
	@echo "$(GREEN)✅ Compilación completada$(NC)"

# Compilar para la plataforma actual
build:
	@echo "$(BLUE)📦 Compilando para la plataforma actual...$(NC)"
	@go build $(LDFLAGS) -o $(BINARY_NAME) $(GO_FILES)
	@echo "$(GREEN)✅ Compilación completada$(NC)"

# Ejecutar pruebas
test:
	@echo "$(BLUE)🧪 Ejecutando pruebas...$(NC)"
	@go test -v ./...
	@echo "$(GREEN)✅ Pruebas completadas$(NC)"

# Instalar en el sistema
install:
	@echo "$(BLUE)📋 Instalando Treew...$(NC)"
	@./scripts/install.sh
	@echo "$(GREEN)✅ Instalación completada$(NC)"

# Desinstalar del sistema
uninstall:
	@echo "$(BLUE)🗑️ Desinstalando Treew...$(NC)"
	@rm -f /usr/local/bin/$(BINARY_NAME) $(HOME)/.local/bin/$(BINARY_NAME)
	@echo "$(GREEN)✅ Desinstalación completada$(NC)"

# Crear archivos de release
release: build-all
	@echo "$(BLUE)📦 Creando archivos de release...$(NC)"
	@mkdir -p $(BUILD_DIR)/release
	@for platform in $(PLATFORMS); do \
		os=$${platform%/*}; \
		arch=$${platform#*/}; \
		echo "$(BLUE)📝 Creando release para $$os/$$arch...$(NC)"; \
		mkdir -p $(BUILD_DIR)/release/treew-$(VERSION)-$$os-$$arch; \
		cp $(BUILD_DIR)/$(BINARY_NAME)-$$os-$$arch* $(BUILD_DIR)/release/treew-$(VERSION)-$$os-$$arch/; \
		cp README.md $(BUILD_DIR)/release/treew-$(VERSION)-$$os-$$arch/; \
		if [ -f "LICENSE" ]; then \
			cp LICENSE $(BUILD_DIR)/release/treew-$(VERSION)-$$os-$$arch/; \
		fi; \
		if [ "$$os" = "windows" ]; then \
			cp scripts/install.ps1 scripts/uninstall.ps1 $(BUILD_DIR)/release/treew-$(VERSION)-$$os-$$arch/; \
		else \
			cp scripts/install.sh scripts/uninstall.sh $(BUILD_DIR)/release/treew-$(VERSION)-$$os-$$arch/; \
		fi; \
		cd $(BUILD_DIR)/release && tar -czf treew-$(VERSION)-$$os-$$arch.tar.gz treew-$(VERSION)-$$os-$$arch; \
		cd ../..; \
	done
	@echo "$(GREEN)✅ Archivos de release creados$(NC)"

# Ayuda
help:
	@echo "$(BLUE)🌲 Treew - Makefile$(NC)"
	@echo ""
	@echo "Uso:"
	@echo "  make [comando]"
	@echo ""
	@echo "Comandos:"
	@echo "  all         Limpia y compila para todas las plataformas"
	@echo "  clean        Limpia archivos compilados"
	@echo "  build-all    Compila para todas las plataformas"
	@echo "  build        Compila para la plataforma actual"
	@echo "  test         Ejecuta las pruebas"
	@echo "  install      Instala Treew en el sistema"
	@echo "  uninstall    Desinstala Treew del sistema"
	@echo "  release      Crea archivos de release para todas las plataformas"
	@echo "  help         Muestra esta ayuda" 