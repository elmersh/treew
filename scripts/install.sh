#!/bin/bash

set -e

# Colores para mejorar la salida
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

echo -e "${BLUE}🌲 Instalando Treew...${NC}"

# Obtener la ruta del script actual
SCRIPT_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Detectar sistema operativo y arquitectura
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Mapear arquitectura a nombres de binario
case $ARCH in
    x86_64)
        ARCH="amd64"
        ;;
    aarch64|arm64)
        ARCH="arm64"
        ;;
    *)
        echo -e "${RED}❌ Error: Arquitectura no soportada: $ARCH${NC}"
        exit 1
        ;;
esac

# Construir nombre del binario
BINARY_NAME="treew-${OS}-${ARCH}"
BINARY_PATH="$SCRIPT_PATH/$BINARY_NAME"

# Verificar si el binario existe
if [ ! -f "$BINARY_PATH" ]; then
    echo -e "${RED}❌ Error: No se encontró el archivo $BINARY_NAME en el directorio actual.${NC}"
    echo -e "${RED}Por favor, asegúrate de haber descargado la versión correcta para tu sistema.${NC}"
    echo -e "${YELLOW}Sistema detectado: $OS $ARCH${NC}"
    exit 1
fi

# Determinar dónde instalarlo
INSTALL_DIR="/usr/local/bin"
if [ ! -w "/usr/local/bin" ]; then
    INSTALL_DIR="$HOME/.local/bin"
fi

# Crear directorio si no existe
mkdir -p "$INSTALL_DIR"

# Mover el binario
echo -e "${BLUE}📋 Instalando binario en $INSTALL_DIR...${NC}"
cp "$BINARY_PATH" "$INSTALL_DIR/treew"
chmod +x "$INSTALL_DIR/treew"

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
use_nerd_fonts: false
EOF
fi

# Verificar PATH
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo -e "${BLUE}🔄 Añadiendo $INSTALL_DIR a tu PATH...${NC}"
    if [ -f "$HOME/.bashrc" ]; then
        echo "export PATH=\"$INSTALL_DIR:\$PATH\"" >> "$HOME/.bashrc"
    fi
    if [ -f "$HOME/.zshrc" ]; then
        echo "export PATH=\"$INSTALL_DIR:\$PATH\"" >> "$HOME/.zshrc"
    fi
    echo -e "${BLUE}📝 Por favor, reinicia tu terminal para que los cambios surtan efecto.${NC}"
fi

echo -e "${GREEN}✅ ¡Treew instalado correctamente!${NC}"
echo -e "${BLUE}Para probar, abre una nueva terminal y ejecuta: treew${NC}"