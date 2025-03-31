package auth

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

func Connect() (*discordgo.Session, error) {
	activeTokenName := viper.GetString("active_token")
	if activeTokenName == "" {
		return nil, fmt.Errorf("token ativo não definido. Por favor, execute 'godel token-switch [nome]' para definir o token a ser usado")
	}

	tokens := viper.GetStringMapString("tokens")
	token, exists := tokens[activeTokenName]
	if !exists || token == "" {
		return nil, fmt.Errorf("token para a conta ativa '%s' não encontrado. Por favor, execute 'godel set-token [nome] [token]' para registrar o token", activeTokenName)
	}

	dg, err := discordgo.New(token)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar sessão do Discord: %v", err)
	}

	dg.Identify.Intents = discordgo.IntentsAll
	dg.Identify.Intents = discordgo.IntentsDirectMessages

	dg.Token = token
	dg.Identify.Token = token
	dg.StateEnabled = false
	dg.Identify.Compress = false
	dg.Identify.LargeThreshold = 0

	err = dg.Open()
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao Discord: %v", err)
	}

	fmt.Println("Conectado como: " + dg.State.User.Username)
	return dg, nil
}
