/*
Copyright © 2024 Silicon Labs
*/
package cmd

import (
	"fmt"
	"os"
	"silabs/get-zap/github"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const ownerArg = "owner"
const repoArg = "repo"

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "get-zap",
	Short: "Application to retrieve artifacts from github.",
	Long:  `This application by default retrieves zap artifacts, with the right arguments, it can be used to retrieve assets from any public github repo.`,
	Run: func(cmd *cobra.Command, args []string) {
		owner, err := cmd.Flags().GetString(ownerArg)
		cobra.CheckErr(err)
		repo, err := cmd.Flags().GetString(repoArg)
		cobra.CheckErr(err)
		github.DownloadLatestRelease(owner, repo)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.get-zap.yaml)")
	rootCmd.PersistentFlags().StringP(ownerArg, "o", "project-chip", "Owner of the github repository.")
	rootCmd.PersistentFlags().StringP(repoArg, "r", "zap", "Name of the github repository.")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".get-zap" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".get-zap")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
