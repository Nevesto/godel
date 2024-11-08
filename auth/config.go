package auth

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func LoadConfig() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println(err)
		return
	}

	configPath := filepath.Join(configDir, "godel")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configPath)

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
	}
}
