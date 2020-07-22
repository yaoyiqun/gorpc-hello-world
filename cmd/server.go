package cmd

import (
	"github.com/spf13/cobra"
	"grpc-hello-world/server"
	"log"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC hello-world server",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover error:%v\n", err)
			}
		}()
		server.Server()
	},
}

func init() {
	serverCmd.Flags().StringVarP(&server.ServerPort, "port", "p", "50052", "server prot")
	serverCmd.Flags().StringVarP(&server.CertPemPath, "cert-pem", "", "./certs/server.pem", "cert pem path")
	serverCmd.Flags().StringVarP(&server.CertKeyPath, "cert-key", "", "./certs/server.key", "cert key path")
	serverCmd.Flags().StringVarP(&server.CertName, "cert-name", "", "grpc.abc", "server's hostname")
	rootCmd.AddCommand(serverCmd)
}
