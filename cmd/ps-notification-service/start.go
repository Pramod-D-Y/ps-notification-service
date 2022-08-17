package main

import (
	"github.com/roppenlabs/ps-notification-service/internal/server"
	"github.com/spf13/cobra"
)

func initCLI() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "ps-user-service",
		Short: "A service used for managing teams in rapido",
	}

	rootCmd.AddCommand(startCommand())
	return rootCmd

}

func startCommand() *cobra.Command {
	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "Starts the server",
		Run: func(cmd *cobra.Command, args []string) {
			server.StartServer()
		},
	}
	return startCmd
}