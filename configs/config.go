package configs

import "github.com/spf13/viper"

var cfg *Config

type Config struct {
	API   APIConfig
	DB    DBConfig
	PRICE UDrivePriceConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

type UDrivePriceConfig struct {
	Host string
	Port string
}

func init() {
	viper.SetDefault("api.port", "3200")
	viper.SetDefault("database.host", "udrive-service.cy0mf1jtohal.us-east-2.rds.amazonaws.com")
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
	cfg = new(Config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	cfg.PRICE = UDrivePriceConfig{
		Host: viper.GetString("udrive-price.host"),
		Port: viper.GetString("udrive-price.port"),
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetAPI() string {
	return cfg.API.Port
}

func GetPRICE() UDrivePriceConfig {
	return cfg.PRICE
}
