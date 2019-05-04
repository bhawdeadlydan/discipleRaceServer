package cmd



import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/discipleRaceServer/config"
	"fmt"
)

var (
	cfgFile         string
	devmode         bool
	cfg             *viper.Viper
	appConfig       *config.Config
	initConfigError error
)

var RootCmd = &cobra.Command{
	Use:   "discipleRaceServer --config config/app.conf",
	Short: "discipleRaceServer is a gRPC API backend.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if initConfigError != nil {
			return initConfigError
		}

		fmt.Println("Hello World!")
		fmt.Println("Read the port", appConfig.GetAppPort())

		return nil
	},
}

func Execute() error {
	if err := RootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file (defaults to config/app.json)")
	RootCmd.PersistentFlags().BoolVarP(&devmode, "devmode", "d", false, "Enables development mode")
	RootCmd.PersistentFlags().Int("app-port", 6076, "Port for disciple race service to listen on")

	RootCmd.MarkFlagRequired("config")
}

func initConfig() {
	cfg = viper.New()
	if cfgFile != "" { // enable ability to specify config file via flag
		cfg.SetConfigFile(cfgFile)
	}

	cfg.AutomaticEnv()        // read in environment variables that match
	cfg.SetConfigType("json") // Set config type to json as it has the suffix .conf instead of .json

	// If a config file is found, read it in.
	if err := cfg.ReadInConfig(); err != nil {
		initConfigError = err
	}
	cfg.BindPFlag("APP_PORT", RootCmd.Flags().Lookup("app-port"))

	appConfig, _ = config.New(cfg)
}
