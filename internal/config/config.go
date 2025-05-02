package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port string
}

// Retorna un objeto Config que contiene la configuración de la aplicación, esto aplica que otras partes del programa puedan acceder a los valores de configuración
func LoadConfig() *Config {
	viper.SetConfigName("config")
	//Se especifica el tipo de configuración que se usara
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./web/")

	//Verificar si el archivo de configuración es leido
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("No se pudo leer el archivo de configuración: %s", err)
	}

	cfg := &Config{
		Port: viper.GetString("port"),
	}

	return cfg
}
