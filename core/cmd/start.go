package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "nicedice",
	Short: "Tri-topora core",
	Long: "Too start please define config file",
	RunE: start,
}

var cfgFile string

func init() {
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Path to config file")
	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
}

func Execute() {
	RootCmd.Execute()
}

func start(cmd *cobra.Command, args []string) error {
    return nil
}
