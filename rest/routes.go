package rest

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tirith-tech/dlc-oracle/logging"
)

// Init REST API
func Init() {
	// Rest API
	r := mux.NewRouter()
	r.HandleFunc("/api/datasources", ListDataSourcesHandler)
	r.HandleFunc("/api/pubkey", PubKeyHandler)
	r.HandleFunc("/api/rpoint/{id}/{timestamp}", RPointHandler)
	r.HandleFunc("/api/pub/rpoint/{R}", PublicationHandler)
	r.HandleFunc("/api/pub/tradepair/{base}/{quote}/{timestamp}", PublicationByNameAndTimestampHandler)
	r.HandleFunc("/api/pubs/tradepair/{base}/{quote}", PublicationsHandler)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	// CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	logging.Info.Println("Listening on port 80")

	logging.Error.Fatal(http.ListenAndServe(":80", handlers.CORS(originsOk, headersOk, methodsOk)(logging.WebLoggingMiddleware(r))))
}
