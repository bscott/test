package main


import(
	"net/http"


	log "github.com/Sirupsen/logrus"
	"os"
	"regexp"
	"fmt"
)

func RootHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("HOME"))
	log.Debug("RootHandler Invoked")
}

func VersionHandler(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("CHATAPI_VERSION"))
	log.Debug("VersionHandler Invoked")
}

func HealthcheckHandler(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("HEALTHCHECK"))
	log.Debug("HealthCheckHandler Invoked")
}

// Main function, start of the application
func main() {

	log.Info("Starting the ChatAPI Server...")

	router := NewRouter(&Routes{
		Route{
			Name:        "Version",
			Method:      "GET",
			Pattern:     "/version",
			HandlerFunc: VersionHandler,
		},
		Route{
			Name:        "Healthcheck",
			Method:      "GET",
			Pattern:     "/api/status",
			HandlerFunc: HealthcheckHandler,
		},
		Route{
			Name:	"Root",
			Method:	"GET",
			Pattern: "/",
			HandlerFunc: RootHandler,
		}})

	// Setup some port variables
	port := os.Getenv("CHATAPI_HTTP_PORT")
	port_regex_str := "^[0-9]+$"
	port_regex, err := regexp.Compile(port_regex_str)

	if err != nil {
		log.Warn(fmt.Sprintf("Port regex \"%s\"could not be compiled, details: %s", port_regex_str, err))
		port = "8080"
		return
	} else {
		if !port_regex.MatchString(port) {
			log.Warn(fmt.Sprintf("CHATAPI_HTTP_PORT value \"%s\" is either invalid or it was not provided.  Defaulting to 8080", port))
			port = "8080"
		}
	}

	host := os.Getenv("CHATAPI_HTTP_HOST")

	if host == "" {
		host = "0.0.0.0"
	}

	// Start the API Server
	log.Info(fmt.Sprintf("Serving ChatOps API on %s:%s", host, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router))


}

