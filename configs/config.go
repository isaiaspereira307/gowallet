package configs

import "github.com/spf13/viper"

var (
	cfg    *config
	logger *Logger
)

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port      string
	JwtSecret string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}
	cfg = new(config)
	cfg.API = APIConfig{
		Port:      viper.GetString("api.port"),
		JwtSecret: viper.GetString("api.jwt_secret"),
	}
	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}
	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}

func GetJwtSecret() string {
	return cfg.API.JwtSecret
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
