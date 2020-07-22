package cmd

import (
	"github.com/spf13/cobra"
	"grpc-hello-world/client"
	"log"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Run the gRPC hello-world client",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover error:#{err}\n")
			}
		}()
		client.Client()
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
