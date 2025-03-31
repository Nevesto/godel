package auth

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func SaveToken(tokenName, newToken string) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Error determining configuration directory:", err)
		return
	}

	configPath := filepath.Join(configDir, "godel")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err = os.Mkdir(configPath, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating configuration directory:", err)
			return
		}
	}

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configPath)

	configFilePath := filepath.Join(configPath, "config.json")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("No existing configuration found, creating a new one.")
	}

	tokens := viper.GetStringMapString("tokens")
	if tokens == nil {
		tokens = make(map[string]string)
	}

	if current, exists := tokens[tokenName]; exists && current == newToken {
		fmt.Println("Token is already registered with this name and value.")
		return
	}

	tokens[tokenName] = newToken
	viper.Set("tokens", tokens)

	if viper.GetString("active_token") == "" {
		viper.Set("active_token", tokenName)
	}

	if err := viper.WriteConfig(); err != nil {
		if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
			err = viper.WriteConfigAs(configFilePath)
			if err != nil {
				fmt.Println("Error creating new configuration file:", err)
				return
			}
		} else {
			fmt.Println("Error saving token to file:", err)
			return
		}
	}

	fmt.Println("Token saved successfully!")
}
