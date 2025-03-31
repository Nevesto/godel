package scripts

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func SwitchToken(tokenName string) error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("could not determine user config directory: %w", err)
	}

	configPath := filepath.Join(configDir, "godel")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return fmt.Errorf("configuration directory not found. Execute 'godel set-token [name] [token]' first")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("could not read configuration file: %w", err)
	}

	tokens := viper.GetStringMapString("tokens")
	if tokens == nil || len(tokens) == 0 {
		return fmt.Errorf("no tokens found. Register a token with 'godel set-token [name] [token]'")
	}

	if _, exists := tokens[tokenName]; !exists {
		return fmt.Errorf("token '%s' not found", tokenName)
	}

	viper.Set("active_token", tokenName)

	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("error saving configuration: %w", err)
	}

	return nil
}
