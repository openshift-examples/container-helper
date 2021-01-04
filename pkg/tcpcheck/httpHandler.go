package tcpcheck

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

// https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html
// To run:
// go run main.go

type Error struct {
	Message string `json:"Message"`
}

// var (
// 	tcpChecksValidator = validator.NewValidator()
// )

func GetTcpChecks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Printf("%+v\n", Checks)
	err := json.NewEncoder(w).Encode(Checks)
	if err != nil {
		e := Error{Message: "Internal Server Error"}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(e)
		return
	}
}

// func getTcpCheckByUuid(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for _, tcpcheck := range Checks {
// 		if tcpcheck.Uuid == params["uuid"] {
// 			_ = json.NewEncoder(w).Encode(tcpcheck)
// 			return
// 		}
// 	}
// 	e := Error{Message: "UUID Not Found"}
// 	w.WriteHeader(http.StatusNotFound)
// 	if err := json.NewEncoder(w).Encode(e); err != nil {
// 		e := Error{Message: "Internal Server Error"}
// 		http.Error(w, e.Message, http.StatusInternalServerError)
// 	}
// }

// func createTcpCheck(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	dec := json.NewDecoder(r.Body)
// 	dec.DisallowUnknownFields()
// 	var tc Data
// 	err := dec.Decode(&tc)
// 	if err != nil {
// 		var syntaxError *json.SyntaxError
// 		var unmarshalTypeError *json.UnmarshalTypeError

// 		switch {
// 		// Catch any syntax errors in the JSON and send an error message
// 		// which interpolates the location of the problem to make it
// 		// easier for the client to fix.
// 		case errors.As(err, &syntaxError):
// 			e := Error{Message: fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)}
// 			w.WriteHeader(http.StatusBadRequest)
// 			_ = json.NewEncoder(w).Encode(e)
// 		// In some circumstances Decode() may also return an
// 		// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
// 		// is an open issue regarding this at
// 		// https://github.com/golang/go/issues/25956.
// 		case errors.Is(err, io.ErrUnexpectedEOF):
// 			e := Error{Message: fmt.Sprintf("Request body contains badly-formed JSON")}
// 			w.WriteHeader(http.StatusBadRequest)
// 			_ = json.NewEncoder(w).Encode(e)
// 		// Catch any type errors, like trying to assign a string in the
// 		// JSON request body to a int field in our Person struct. We can
// 		// interpolate the relevant field name and position into the error
// 		// message to make it easier for the client to fix.
// 		case errors.As(err, &unmarshalTypeError):
// 			e := Error{Message: fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)}
// 			w.WriteHeader(http.StatusBadRequest)
// 			_ = json.NewEncoder(w).Encode(e)

// 		// Catch the error caused by extra unexpected fields in the request
// 		// body. We extract the field name from the error message and
// 		// interpolate it in our custom error message. There is an open
// 		// issue at https://github.com/golang/go/issues/29035 regarding
// 		// turning this into a sentinel error.
// 		case strings.HasPrefix(err.Error(), "json: unknown field "):
// 			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
// 			e := Error{Message: fmt.Sprintf("Request body contains unknown field %s", fieldName)}
// 			w.WriteHeader(http.StatusBadRequest)
// 			_ = json.NewEncoder(w).Encode(e)
// 		// An io.EOF error is returned by Decode() if the request body is
// 		// empty.
// 		case errors.Is(err, io.EOF):
// 			e := Error{Message: "Request body must not be empty"}
// 			w.WriteHeader(http.StatusBadRequest)
// 			_ = json.NewEncoder(w).Encode(e)
// 		// Catch the error caused by the request body being too large. Again
// 		// there is an open issue regarding turning this into a sentinel
// 		// error at https://github.com/golang/go/issues/30715.
// 		case err.Error() == "http: request body too large":
// 			e := Error{Message: "Request body must not be larger than 1MB"}
// 			w.WriteHeader(http.StatusBadRequest)
// 			_ = json.NewEncoder(w).Encode(e)

// 		// Otherwise default to logging the error and sending a 500 Internal
// 		// Server Error response.
// 		default:
// 			e := Error{Message: err.Error()}
// 			w.WriteHeader(http.StatusBadRequest)
// 			_ = json.NewEncoder(w).Encode(e)
// 		}
// 		return
// 	}
// 	tc.Uuid = guuid.New().String()
// 	if err := tcpChecksValidator.Validate(&tc); err != nil {
// 		e := Error{Message: "Bad Request - Improper Types Passed"}

// 		w.WriteHeader(http.StatusUnprocessableEntity)
// 		_ = json.NewEncoder(w).Encode(e)
// 		return
// 	}
// 	Checks = append(Checks, tc)
// 	return
// }

// func deleteTcpCheck(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	for idx, tc := range Checks {
// 		if tc.Uuid == params["uuid"] {
// 			Checks = append(Checks[:idx], Checks[idx+1:]...)
// 			w.WriteHeader(http.StatusNoContent)
// 			return
// 		}
// 	}
// 	e := Error{Message: "UUID Not Found"}
// 	w.WriteHeader(http.StatusNotFound)
// 	if err := json.NewEncoder(w).Encode(e); err != nil {
// 		e := Error{Message: "Internal Server Error"}
// 		http.Error(w, e.Message, http.StatusInternalServerError)
// 	}

// }

func HttpCheck(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	w.Header().Set("Content-Type", "text/plain")
	CheckAll()

	fmt.Fprintln(w, viper.GetString("response"))
	fmt.Fprintln(w)
	fmt.Fprintf(w, "Namespace: %-15s\n", viper.GetString("namespace"))
	fmt.Fprintln(w, strings.Repeat("=", 55))
	// fmt.Fprintln(w, "%s\n", )

	// fmt.Fprintf(w, "%s\n")
	hostname, _ := os.Hostname()
	for _, h := range Checks {
		var available string
		switch h.Available {
		case -1:
			available = "UNKNOWN"
		case 1:
			available = "PASS"
		case 0:
			available = "FAIL"
		default:
			available = "UNKNOWN"
		}
		fmt.Fprintf(w, "%-15s -> %-20s : %8s\n",
			hostname,
			fmt.Sprintf("%s (%s)", h.Name, net.JoinHostPort(h.Host, strconv.Itoa(h.Port))),
			available,
		)
	}
}
