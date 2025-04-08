package cmd

import (
	"os"
	"path/filepath"

	"github.com/elmersh/treew/internal/config"
	"github.com/elmersh/treew/internal/walker"
	"github.com/elmersh/treew/pkg/emoji"
	"github.com/spf13/cobra"
)

var (
	// Opciones de línea de comandos
	cfgFile           string
	excludeFolders    []string
	excludeExtensions []string
	showHidden        bool
	showFileSize      bool
	showLastModified  bool
	maxDepth          int
	outputFile        string
	saveConfig        bool
)

// rootCmd representa el comando base
var rootCmd = &cobra.Command{
	Use:   "treew [path]",
	Short: "Muestra un árbol de directorios con íconos",
	Long: `Treew muestra un árbol de directorios con íconos para diferentes tipos de archivos.
Permite filtrar carpetas y extensiones específicas, mostrar tamaños de archivos
y fechas de modificación, y guardar la salida en un archivo.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}

		// Convertir a ruta absoluta
		absPath, err := filepath.Abs(path)
		if err != nil {
			return err
		}

		// Cargar configuración
		cfg, err := config.LoadConfig()
		if err != nil {
			return err
		}

		// Inicializar el paquete emoji con la configuración
		emoji.InitFromConfig(cfg)

		// Usar valores de línea de comandos si se proporcionaron
		if cmd.Flags().Changed("exclude-folders") {
			cfg.ExcludeFolders = excludeFolders
		}
		if cmd.Flags().Changed("exclude-extensions") {
			cfg.ExcludeExtensions = excludeExtensions
		}
		if cmd.Flags().Changed("show-hidden") {
			cfg.ShowHidden = showHidden
		}
		if cmd.Flags().Changed("show-file-size") {
			cfg.ShowFileSize = showFileSize
		}
		if cmd.Flags().Changed("show-last-modified") {
			cfg.ShowLastModified = showLastModified
		}
		if cmd.Flags().Changed("max-depth") {
			cfg.MaxDepth = maxDepth
		}

		// Guardar configuración si se solicitó
		if saveConfig {
			if err := config.SaveConfig(cfg); err != nil {
				return err
			}
		}

		// Configurar opciones para el recorrido
		options := &walker.TreeOptions{
			Path:              absPath,
			ExcludeFolders:    cfg.ExcludeFolders,
			ExcludeExtensions: cfg.ExcludeExtensions,
			ShowHidden:        cfg.ShowHidden,
			ShowFileSize:      cfg.ShowFileSize,
			ShowLastModified:  cfg.ShowLastModified,
			MaxDepth:          cfg.MaxDepth,
			OutputFile:        outputFile,
		}

		// Crear walker y ejecutar
		w, err := walker.NewWalker(options)
		if err != nil {
			return err
		}

		return w.Walk()
	},
}

// Execute añade todos los comandos secundarios al comando raíz y establece las banderas apropiadamente.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Cargar configuración predeterminada
	defaultCfg := config.DefaultConfig()

	// Definir banderas
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "archivo de configuración (predeterminado: $HOME/.config/treew.yaml)")
	rootCmd.Flags().StringSliceVar(&excludeFolders, "exclude-folders", defaultCfg.ExcludeFolders, "carpetas a excluir")
	rootCmd.Flags().StringSliceVar(&excludeExtensions, "exclude-extensions", defaultCfg.ExcludeExtensions, "extensiones a excluir")
	rootCmd.Flags().BoolVarP(&showHidden, "show-hidden", "a", defaultCfg.ShowHidden, "mostrar archivos ocultos")
	rootCmd.Flags().BoolVarP(&showFileSize, "show-file-size", "s", defaultCfg.ShowFileSize, "mostrar tamaño de archivos")
	rootCmd.Flags().BoolVar(&showLastModified, "show-last-modified", defaultCfg.ShowLastModified, "mostrar fecha de modificación")
	rootCmd.Flags().IntVarP(&maxDepth, "max-depth", "d", defaultCfg.MaxDepth, "profundidad máxima (-1 para ilimitado)")
	rootCmd.Flags().StringVar(&outputFile, "output-file", "", "guardar salida en archivo")
	rootCmd.Flags().BoolVar(&saveConfig, "save-config", false, "guardar configuración actual como predeterminada")
}
