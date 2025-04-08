# Script de instalación para Windows
# Ejecutar como administrador: Start-Process powershell -Verb RunAs

# Colores para la salida
$Green = [System.ConsoleColor]::Green
$Blue = [System.ConsoleColor]::Blue
$Red = [System.ConsoleColor]::Red

# Función para escribir mensajes con color
function Write-ColorOutput($ForegroundColor) {
    $fc = $host.UI.RawUI.ForegroundColor
    $host.UI.RawUI.ForegroundColor = $ForegroundColor
    if ($args) {
        Write-Output $args
    }
    $host.UI.RawUI.ForegroundColor = $fc
}

Write-ColorOutput $Blue "🌲 Instalando Treew..."

# Obtener la ruta del script actual
$scriptPath = Split-Path -Parent $MyInvocation.MyCommand.Path

# Detectar arquitectura
$arch = "amd64"  # valor por defecto
$processorArchitecture = [System.Environment]::GetEnvironmentVariable("PROCESSOR_ARCHITECTURE")
$processorArchitectureW6432 = [System.Environment]::GetEnvironmentVariable("PROCESSOR_ARCHITEW6432")

if ($processorArchitecture -eq "ARM64" -or $processorArchitectureW6432 -eq "ARM64") {
    $arch = "arm64"
} elseif (-not [Environment]::Is64BitOperatingSystem) {
    $arch = "386"
}

$binaryName = "treew-windows-$arch.exe"
$binaryPath = Join-Path $scriptPath $binaryName

# Verificar si el binario existe
if (-not (Test-Path $binaryPath)) {
    Write-ColorOutput $Red "❌ Error: No se encontró el archivo $binaryName en el directorio actual."
    Write-ColorOutput $Red "Por favor, asegúrate de haber descargado la versión correcta para tu sistema."
    Write-ColorOutput $Red "Arquitectura detectada: $arch"
    exit 1
}

# Determinar dónde instalarlo
$installDir = "$env:USERPROFILE\AppData\Local\Programs\Treew"
$binDir = "$installDir\bin"

# Crear directorios si no existen
if (-not (Test-Path $installDir)) {
    New-Item -ItemType Directory -Path $installDir | Out-Null
    Write-ColorOutput $Blue "📁 Creado directorio de instalación: $installDir"
}

if (-not (Test-Path $binDir)) {
    New-Item -ItemType Directory -Path $binDir | Out-Null
    Write-ColorOutput $Blue "📁 Creado directorio de binarios: $binDir"
}

# Copiar el binario
Copy-Item -Path $binaryPath -Destination $binDir -Force
Write-ColorOutput $Blue "📋 Instalando binario en $binDir..."

# Renombrar el binario a treew.exe
$finalBinaryPath = Join-Path $binDir "treew.exe"
Rename-Item -Path (Join-Path $binDir $binaryName) -NewName "treew.exe" -Force
Write-ColorOutput $Blue "📝 Renombrando binario a treew.exe..."

# Crear archivo de configuración básico
$configDir = "$env:USERPROFILE\.config"
if (-not (Test-Path $configDir)) {
    New-Item -ItemType Directory -Path $configDir | Out-Null
    Write-ColorOutput $Blue "📁 Creado directorio de configuración: $configDir"
}

$configFile = "$configDir\treew.yaml"
if (-not (Test-Path $configFile)) {
    Write-ColorOutput $Blue "⚙️ Creando archivo de configuración..."
    @"
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
"@ | Out-File -FilePath $configFile -Encoding utf8
}

# Añadir al PATH
$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($userPath -notlike "*$binDir*") {
    [Environment]::SetEnvironmentVariable("Path", "$userPath;$binDir", "User")
    Write-ColorOutput $Blue "🔄 Añadido $binDir a tu PATH de usuario."
    Write-ColorOutput $Blue "📝 Por favor, reinicia tu terminal para que los cambios surtan efecto."
}

Write-ColorOutput $Green "✅ ¡Treew instalado correctamente!"
Write-ColorOutput $Blue "Para probar, abre una nueva terminal y ejecuta: treew" 