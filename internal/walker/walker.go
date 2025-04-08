package walker

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/elmersh/treew/internal/formatter"
)

// TreeOptions contiene todas las opciones para recorrer el árbol de directorios
type TreeOptions struct {
	Path              string
	ExcludeFolders    []string
	ExcludeExtensions []string
	ShowHidden        bool
	ShowFileSize      bool
	ShowLastModified  bool
	MaxDepth          int
	OutputFile        string
}

// Walker maneja el recorrido del árbol de directorios
type Walker struct {
	options   *TreeOptions
	formatter *formatter.Formatter
	output    []string
	file      *os.File
}

// NewWalker crea una nueva instancia del recorredor de directorios
func NewWalker(options *TreeOptions) (*Walker, error) {
	var file *os.File
	var err error

	// Abrir archivo de salida si se especificó
	if options.OutputFile != "" {
		file, err = os.Create(options.OutputFile)
		if err != nil {
			return nil, fmt.Errorf("error al crear archivo de salida: %w", err)
		}
	}

	return &Walker{
		options:   options,
		formatter: formatter.NewFormatter(options.ShowFileSize, options.ShowLastModified),
		output:    []string{},
		file:      file,
	}, nil
}

// Walk recorre el árbol de directorios y genera la salida
func (w *Walker) Walk() error {
	// Verificar si la ruta existe
	fileInfo, err := os.Stat(w.options.Path)
	if err != nil {
		return fmt.Errorf("error al acceder a la ruta: %w", err)
	}

	if !fileInfo.IsDir() {
		return fmt.Errorf("la ruta especificada no es un directorio")
	}

	// Formatear y mostrar la raíz
	rootLine := w.formatter.FormatRoot(w.options.Path, fileInfo.ModTime())
	w.printLine(rootLine)

	// Iniciar recorrido recursivo
	err = w.walkDir(w.options.Path, "", 0)
	if err != nil {
		return err
	}

	// Cerrar archivo si estaba abierto
	if w.file != nil {
		return w.file.Close()
	}

	return nil
}

// walkDir recorre un directorio recursivamente
func (w *Walker) walkDir(path, indent string, depth int) error {
	// Verificar profundidad máxima
	if w.options.MaxDepth != -1 && depth >= w.options.MaxDepth {
		return nil
	}

	// Leer entradas del directorio
	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("error al leer directorio %s: %w", path, err)
	}

	// Filtrar entradas según configuración
	var filteredEntries []fs.DirEntry
	for _, entry := range entries {
		name := entry.Name()

		// Verificar si está oculto y debe excluirse
		if !w.options.ShowHidden && strings.HasPrefix(name, ".") {
			continue
		}

		// Verificar exclusiones de carpetas
		if entry.IsDir() && contains(w.options.ExcludeFolders, name) {
			continue
		}

		// Verificar exclusiones de extensiones
		if !entry.IsDir() {
			ext := filepath.Ext(name)
			if contains(w.options.ExcludeExtensions, ext) {
				continue
			}
		}

		filteredEntries = append(filteredEntries, entry)
	}

	// Ordenar entradas: primero directorios, luego archivos
	sort.Slice(filteredEntries, func(i, j int) bool {
		if filteredEntries[i].IsDir() && !filteredEntries[j].IsDir() {
			return true
		}
		if !filteredEntries[i].IsDir() && filteredEntries[j].IsDir() {
			return false
		}
		return filteredEntries[i].Name() < filteredEntries[j].Name()
	})

	// Procesar cada entrada
	for i, entry := range filteredEntries {
		isLast := i == len(filteredEntries)-1
		connector := "├── "
		if isLast {
			connector = "└── "
		}

		entryPath := filepath.Join(path, entry.Name())
		fileInfo, err := os.Stat(entryPath)
		if err != nil {
			// Continuar si hay un error al obtener información del archivo
			continue
		}

		if entry.IsDir() {
			// Es un directorio
			line := w.formatter.FormatFolder(entry.Name(), indent, connector, fileInfo.ModTime())
			w.printLine(line)

			// Calcular nueva indentación para los hijos
			var newIndent string
			if isLast {
				newIndent = indent + "    "
			} else {
				newIndent = indent + "│   "
			}

			// Recursión para subdirectorios
			err = w.walkDir(entryPath, newIndent, depth+1)
			if err != nil {
				return err
			}
		} else {
			// Es un archivo
			line := w.formatter.FormatFile(entry.Name(), indent, connector, fileInfo.Size(), fileInfo.ModTime())
			w.printLine(line)
		}
	}

	return nil
}

// printLine imprime una línea en la salida
func (w *Walker) printLine(line string) {
	fmt.Println(line)
	w.output = append(w.output, line)

	if w.file != nil {
		fmt.Fprintln(w.file, line)
	}
}

// GetOutput devuelve todas las líneas de salida
func (w *Walker) GetOutput() []string {
	return w.output
}

// contains verifica si un slice contiene un elemento
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
