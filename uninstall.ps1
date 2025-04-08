# Script de desinstalación para Windows
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

Write-ColorOutput $Blue "🗑️ Desinstalando Treew..."

# Determinar la ubicación de instalación
$installDir = "$env:USERPROFILE\AppData\Local\Programs\Treew"
$binDir = "$installDir\bin"
$binaryPath = "$binDir\treew.exe"

# Verificar si Treew está instalado
if (-not (Test-Path $binaryPath)) {
    Write-ColorOutput $Red "❌ Error: Treew no está instalado en la ubicación esperada: $binaryPath"
    Write-ColorOutput $Red "Si lo instalaste en otra ubicación, por favor elimínalo manualmente."
    exit 1
}

# Eliminar el binario
Remove-Item -Path $binaryPath -Force
Write-ColorOutput $Blue "🗑️ Eliminado binario: $binaryPath"

# Eliminar directorios si están vacíos
if ((Get-ChildItem -Path $binDir -ErrorAction SilentlyContinue).Count -eq 0) {
    Remove-Item -Path $binDir -Force
    Write-ColorOutput $Blue "🗑️ Eliminado directorio de binarios: $binDir"
}

if ((Get-ChildItem -Path $installDir -ErrorAction SilentlyContinue).Count -eq 0) {
    Remove-Item -Path $installDir -Force
    Write-ColorOutput $Blue "🗑️ Eliminado directorio de instalación: $installDir"
}

# Eliminar del PATH
$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($userPath -like "*$binDir*") {
    $newPath = ($userPath.Split(';') | Where-Object { $_ -ne $binDir }) -join ';'
    [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
    Write-ColorOutput $Blue "🔄 Eliminado $binDir de tu PATH de usuario."
    Write-ColorOutput $Blue "📝 Por favor, reinicia tu terminal para que los cambios surtan efecto."
}

# Preguntar si desea eliminar el archivo de configuración
$configFile = "$env:USERPROFILE\.config\treew.yaml"
if (Test-Path $configFile) {
    $response = Read-Host "¿Deseas eliminar el archivo de configuración? (s/n)"
    if ($response -eq "s" -or $response -eq "S") {
        Remove-Item -Path $configFile -Force
        Write-ColorOutput $Blue "🗑️ Eliminado archivo de configuración: $configFile"
    }
}

Write-ColorOutput $Green "✅ ¡Treew desinstalado correctamente!" 