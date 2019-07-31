package flyctl

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/superfly/flyctl/terminal"
)

const (
	ConfigAPIAccessToken = "api_access_token"
	ConfigAPIBaseURL     = "api_base_url"
	ConfigAppName        = "app"
	ConfigVerboseOutput  = "verbose"
)

func InitConfig() {
	viper.SetDefault(ConfigAPIBaseURL, "https://fly.io")

	viper.SetEnvPrefix("FLY")
	viper.AutomaticEnv()

	if viper.GetBool("verbose") {
		terminal.SetLogLevel(terminal.LevelDebug)
	}

	// if configFile != "" {
	// 	viper.SetConfigFile(configFile)
	// } else {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	viper.SetConfigName("fly")
	viper.AddConfigPath(cwd)
	// }

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		terminal.Debug("Using config file:", viper.ConfigFileUsed())
	}

	if token, err := GetSavedAccessToken(); err == nil {
		viper.SetDefault(ConfigAPIAccessToken, token)
	}
}