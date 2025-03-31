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
		fmt.Println("Erro ao determinar o diretório de configuração:", err)
		return
	}

	configPath := filepath.Join(configDir, "godel")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err = os.Mkdir(configPath, os.ModePerm)
		if err != nil {
			fmt.Println("Erro ao criar diretório de configuração:", err)
			return
		}
	}

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configPath)

	configFilePath := filepath.Join(configPath, "config.json")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Nenhuma configuração existente encontrada, criando uma nova.")
	}

	tokens := viper.GetStringMapString("tokens")
	if tokens == nil {
		tokens = make(map[string]string)
	}

	if current, exists := tokens[tokenName]; exists && current == newToken {
		fmt.Println("Token já está registrado com esse nome e valor.")
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
				fmt.Println("Erro ao criar novo arquivo de configuração:", err)
				return
			}
		} else {
			fmt.Println("Erro ao salvar token no arquivo:", err)
			return
		}
	}

	fmt.Println("Token salvo com sucesso!")
}
