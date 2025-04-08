package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Config almacena la configuración de la aplicación
type Config struct {
	ExcludeFolders    []string `mapstructure:"exclude_folders"`
	ExcludeExtensions []string `mapstructure:"exclude_extensions"`
	ShowHidden        bool     `mapstructure:"show_hidden"`
	ShowFileSize      bool     `mapstructure:"show_file_size"`
	ShowLastModified  bool     `mapstructure:"show_last_modified"`
	MaxDepth          int      `mapstructure:"max_depth"`
	UseNerdFonts      bool     `mapstructure:"use_nerd_fonts"`
}

// DefaultConfig devuelve la configuración predeterminada
func DefaultConfig() *Config {
	return &Config{
		ExcludeFolders:    []string{"node_modules", "bin", "obj", ".git", "packages"},
		ExcludeExtensions: []string{},
		ShowHidden:        false,
		ShowFileSize:      false,
		ShowLastModified:  false,
		MaxDepth:          -1, // -1 indica sin límite
		UseNerdFonts:      false,
	}
}

// LoadConfig carga la configuración del archivo o utiliza los valores predeterminados
func LoadConfig() (*Config, error) {
	config := DefaultConfig()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return config, err
	}

	// Configurar viper
	viper.SetConfigName("treew")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath(filepath.Join(homeDir, ".config"))

	// Hacer que Viper sea insensible a mayúsculas/minúsculas
	viper.SetEnvPrefix("TREEW")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Definir valores predeterminados
	viper.SetDefault("exclude_folders", config.ExcludeFolders)
	viper.SetDefault("exclude_extensions", config.ExcludeExtensions)
	viper.SetDefault("show_hidden", config.ShowHidden)
	viper.SetDefault("show_file_size", config.ShowFileSize)
	viper.SetDefault("show_last_modified", config.ShowLastModified)
	viper.SetDefault("max_depth", config.MaxDepth)
	viper.SetDefault("use_nerd_fonts", config.UseNerdFonts)

	// Leer archivo de configuración si existe
	if err := viper.ReadInConfig(); err != nil {
		// Está bien si no se encuentra el archivo de configuración
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return config, err
		}
	}

	// Decodificar en la estructura de configuración
	if err := viper.Unmarshal(config); err != nil {
		return config, err
	}

	return config, nil
}

// SaveConfig guarda la configuración actual en un archivo
func SaveConfig(config *Config) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configDir := filepath.Join(homeDir, ".config")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	viper.Set("exclude_folders", config.ExcludeFolders)
	viper.Set("exclude_extensions", config.ExcludeExtensions)
	viper.Set("show_hidden", config.ShowHidden)
	viper.Set("show_file_size", config.ShowFileSize)
	viper.Set("show_last_modified", config.ShowLastModified)
	viper.Set("max_depth", config.MaxDepth)
	viper.Set("use_nerd_fonts", config.UseNerdFonts)

	return viper.WriteConfigAs(filepath.Join(configDir, "treew.yaml"))
}
