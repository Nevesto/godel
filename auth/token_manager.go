package auth

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func SaveToken(newToken string) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Error determining config directory:", err)
		return
	}

	configPath := filepath.Join(configDir, "godel")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err = os.Mkdir(configPath, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating config directory:", err)
			return
		}
	}

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configPath)

	configFilePath := filepath.Join(configPath, "config.json")
	err = viper.ReadInConfig()
	if err == nil {
		currentToken := viper.GetString("discord_token")
		if currentToken == newToken {
			fmt.Println("Token is already saved.")
			return
		}
	}

	viper.Set("discord_token", newToken)

	if err := viper.WriteConfig(); err != nil {
		if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
			err = viper.WriteConfigAs(configFilePath)
			if err != nil {
				fmt.Println("Error creating new config file:", err)
				return
			}
		} else {
			fmt.Println("Error saving token to file:", err)
			return
		}
	}

	fmt.Println("New token saved successfully!")
}
