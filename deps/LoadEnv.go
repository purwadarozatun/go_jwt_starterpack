package deps

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"PORT"`

	ClientOrigin       string `mapstructure:"CLIENT_ORIGIN"`
	LDAP_BASE_DN       string `mapstructure:"LDAP_BASE_DN"`
	LDAP_URL           string `mapstructure:"LDAP_URL"`
	LDAP_PORT          int    `mapstructure:"LDAP_PORT"`
	LDAP_BIND_DN       string `mapstructure:"LDAP_BIND_DN"`
	LDAP_BIND_PASSWORD string `mapstructure:"LDAP_BIND_PASSWORD"`

	AUTH_LDLAP string `mapstructure:"LDAP_BIND_PASSWORD"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
