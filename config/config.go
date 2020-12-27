package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Name string
}

func InitCfg(cfg string) {
	var c = Config{Name: cfg}

	err := c.initViper()
	if err != nil {
		log.Fatalf("init viper error: %v", err)
	}
}

func (c *Config) initViper() error {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	//viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	//viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}
	return nil
}
