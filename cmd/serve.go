package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"log"
	"net"
	"net/http"
	"strconv"

	mux "github.com/gorilla/mux"
	tcpcheck "github.com/openshift-examples/container-helper/pkg/tcpcheck"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts a simple webserver",
	Long: `Starts a simple webserver
	long description`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Serving on %d", viper.GetInt("port"))
		router := mux.NewRouter()

		router.HandleFunc("/", httpRootHandler).Methods("GET")
		router.HandleFunc("/metrics", metrics).Methods("GET")
		router.HandleFunc("/check", tcpcheck.HttpCheck).Methods("GET")
		router.HandleFunc("/v1/tcpchecks", tcpcheck.GetTcpChecks).Methods("GET")
		// 	router.HandleFunc("/v1/tcpchecks/{uuid}", tcpcheck.getTcpCheckByUuid).Methods("GET")
		// 	router.HandleFunc("/v1/tcpchecks", tcpcheck.createTcpCheck).Methods("POST")
		// 	router.HandleFunc("/v1/tcpchecks/{uuid}", tcpcheck.deleteTcpCheck).Methods("DELETE")
		srv := &http.Server{
			Handler: logRequest(router),
			Addr:    fmt.Sprintf(":%d", viper.GetInt("port")),
			// Good practice: enforce timeouts for servers you create!
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
		log.Fatal(srv.ListenAndServe())

	},
}

// https://github.com/gorilla/mux/issues/444#issuecomment-459090877
func metrics(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
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
	serveCmd.Flags().StringP("response", "r", "Hello OpenShift!", "response message")
	serveCmd.Flags().StringP("namespace", "n", "Unknown", "Namespace")

	viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
	viper.BindPFlag("response", serveCmd.Flags().Lookup("response"))
	viper.BindPFlag("namespace", serveCmd.Flags().Lookup("namespace"))

	viper.SetDefault("port", 8080)
	viper.SetDefault("response", "Hello OpenShift!")
	viper.SetDefault("namespace", "Unknown")

	// viper.SetEnvPrefix("CT_SERVE") // Set the environment prefix to DEMO_*
	viper.AutomaticEnv() // Automatically search for environment variables
}
