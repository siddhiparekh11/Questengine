package models


type Database struct {
	Host string `mapstructure:"host"`
	Port string	`mapstructure:"port"`
	User string `mapstructure:"user"`
	Password string `mapstructure:"passord"`
	Name string `mapstructure:"name"`
}