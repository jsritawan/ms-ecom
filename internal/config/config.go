package config

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server ServerConfig `yaml:"server"`
		DB     DBConfig     `mapstructure:"database"`
		Logger LoggerConfig `yaml:"logger"`
	}

	ServerConfig struct {
		Port string `yaml:"port,omitempty"`
	}

	DBConfig struct {
		Host         string        `yaml:"host,omitempty"`
		Port         string        `yaml:"port,omitempty"`
		Username     string        `yaml:"username,omitempty"`
		Password     string        `yaml:"password,omitempty"`
		Name         string        `yaml:"name,omitempty"`
		SSLMode      string        `yaml:"sslmode,omitempty"`
		Timezone     string        `yaml:"timezone,omitempty"`
		MaxIdleConns int           `yaml:"maxIdleConns,omitempty"`
		MaxIdleTime  time.Duration `yaml:"maxIdleTime,omitempty"`
		MaxOpenConns int           `yaml:"maxOpenConns,omitempty"`
		MaxLifeTime  time.Duration `yaml:"maxLifeTime,omitempty"`
		SSLCert      string        `yaml:"sslcert,omitempty"`
		SSLKey       string        `yaml:"sslkey,omitempty"`
		SSLRootCert  string        `yaml:"sslrootcert,omitempty"`
	}

	LoggerConfig struct {
		Level       string `yaml:"level"`
		OnCloud     bool   `yaml:"oncloud"`
		Development bool   `yaml:"development"`
		Stacktrace  bool   `yaml:"stacktrace"`
		Caller      bool   `yaml:"caller"`
		DbLevel     string `yaml:"dblevel"`
	}
)

const (
	defaultConfigName = "config"
	defaultConfigType = "yaml"
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
