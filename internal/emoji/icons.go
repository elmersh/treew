package emoji

import (
	"fmt"
	"sync"

	"github.com/elmersh/treew/internal/config"
)

// IconSet representa un conjunto de iconos para diferentes extensiones de archivo
type IconSet struct {
	NerdFont string
	Emoji    string
	Color    string // Color del icono (ANSI)
}

// Colores ANSI
const (
	Reset     = "\033[0m"
	Bold      = "\033[1m"
	Dim       = "\033[2m"
	Underline = "\033[4m"
	Blink     = "\033[5m"
	Reverse   = "\033[7m"
	Hidden    = "\033[8m"

	// Colores de texto
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

// FileIcons contiene íconos de Nerd Fonts y emojis para diferentes extensiones de archivos
var FileIcons = map[string]IconSet{
	// Desarrollo
	".cs":     {NerdFont: "\ueb27", Emoji: "📄", Color: Magenta}, // C#
	".vb":     {NerdFont: "\ueb27", Emoji: "📄", Color: Magenta}, // Visual Basic
	".java":   {NerdFont: "\ueb5b", Emoji: "☕", Color: Red},     // Java
	".py":     {NerdFont: "\ue235", Emoji: "🐍", Color: Blue},    // Python
	".rb":     {NerdFont: "\ue21e", Emoji: "💎", Color: Red},     // Ruby
	".php":    {NerdFont: "\ue73d", Emoji: "🐘", Color: Magenta}, // PHP
	".go":     {NerdFont: "\ue626", Emoji: "🐹", Color: Cyan},    // Go
	".rs":     {NerdFont: "\ue7a8", Emoji: "🦀", Color: Red},     // Rust
	".c":      {NerdFont: "\ue61e", Emoji: "🔧", Color: Blue},    // C
	".cpp":    {NerdFont: "\ue61d", Emoji: "🔧", Color: Blue},    // C++
	".h":      {NerdFont: "\ufb71", Emoji: "🔧", Color: Blue},    // Header
	".hpp":    {NerdFont: "\ufb71", Emoji: "🔧", Color: Blue},    // Header++
	".swift":  {NerdFont: "\ue755", Emoji: "🦅", Color: Yellow},  // Swift
	".kt":     {NerdFont: "\ue634", Emoji: "📱", Color: Magenta}, // Kotlin
	".kts":    {NerdFont: "\ue634", Emoji: "📱", Color: Magenta}, // Kotlin Script
	".scala":  {NerdFont: "\ue737", Emoji: "⚡", Color: Red},     // Scala
	".groovy": {NerdFont: "\ue775", Emoji: "🍇", Color: Magenta}, // Groovy
	".sh":     {NerdFont: "\uf489", Emoji: "🐚", Color: Green},   // Shell
	".bash":   {NerdFont: "\uf489", Emoji: "🐚", Color: Green},   // Bash
	".zsh":    {NerdFont: "\uf489", Emoji: "🐚", Color: Green},   // Zsh
	".fish":   {NerdFont: "\uf489", Emoji: "🐟", Color: Blue},    // Fish
	".ps1":    {NerdFont: "\uf489", Emoji: "💻", Color: Blue},    // PowerShell
	".psm1":   {NerdFont: "\uf489", Emoji: "💻", Color: Blue},    // PowerShell Module
	".bat":    {NerdFont: "\uf489", Emoji: "💻", Color: Blue},    // Batch
	".cmd":    {NerdFont: "\uf489", Emoji: "💻", Color: Blue},    // Command
	".exe":    {NerdFont: "\uf489", Emoji: "⚙️", Color: Green},  // Executable
	".dll":    {NerdFont: "\uf489", Emoji: "🔌", Color: Green},   // DLL
	".so":     {NerdFont: "\uf489", Emoji: "🔌", Color: Green},   // Shared Object
	".dylib":  {NerdFont: "\uf489", Emoji: "🔌", Color: Green},   // Dynamic Library

	// Web
	".html":    {NerdFont: "\uf13b", Emoji: "🌐", Color: Red},     // HTML
	".htm":     {NerdFont: "\uf13b", Emoji: "🌐", Color: Red},     // HTML
	".css":     {NerdFont: "\ue42e", Emoji: "🎨", Color: Blue},    // CSS
	".scss":    {NerdFont: "\ue42e", Emoji: "🎨", Color: Magenta}, // SCSS
	".sass":    {NerdFont: "\ue42e", Emoji: "🎨", Color: Magenta}, // SASS
	".less":    {NerdFont: "\ue42e", Emoji: "🎨", Color: Blue},    // LESS
	".js":      {NerdFont: "\ue781", Emoji: "📜", Color: Yellow},  // JavaScript
	".jsx":     {NerdFont: "\ue781", Emoji: "📜", Color: Yellow},  // JSX
	".ts":      {NerdFont: "\ue8ca", Emoji: "📜", Color: Blue},    // TypeScript
	".tsx":     {NerdFont: "\ue8ca", Emoji: "📜", Color: Blue},    // TSX
	".vue":     {NerdFont: "\ue7fa", Emoji: "🟢", Color: Green},   // Vue
	".svelte":  {NerdFont: "\ue7fa", Emoji: "🟠", Color: Red},     // Svelte
	".astro":   {NerdFont: "\ue7fa", Emoji: "🌠", Color: Magenta}, // Astro
	".graphql": {NerdFont: "\ue799", Emoji: "📊", Color: Magenta}, // GraphQL
	".gql":     {NerdFont: "\ue799", Emoji: "📊", Color: Magenta}, // GraphQL
	".wasm":    {NerdFont: "\ue7fa", Emoji: "🔷", Color: Blue},    // WebAssembly

	// Datos
	".json":     {NerdFont: "\ue60b", Emoji: "📋", Color: Blue},     // JSON
	".xml":      {NerdFont: "\ue799", Emoji: "📋", Color: Magenta},  // XML
	".yaml":     {NerdFont: "\ue60b", Emoji: "⚙️", Color: Green},   // YAML
	".yml":      {NerdFont: "\ue60b", Emoji: "⚙️", Color: Green},   // YAML
	".toml":     {NerdFont: "\ue60b", Emoji: "⚙️", Color: Green},   // TOML
	".ini":      {NerdFont: "\ue615", Emoji: "⚙️", Color: Green},   // INI
	".csv":      {NerdFont: "\uf1c0", Emoji: "📊", Color: Blue},     // CSV
	".sql":      {NerdFont: "\ue428", Emoji: "🗃️", Color: Magenta}, // SQL
	".db":       {NerdFont: "\ue428", Emoji: "🗃️", Color: Magenta}, // Database
	".sqlite":   {NerdFont: "\ue428", Emoji: "🗃️", Color: Magenta}, // SQLite
	".sqlite3":  {NerdFont: "\ue428", Emoji: "🗃️", Color: Magenta}, // SQLite3
	".mdb":      {NerdFont: "\ue428", Emoji: "🗃️", Color: Magenta}, // Access
	".accdb":    {NerdFont: "\ue428", Emoji: "🗃️", Color: Magenta}, // Access
	".parquet":  {NerdFont: "\uf1c0", Emoji: "📊", Color: Blue},     // Parquet
	".avro":     {NerdFont: "\uf1c0", Emoji: "📊", Color: Blue},     // Avro
	".protobuf": {NerdFont: "\uf1c0", Emoji: "📊", Color: Blue},     // Protocol Buffers
	".proto":    {NerdFont: "\uf1c0", Emoji: "📊", Color: Blue},     // Protocol Buffers

	// Documentos
	".md":    {NerdFont: "\uf48a", Emoji: "📝", Color: Blue},    // Markdown
	".txt":   {NerdFont: "\uf15c", Emoji: "📄", Color: Blue},    // Text
	".doc":   {NerdFont: "\uf1c2", Emoji: "📘", Color: Magenta}, // Word
	".docx":  {NerdFont: "\uf1c2", Emoji: "📘", Color: Magenta}, // Word
	".pdf":   {NerdFont: "\uf1c1", Emoji: "📕", Color: Magenta}, // PDF
	".xls":   {NerdFont: "\uf1c3", Emoji: "📊", Color: Magenta}, // Excel
	".xlsx":  {NerdFont: "\uf1c3", Emoji: "📗", Color: Magenta}, // Excel
	".ppt":   {NerdFont: "\uf1c5", Emoji: "📙", Color: Magenta}, // PowerPoint
	".pptx":  {NerdFont: "\uf1c5", Emoji: "📙", Color: Magenta}, // PowerPoint
	".odt":   {NerdFont: "\uf1c2", Emoji: "📘", Color: Magenta}, // OpenDocument Text
	".ods":   {NerdFont: "\uf1c3", Emoji: "📗", Color: Magenta}, // OpenDocument Spreadsheet
	".odp":   {NerdFont: "\uf1c5", Emoji: "📙", Color: Magenta}, // OpenDocument Presentation
	".rtf":   {NerdFont: "\uf15c", Emoji: "📄", Color: Blue},    // Rich Text
	".tex":   {NerdFont: "\uf15c", Emoji: "📝", Color: Blue},    // TeX
	".latex": {NerdFont: "\uf15c", Emoji: "📝", Color: Blue},    // LaTeX
	".epub":  {NerdFont: "\uf02d", Emoji: "📚", Color: Magenta}, // EPUB
	".mobi":  {NerdFont: "\uf02d", Emoji: "📚", Color: Magenta}, // MOBI
	".azw3":  {NerdFont: "\uf02d", Emoji: "📚", Color: Magenta}, // Kindle
	".fb2":   {NerdFont: "\uf02d", Emoji: "📚", Color: Magenta}, // FictionBook

	// Imágenes y Media
	".jpg":  {NerdFont: "\uf1c5", Emoji: "🖼️", Color: Magenta}, // JPEG
	".jpeg": {NerdFont: "\uf1c5", Emoji: "🖼️", Color: Magenta}, // JPEG
	".png":  {NerdFont: "\uf1c5", Emoji: "🖼️", Color: Magenta}, // PNG
	".gif":  {NerdFont: "\uf1c5", Emoji: "🖼️", Color: Magenta}, // GIF
	".svg":  {NerdFont: "\uf1c5", Emoji: "🖼️", Color: Magenta}, // SVG
	".webp": {NerdFont: "\uf1c5", Emoji: "🖼️", Color: Magenta}, // WebP
	".bmp":  {NerdFont: "\uf1c5", Emoji: "🖼️", Color: Magenta}, // BMP
	".tiff": {NerdFont: "\uf1c5", Emoji: "🖼️", Color: Magenta}, // TIFF
	".ico":  {NerdFont: "\uf1c5", Emoji: "🖼️", Color: Magenta}, // ICO
	".mp3":  {NerdFont: "\uf1c7", Emoji: "🎵", Color: Magenta},  // MP3
	".wav":  {NerdFont: "\uf1c7", Emoji: "🎵", Color: Magenta},  // WAV
	".ogg":  {NerdFont: "\uf1c7", Emoji: "🎵", Color: Magenta},  // OGG
	".flac": {NerdFont: "\uf1c7", Emoji: "🎵", Color: Magenta},  // FLAC
	".aac":  {NerdFont: "\uf1c7", Emoji: "🎵", Color: Magenta},  // AAC
	".m4a":  {NerdFont: "\uf1c7", Emoji: "🎵", Color: Magenta},  // M4A
	".mp4":  {NerdFont: "\uf1c8", Emoji: "🎥", Color: Magenta},  // MP4
	".mov":  {NerdFont: "\uf1c8", Emoji: "🎥", Color: Magenta},  // MOV
	".avi":  {NerdFont: "\uf1c8", Emoji: "🎥", Color: Magenta},  // AVI
	".mkv":  {NerdFont: "\uf1c8", Emoji: "🎥", Color: Magenta},  // MKV
	".webm": {NerdFont: "\uf1c8", Emoji: "🎥", Color: Magenta},  // WebM
	".flv":  {NerdFont: "\uf1c8", Emoji: "🎥", Color: Magenta},  // FLV
	".wmv":  {NerdFont: "\uf1c8", Emoji: "🎥", Color: Magenta},  // WMV
	".mpg":  {NerdFont: "\uf1c8", Emoji: "🎥", Color: Magenta},  // MPG
	".mpeg": {NerdFont: "\uf1c8", Emoji: "🎥", Color: Magenta},  // MPEG
	".3gp":  {NerdFont: "\uf1c8", Emoji: "🎥", Color: Magenta},  // 3GP
	".iso":  {NerdFont: "\uf1c0", Emoji: "💿", Color: Magenta},  // ISO
	".cue":  {NerdFont: "\uf1c0", Emoji: "💿", Color: Magenta},  // CUE
	".bin":  {NerdFont: "\uf1c0", Emoji: "💿", Color: Magenta},  // BIN

	// Archivos de proyecto y configuración
	".sln":                          {NerdFont: "\ue77c", Emoji: "🔨", Color: Magenta},  // Solution
	".csproj":                       {NerdFont: "\ue77c", Emoji: "🔧", Color: Magenta},  // C# Project
	".vbproj":                       {NerdFont: "\ue77c", Emoji: "🔧", Color: Magenta},  // VB Project
	".fsproj":                       {NerdFont: "\ue77c", Emoji: "🔧", Color: Magenta},  // F# Project
	".conf":                         {NerdFont: "\ue615", Emoji: "⚙️", Color: Green},   // Config
	".config":                       {NerdFont: "\ue615", Emoji: "⚙️", Color: Green},   // Config
	".env":                          {NerdFont: "\uf013", Emoji: "🔒", Color: Magenta},  // Environment
	".gitignore":                    {NerdFont: "\uf1d3", Emoji: "👁️", Color: Magenta}, // Git Ignore
	".gitattributes":                {NerdFont: "\uf1d3", Emoji: "👁️", Color: Magenta}, // Git Attributes
	".editorconfig":                 {NerdFont: "\ue615", Emoji: "⚙️", Color: Green},   // Editor Config
	".dockerignore":                 {NerdFont: "\uf308", Emoji: "🐳", Color: Magenta},  // Docker Ignore
	".dockerfile":                   {NerdFont: "\uf308", Emoji: "🐳", Color: Magenta},  // Dockerfile
	".docker":                       {NerdFont: "\uf308", Emoji: "🐳", Color: Magenta},  // Docker
	".compose":                      {NerdFont: "\uf308", Emoji: "🐳", Color: Magenta},  // Docker Compose
	".properties":                   {NerdFont: "\ue615", Emoji: "⚙️", Color: Green},   // Properties
	".gradle":                       {NerdFont: "\ue615", Emoji: "⚙️", Color: Green},   // Gradle
	".maven":                        {NerdFont: "\ue615", Emoji: "⚙️", Color: Green},   // Maven
	".pom":                          {NerdFont: "\ue615", Emoji: "⚙️", Color: Green},   // POM
	".lock":                         {NerdFont: "\uf023", Emoji: "🔒", Color: Magenta},  // Lock
	".package-lock.json":            {NerdFont: "\uf023", Emoji: "🔒", Color: Magenta},  // Package Lock
	".yarn.lock":                    {NerdFont: "\uf023", Emoji: "🔒", Color: Magenta},  // Yarn Lock
	".cargo.lock":                   {NerdFont: "\uf023", Emoji: "🔒", Color: Magenta},  // Cargo Lock
	".poetry.lock":                  {NerdFont: "\uf023", Emoji: "🔒", Color: Magenta},  // Poetry Lock
	".composer.lock":                {NerdFont: "\uf023", Emoji: "🔒", Color: Magenta},  // Composer Lock
	".gemfile.lock":                 {NerdFont: "\uf023", Emoji: "🔒", Color: Magenta},  // Gemfile Lock
	".requirements.txt":             {NerdFont: "\uf1d6", Emoji: "📦", Color: Magenta},  // Requirements
	".setup.py":                     {NerdFont: "\uf1d6", Emoji: "📦", Color: Magenta},  // Setup
	".package.json":                 {NerdFont: "\uf1d6", Emoji: "📦", Color: Magenta},  // Package
	".cargo.toml":                   {NerdFont: "\uf1d6", Emoji: "📦", Color: Magenta},  // Cargo
	".go.mod":                       {NerdFont: "\uf1d6", Emoji: "📦", Color: Magenta},  // Go Module
	".go.sum":                       {NerdFont: "\uf1d6", Emoji: "📦", Color: Magenta},  // Go Sum
	".composer.json":                {NerdFont: "\uf1d6", Emoji: "📦", Color: Magenta},  // Composer
	".gemfile":                      {NerdFont: "\uf1d6", Emoji: "📦", Color: Magenta},  // Gemfile
	".rakefile":                     {NerdFont: "\uf1d6", Emoji: "📦", Color: Magenta},  // Rakefile
	".makefile":                     {NerdFont: "\uf1d6", Emoji: "📦", Color: Magenta},  // Makefile
	".cmake":                        {NerdFont: "\uf1d6", Emoji: "📦", Color: Magenta},  // CMake
	".ninja":                        {NerdFont: "\uf1d6", Emoji: "📦", Color: Magenta},  // Ninja
	".bazel":                        {NerdFont: "\uf1d6", Emoji: "📦", Color: Magenta},  // Bazel
	".buck":                         {NerdFont: "\uf1d6", Emoji: "📦", Color: Magenta},  // Buck
	".pants":                        {NerdFont: "\uf1d6", Emoji: "📦", Color: Magenta},  // Pants
	".bazelrc":                      {NerdFont: "\uf1d6", Emoji: "📦", Color: Magenta},  // Bazel RC
	".buckconfig":                   {NerdFont: "\uf1d6", Emoji: "📦", Color: Magenta},  // Buck Config
	".pants.ini":                    {NerdFont: "\uf1d6", Emoji: "📦", Color: Magenta},  // Pants INI
	".travis.yml":                   {NerdFont: "\ue77e", Emoji: "🔧", Color: Magenta},  // Travis
	".circleci":                     {NerdFont: "\ue77e", Emoji: "🔧", Color: Magenta},  // CircleCI
	".github":                       {NerdFont: "\uf408", Emoji: "🔧", Color: Magenta},  // GitHub
	".gitlab":                       {NerdFont: "\uf296", Emoji: "🔧", Color: Magenta},  // GitLab
	".bitbucket":                    {NerdFont: "\uf171", Emoji: "🔧", Color: Magenta},  // Bitbucket
	".jenkins":                      {NerdFont: "\ue77e", Emoji: "🔧", Color: Magenta},  // Jenkins
	".drone":                        {NerdFont: "\ue77e", Emoji: "🔧", Color: Magenta},  // Drone
	".azure-pipelines.yml":          {NerdFont: "\ufd03", Emoji: "🔧", Color: Magenta},  // Azure
	".appveyor.yml":                 {NerdFont: "\ue77e", Emoji: "🔧", Color: Magenta},  // AppVeyor
	".codeship":                     {NerdFont: "\ue77e", Emoji: "🔧", Color: Magenta},  // Codeship
	".semaphore":                    {NerdFont: "\ue77e", Emoji: "🔧", Color: Magenta},  // Semaphore
	".wercker":                      {NerdFont: "\ue77e", Emoji: "🔧", Color: Magenta},  // Wercker
	".buildkite":                    {NerdFont: "\ue77e", Emoji: "🔧", Color: Magenta},  // Buildkite
	".teamcity":                     {NerdFont: "\ue77e", Emoji: "🔧", Color: Magenta},  // TeamCity
	".bamboo":                       {NerdFont: "\ue77e", Emoji: "🔧", Color: Magenta},  // Bamboo
	".concourse":                    {NerdFont: "\ue77e", Emoji: "🔧", Color: Magenta},  // Concourse
	".spinnaker":                    {NerdFont: "\ue77e", Emoji: "🔧", Color: Magenta},  // Spinnaker
	".argocd":                       {NerdFont: "\ue77e", Emoji: "🔧", Color: Magenta},  // ArgoCD
	".flux":                         {NerdFont: "\ue77e", Emoji: "🔧", Color: Magenta},  // Flux
	".helm":                         {NerdFont: "\ufd31", Emoji: "🔧", Color: Magenta},  // Helm
	".kubernetes":                   {NerdFont: "\ufd31", Emoji: "🔧", Color: Magenta},  // Kubernetes
	".k8s":                          {NerdFont: "\ufd31", Emoji: "🔧", Color: Magenta},  // K8s
	".docker-compose.yml":           {NerdFont: "\uf308", Emoji: "🐳", Color: Magenta},  // Docker Compose
	".docker-compose.yaml":          {NerdFont: "\uf308", Emoji: "🐳", Color: Magenta},  // Docker Compose
	".docker-compose.override.yml":  {NerdFont: "\uf308", Emoji: "🐳", Color: Magenta},  // Docker Compose Override
	".docker-compose.override.yaml": {NerdFont: "\uf308", Emoji: "🐳", Color: Magenta},  // Docker Compose Override
	".vagrant":                      {NerdFont: "\ue77e", Emoji: "🔧", Color: Magenta},  // Vagrant
	".vagrantfile":                  {NerdFont: "\ue77e", Emoji: "🔧", Color: Magenta},  // Vagrantfile
	".terraform":                    {NerdFont: "\ufd31", Emoji: "🔧", Color: Magenta},  // Terraform
	".tf":                           {NerdFont: "\ufd31", Emoji: "🔧", Color: Magenta},  // Terraform
	".tfvars":                       {NerdFont: "\ufd31", Emoji: "🔧", Color: Magenta},  // Terraform Vars
	".tfstate":                      {NerdFont: "\ufd31", Emoji: "🔧", Color: Magenta},  // Terraform State
	".tfstate.backup":               {NerdFont: "\ufd31", Emoji: "🔧", Color: Magenta},  // Terraform State Backup
	".tfplan":                       {NerdFont: "\ufd31", Emoji: "🔧", Color: Magenta},  // Terraform Plan
	".tf.json":                      {NerdFont: "\ufd31", Emoji: "🔧", Color: Magenta},  // Terraform JSON
	".tfvars.json":                  {NerdFont: "\ufd31", Emoji: "📦", Color: Magenta},  // Terraform Vars JSON
	".tfvars.yaml":                  {NerdFont: "\ufd31", Emoji: "🔧", Color: Magenta},  // Terraform Vars YAML
	".tfvars.yml":                   {NerdFont: "\ufd31", Emoji: "🔧", Color: Magenta},  // Terraform Vars YAML
	".tfvars.hcl":                   {NerdFont: "\ufd31", Emoji: "🔧", Color: Magenta},  // Terraform Vars HCL

	// Comprimidos
	".zip":      {NerdFont: "\uf1c6", Emoji: "📦", Color: Magenta}, // ZIP
	".rar":      {NerdFont: "\uf1c6", Emoji: "📦", Color: Magenta}, // RAR
	".7z":       {NerdFont: "\uf1c6", Emoji: "📦", Color: Magenta}, // 7Z
	".tar":      {NerdFont: "\uf1c6", Emoji: "📦", Color: Magenta}, // TAR
	".gz":       {NerdFont: "\uf1c6", Emoji: "📦", Color: Magenta}, // GZ
	".bz2":      {NerdFont: "\uf1c6", Emoji: "📦", Color: Magenta}, // BZ2
	".tar.gz":   {NerdFont: "\uf1c6", Emoji: "📦", Color: Magenta}, // TAR GZ
	".tar.bz2":  {NerdFont: "\uf1c6", Emoji: "📦", Color: Magenta}, // TAR BZ2
	".tar.xz":   {NerdFont: "\uf1c6", Emoji: "📦", Color: Magenta}, // TAR XZ
	".tar.lz":   {NerdFont: "\uf1c6", Emoji: "📦", Color: Magenta}, // TAR LZ
	".tar.lzma": {NerdFont: "\uf1c6", Emoji: "📦", Color: Magenta}, // TAR LZMA
	".tar.lzo":  {NerdFont: "\uf1c6", Emoji: "📦", Color: Magenta}, // TAR LZO

	// Por defecto
	"default": {NerdFont: "\uf15c", Emoji: "📄", Color: Blue}, // Default
	"folder":  {NerdFont: "\uf115", Emoji: "📁", Color: Cyan}, // Folder
}

var (
	UseNerdFonts bool
	configMutex  sync.RWMutex
)

// InitFromConfig inicializa la configuración de iconos desde la configuración del usuario
func InitFromConfig(cfg *config.Config) {
	configMutex.Lock()
	defer configMutex.Unlock()
	UseNerdFonts = cfg.UseNerdFonts
}

// SetUseNerdFonts establece si se deben usar iconos de Nerd Fonts
func SetUseNerdFonts(use bool) {
	configMutex.Lock()
	defer configMutex.Unlock()
	UseNerdFonts = use
}

// FolderIcon obtiene el ícono de carpeta
func FolderIcon() string {
	configMutex.RLock()
	defer configMutex.RUnlock()

	icon, exists := FileIcons["folder"]
	if !exists {
		return "📁"
	}
	if UseNerdFonts {
		return fmt.Sprintf("%s%s%s", icon.Color, icon.NerdFont, Reset)
	}
	return icon.Emoji
}

// GetFileIcon obtiene el ícono para una extensión de archivo
func GetFileIcon(extension string) string {
	configMutex.RLock()
	defer configMutex.RUnlock()

	icon, exists := FileIcons[extension]
	if !exists {
		if UseNerdFonts {
			defaultIcon := FileIcons["default"]
			return fmt.Sprintf("%s%s%s", defaultIcon.Color, defaultIcon.NerdFont, Reset)
		}
		return FileIcons["default"].Emoji
	}
	if UseNerdFonts {
		return fmt.Sprintf("%s%s%s", icon.Color, icon.NerdFont, Reset)
	}
	return icon.Emoji
}

// GetNerdFontIcon obtiene el ícono de Nerd Fonts para una extensión de archivo
func GetNerdFontIcon(extension string) string {
	configMutex.RLock()
	defer configMutex.RUnlock()

	icon, exists := FileIcons[extension]
	if !exists {
		defaultIcon := FileIcons["default"]
		return fmt.Sprintf("%s%s%s", defaultIcon.Color, defaultIcon.NerdFont, Reset)
	}
	return fmt.Sprintf("%s%s%s", icon.Color, icon.NerdFont, Reset)
}
