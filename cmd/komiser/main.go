package main

import (
	"fmt"
	"github.com/flemzord/komiser/src/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	Version   = "develop"
	BuildDate = "-"
	Commit    = "-"

	root = &cobra.Command{
		Use:               "komiser",
		Short:             "komiser server",
		DisableAutoGenTag: true,
	}
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Get version",
		Run:   PrintVersion,
	}
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Launch Server",
		Run:   server.Server,
	}
)

func PrintVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("Version: %s \n", Version)
	fmt.Printf("Date: %s \n", BuildDate)
	fmt.Printf("Commit: %s \n", Commit)
}

func main() {
	viper.SetDefault("version", Version)

	root.AddCommand(versionCmd)
	root.AddCommand(serverCmd)

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
