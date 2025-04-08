package formatter

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/elmersh/treew/internal/emoji"
	"github.com/fatih/color"
)

// Formatter gestiona el formato de salida del árbol de directorios
type Formatter struct {
	ShowFileSize     bool
	ShowLastModified bool
}

// NewFormatter crea una nueva instancia del formateador
func NewFormatter(showFileSize, showLastModified bool) *Formatter {
	return &Formatter{
		ShowFileSize:     showFileSize,
		ShowLastModified: showLastModified,
	}
}

// FormatRoot formatea la línea raíz del árbol
func (f *Formatter) FormatRoot(path string, modTime time.Time) string {
	folderName := filepath.Base(path)

	// Usar color para el nombre de la carpeta raíz
	boldCyan := color.New(color.FgCyan, color.Bold).SprintFunc()

	output := fmt.Sprintf("%s %s", emoji.FolderIcon(), boldCyan(folderName))

	if f.ShowLastModified {
		output += fmt.Sprintf(" (Modified: %s)", modTime.Format("2006-01-02 15:04:05"))
	}

	return output
}

// FormatFolder formatea una línea de carpeta
func (f *Formatter) FormatFolder(name string, indent string, connector string, modTime time.Time) string {
	// Usar color para el nombre de la carpeta
	cyan := color.New(color.FgCyan).SprintFunc()

	output := fmt.Sprintf("%s%s%s %s", indent, connector, emoji.FolderIcon(), cyan(name))

	if f.ShowLastModified {
		output += fmt.Sprintf(" (Modified: %s)", modTime.Format("2006-01-02 15:04:05"))
	}

	return output
}

// FormatFile formatea una línea de archivo
func (f *Formatter) FormatFile(name string, indent string, connector string, size int64, modTime time.Time) string {
	ext := filepath.Ext(name)
	fileIcon := emoji.GetFileIcon(ext)

	output := fmt.Sprintf("%s%s%s %s", indent, connector, fileIcon, name)

	if f.ShowFileSize {
		output += fmt.Sprintf(" (%s)", FormatFileSize(size))
	}

	if f.ShowLastModified {
		output += fmt.Sprintf(" (Modified: %s)", modTime.Format("2006-01-02 15:04:05"))
	}

	return output
}

// FormatFileSize formatea un tamaño de archivo a un formato legible
func FormatFileSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}

	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.2f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}
