package scripts

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/spf13/viper"
)

func ListTokens() ([]string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("could not determine user config directory: %w", err)
	}

	configPath := filepath.Join(configDir, "godel")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("configuration directory not found. Register a token first with 'godel set-token [name] [token]'")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("could not read configuration file: %w", err)
	}

	tokens := viper.GetStringMapString("tokens")
	if tokens == nil || len(tokens) == 0 {
		return nil, fmt.Errorf("no tokens registered")
	}

	var aliases []string
	for alias := range tokens {
		aliases = append(aliases, alias)
	}

	sort.Strings(aliases)

	return aliases, nil
}
