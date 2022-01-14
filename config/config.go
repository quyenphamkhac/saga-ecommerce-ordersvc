package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	PostgresDB
}

type PostgresDB struct {
	Addr     string `mapstructure:"ORDERSVC_POSTGRESDB_ADDR"`
	Password string `mapstructure:"ORDERSVC_POSTGRESDB_PASSWORD"`
	User     string `mapstructure:"ORDERSVC_POSTGRESDB_USER"`
	Database string `mapstructure:"ORDERSVC_POSTGRESDB_DATABASE"`
}

func NewServiceConfig() (*Config, error) {
	v, err := loadConfig()
	if err != nil {
		return nil, err
	}
	c, err := parseConfig(v)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func loadConfig() (*viper.Viper, error) {
	v := viper.New()
	v.SetEnvPrefix("ordersvc")
	v.AutomaticEnv()
	return v, nil
}
func parseConfig(v *viper.Viper) (*Config, error) {
	c := &Config{
		PostgresDB: PostgresDB{
			Addr:     v.GetString("postgresdb_addr"),
			User:     v.GetString("postgresdb_user"),
			Database: v.GetString("postgresdb_database"),
			Password: v.GetString("postgresdb_password"),
		},
	}
	return c, nil
}
