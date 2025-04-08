#!/bin/bash

set -e

# Colores para mejorar la salida
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${BLUE}🗑️ Desinstalando Treew...${NC}"

# Determinar la ubicación de instalación
INSTALL_DIR="/usr/local/bin"
if [ -f "$HOME/.local/bin/treew" ]; then
    INSTALL_DIR="$HOME/.local/bin"
fi

BINARY_PATH="$INSTALL_DIR/treew"

# Verificar si Treew está instalado
if [ ! -f "$BINARY_PATH" ]; then
    echo -e "${RED}❌ Error: Treew no está instalado en la ubicación esperada: $BINARY_PATH${NC}"
    echo -e "${RED}Si lo instalaste en otra ubicación, por favor elimínalo manualmente.${NC}"
    exit 1
fi

# Eliminar el binario
echo -e "${BLUE}🗑️ Eliminando binario: $BINARY_PATH${NC}"
rm -f "$BINARY_PATH"

# Preguntar si desea eliminar el archivo de configuración
CONFIG_FILE="$HOME/.config/treew.yaml"
if [ -f "$CONFIG_FILE" ]; then
    read -p "¿Deseas eliminar el archivo de configuración? (s/n): " response
    if [[ "$response" =~ ^[Ss]$ ]]; then
        echo -e "${BLUE}🗑️ Eliminando archivo de configuración: $CONFIG_FILE${NC}"
        rm -f "$CONFIG_FILE"
    fi
fi

echo -e "${GREEN}✅ ¡Treew desinstalado correctamente!${NC}" 