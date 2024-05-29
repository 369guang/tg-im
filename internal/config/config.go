package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	DEBUG  bool `mapstructure:"debug"`
	Common struct {
		StoreMessageHistory     bool   `mapstructure:"store_message_history"`
		StoreMessageHistoryDays int    `mapstructure:"store_message_history_days"`
		StoreOfflineMessage     bool   `mapstructure:"store_offline_message"`
		SecretKey               string `mapstructure:"secret_key"`
	} `mapstructure:"common"`

	Server struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
		Id   string `mapstructure:"id"`
	} `mapstructure:"server"`

	Rpc struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
		Name string `mapstructure:"name"`
		// 多个etcd
		Etcd []string `mapstructure:"etcd"`
	} `mapstructure:"rpc"`

	Tls struct {
		CertFile string `mapstructure:"cert_file"`
		KeyFile  string `mapstructure:"key_file"`
	} `mapstructure:"tls"`

	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DbName   string `mapstructure:"db_name"`
	} `mapstructure:"database"`

	Cache struct { // redis cluster
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
	} `mapstructure:"cache"`

	JWT struct {
		Secret string `mapstructure:"secret"`
		Expire int    `mapstructure:"expire"`
	} `mapstructure:"jwt"`

	Logs struct {
		Directory  string `mapstructure:"directory"`
		FileName   string `mapstructure:"file_name"`
		ToFile     bool   `mapstructure:"to_file"`
		Level      string `mapstructure:"level"`
		MaxSize    int    `mapstructure:"max_size"`
		MaxAge     int    `mapstructure:"max_age"`
		MaxBackups int    `mapstructure:"max_backups"`
		Compress   bool   `mapstructure:"compress"`
	} `mapstructure:"logs"`
}

func LoadConfig(path, configName string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")

	var config Config
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
