/*
Copyright © 2026 Codurance
*/
package cmd

import (
	"context"

	"github.com/CoduranceExperiments/go-test-skeleton/api"
	"github.com/CoduranceExperiments/go-test-skeleton/api/routes"
	"github.com/spf13/cobra"
)

var port int

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the API server",
	Long:  `Runs the Codurance Test API server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		v1, err := routes.New(routes.RouteVersionOne)
		if err != nil {
			return err
		}
		opts := []api.Opt{
			api.WithPort(port),
			api.WithRouter(v1.Router()),
		}
		srv := api.New(opts...)
		return srv.Serve(context.Background())
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&port, "port", "p", 3000, "Print this help message")
}
