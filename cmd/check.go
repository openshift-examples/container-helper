package cmd

import (
	"fmt"
	"log"
	"time"

	tcpcheck "github.com/openshift-examples/container-helper/pkg/tcpcheck"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check",
	Long:  `check test`,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()

		tcpcheck.CheckAll()
		// fmt.Printf("\n\nHier: %v\n", tcpcheck.Get())
		for _, tc := range tcpcheck.Get() {
			fmt.Printf("%-25s : %d\n", tc.Name, tc.Available)
		}
		elapsed := time.Since(start)
		log.Printf("Checks took %s", elapsed)

	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// serveCmd.Flags().IntP("port", "p", 8080, "port on which the server will listen")
	// serveCmd.Flags().StringP("response", "r", "Hello OpenShift!", "response message")

	// viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
	// viper.BindPFlag("port", serveCmd.Flags().Lookup("response"))

	// viper.SetDefault("port", 8080)
	// viper.SetDefault("response", "Hello OpenShift!")

	// // viper.SetEnvPrefix("CT_SERVE") // Set the environment prefix to DEMO_*
	// viper.AutomaticEnv() // Automatically search for environment variables
}
