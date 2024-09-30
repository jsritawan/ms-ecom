package config

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server ServerConfig `yaml:"server"`
	}

	ServerConfig struct {
		Port string `yaml:"port,omitempty"`
	}
)

const (
	defaultConfigName = "config"
	defaultConfigType = "yml"
	defaultConfigPath = "../../etc"
)

var (
	configPath string
	configName string
)

func init() {
	flag.StringVar(&configPath, "configPath", defaultConfigPath, "path to config file")
	flag.StringVar(&configName, "configName", defaultConfigName, "config file name")
}

func LoadConfig(path string) *Config {
	if path != "" {
		configPath = path
	}

	var cfg = &Config{}
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	filePath := filepath.Join(wd, configPath)
	if err := os.Chdir(filePath); err != nil {
		log.Fatalln(err)
	}

	viper.SetConfigName(configName)
	viper.SetConfigType(defaultConfigType)
	viper.AddConfigPath(filePath)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.MergeInConfig(); err != nil {
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalln(err)
	}

	return cfg
}
