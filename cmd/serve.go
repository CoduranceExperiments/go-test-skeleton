/*
Copyright © 2026 Codurance
*/
package cmd

import (
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
		router, err := routes.NewRouter(routes.All()...)
		if err != nil {
			return err
		}

		srv, err := api.New(
			api.WithPort(port),
			api.WithHandler(router),
		)
		if err != nil {
			return err
		}

		// cmd.Context() is currently never cancelled. See the TODO on
		// api.RestService.Serve.
		return srv.Serve(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&port, "port", "p", 3000, "Port the API server listens on")
}
