#!/bin/bash

set -e

# Colores para mejorar la salida
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🌲 Instalando Treew...${NC}"

# Verificar si Go está instalado
if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Error: Go no está instalado.${NC}"
    echo "Por favor, instala Go desde https://golang.org/dl/"
    exit 1
fi

# Verificar la versión de Go
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
REQUIRED_VERSION="1.16.0"

if ! [[ "$(printf '%s\n' "$REQUIRED_VERSION" "$GO_VERSION" | sort -V | head -n1)" == "$REQUIRED_VERSION" ]]; then
    echo -e "${YELLOW}⚠️  Advertencia: Se recomienda Go 1.16.0 o superior (tienes $GO_VERSION)${NC}"
fi

# Compilar el programa
echo -e "${BLUE}📦 Compilando Treew...${NC}"
go build -o treew

# Determinar dónde instalarlo
INSTALL_DIR="$HOME/.local/bin"
if [ -d "/usr/local/bin" ] && [ -w "/usr/local/bin" ]; then
    INSTALL_DIR="/usr/local/bin"
fi

# Crear directorio si no existe
mkdir -p "$INSTALL_DIR"

# Mover el binario
echo -e "${BLUE}📋 Instalando binario en $INSTALL_DIR...${NC}"
mv treew "$INSTALL_DIR"

# Crear archivo de configuración básico
CONFIG_DIR="$HOME/.config"
mkdir -p "$CONFIG_DIR"

if [ ! -f "$CONFIG_DIR/treew.yaml" ]; then
    echo -e "${BLUE}⚙️  Creando archivo de configuración...${NC}"
    cat > "$CONFIG_DIR/treew.yaml" << EOF
exclude_folders:
  - node_modules
  - bin
  - obj
  - .git
  - packages
exclude_extensions: []
show_hidden: false
show_file_size: false
show_last_modified: false
max_depth: -1
EOF
fi

# Verificar PATH
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo -e "${BLUE}🔄 Añadiendo $INSTALL_DIR a tu PATH...${NC}"
    
    # Determinar el archivo de shell
    SHELL_FILE=""
    if [ -n "$ZSH_VERSION" ]; then
        SHELL_FILE="$HOME/.zshrc"
    elif [ -n "$BASH_VERSION" ]; then
        SHELL_FILE="$HOME/.bashrc"
    fi
    
    if [ -n "$SHELL_FILE" ]; then
        echo "export PATH=\"\$PATH:$INSTALL_DIR\"" >> "$SHELL_FILE"
        echo -e "${BLUE}📝 Añadido a $SHELL_FILE. Por favor, reinicia tu terminal o ejecuta: source $SHELL_FILE${NC}"
    else
        echo -e "${YELLOW}⚠️  No se pudo determinar tu shell. Por favor, añade manualmente $INSTALL_DIR a tu PATH.${NC}"
    fi
fi

echo -e "${GREEN}✅ ¡Treew instalado correctamente!${NC}"
echo -e "${BLUE}Para probar, ejecuta: treew${NC}"