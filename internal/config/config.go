package config

import (
	"github.com/spf13/viper"
	"strings"
	"time"
)

type (
	Config struct {
		Server ServerConfig
		Auth   AuthConfig
		DB     PostgresConfig
	}

	ServerConfig struct {
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}

	AuthConfig struct {
		AccessTokenTTL  time.Duration `mapstrcuture:"accessTokenTTL"`
		RefreshTokenTTL time.Duration `mapstrcuture:"refreshTokenTTL"`
		SigningKey      string
		PasswordSalt    string
	}

	PostgresConfig struct {
		Host     string `mapstrcuture:"host"`
		Port     string `mapstrcuture:"port"`
		Username string `mapstrcuture:"username"`
		Password string
		DBName   string `mapstrcuture:"dbname"`
		SSLMode  string `mapstrcuture:"sslmode"`
	}
)

func Init(configPath string) (*Config, error) {
	if err := parseConfigPath(configPath); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}
	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseConfigPath(filepath string) error {
	path := strings.Split(filepath, "/")

	viper.AddConfigPath(path[0])
	viper.SetConfigName(path[1])

	return viper.ReadInConfig()
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.Server); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("auth", &cfg.Auth); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("postgres", &cfg.DB); err != nil {
		return err
	}
	return nil
}

func parseEnv(cfg *Config) error {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	cfg.DB.Password = viper.GetString("POSTGRES_PASSWORD")
	cfg.Auth.SigningKey = viper.GetString("SIGNING_KEY")
	cfg.Auth.PasswordSalt = viper.GetString("PASSWORD_SALT")
	return nil
}
