/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"log"
	"net"
	"net/http"
	"strconv"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts a simple webserver",
	Long: `Starts a simple webserver
	long description`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Serving on %d",viper.GetInt("port"))
		http.HandleFunc("/", httpRootHandler)

		err := http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("port")), logRequest(http.DefaultServeMux))

		if err != nil {
			panic("ListenAndServe: " + err.Error())
		}
	},
}

func httpRootHandler(w http.ResponseWriter, r *http.Request) {

	// Echo back the port the request was received on
	// via a "request-port" header.
	addr := r.Context().Value(http.LocalAddrContextKey).(net.Addr)
	if tcpAddr, ok := addr.(*net.TCPAddr); ok {
		w.Header().Set("x-request-port", strconv.Itoa(tcpAddr.Port))
	}

	fmt.Fprintln(w, viper.GetString("response"))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request from %s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func init() {
	rootCmd.AddCommand(serveCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	serveCmd.Flags().IntP("port", "p", 8080, "port on which the server will listen")
	serveCmd.Flags().StringP("response","r", "Hello OpenShift!", "response message")

	viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
	viper.BindPFlag("port", serveCmd.Flags().Lookup("response"))

	viper.SetDefault("port", 8080)
	viper.SetDefault("response", "Hello OpenShift!")

	// viper.SetEnvPrefix("CT_SERVE") // Set the environment prefix to DEMO_*
	viper.AutomaticEnv() // Automatically search for environment variables
}
