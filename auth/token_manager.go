package auth

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// SaveToken salva (ou atualiza) um token identificado por um nome (alias).
// Se não houver token ativo, define o token salvo como ativo.
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
		// Não existe arquivo de configuração ainda
		fmt.Println("Nenhuma configuração existente encontrada, criando uma nova.")
	}

	// Recupera o mapa de tokens salvos (caso exista)
	tokens := viper.GetStringMapString("tokens")
	if tokens == nil {
		tokens = make(map[string]string)
	}

	// Se o token com o mesmo nome já existir e for idêntico, não há necessidade de atualizá-lo.
	if current, exists := tokens[tokenName]; exists && current == newToken {
		fmt.Println("Token já está registrado com esse nome e valor.")
		return
	}

	// Atualiza ou adiciona o token no mapa
	tokens[tokenName] = newToken
	viper.Set("tokens", tokens)

	// Se não houver token ativo definido, define este token como ativo
	if viper.GetString("active_token") == "" {
		viper.Set("active_token", tokenName)
	}

	// Salva as configurações no arquivo
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
