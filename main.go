package main

import (
	"log"
	"net/http"
	"os"
	"os/user"
	"path"

	"github.com/tirith-tech/dlc-oracle/crypto"
	"github.com/tirith-tech/dlc-oracle/logging"
	"github.com/tirith-tech/dlc-oracle/publisher"
	"github.com/tirith-tech/dlc-oracle/routes"
	"github.com/tirith-tech/dlc-oracle/store"

	"github.com/awnumar/memguard"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	dlcoracle "github.com/mit-dci/dlc-oracle-go"
)

func main() {
	logging.Init(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	logging.Info.Println("Tirith Discreet Log Oracle starting...")

	// Create app folder in user's homedir
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dataDir := path.Join(usr.HomeDir, ".dlcoracle")
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		os.Mkdir(dataDir, 0700)
	}

	// Read or create a keyfile
	keyFilePath := path.Join(dataDir, "privkey.hex")

	key, err := dlcoracle.ReadKeyFile(keyFilePath)
	if err != nil {
		logging.Error.Fatal("Could not open or create keyfile:", err)
		os.Exit(1)
	}
	crypto.StoreKey(key)
	// Safely terminate in case of an interrupt signal
	memguard.CatchInterrupt()

	// Purge the session when we return
	defer memguard.Purge()

	store.Init()
	logging.Info.Println("Connecting to MongoDB...")

	publisher.Init()
	logging.Info.Println("Started publisher...")

	r := mux.NewRouter()
	r.HandleFunc("/api/datasources", routes.ListDataSourcesHandler)
	r.HandleFunc("/api/datasource/{id}/value", routes.DataSourceValueHandler)
	r.HandleFunc("/api/pubkey", routes.PubKeyHandler)
	r.HandleFunc("/api/rpoint/{id}/{timestamp}", routes.RPointHandler)
	r.HandleFunc("/api/publication/{R}", routes.PublicationHandler)
	r.HandleFunc("/api/publications/tradepair/{base}/{quote}", routes.PublicationsHandler)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	// CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	logging.Info.Println("Listening on port 3000")

	logging.Error.Fatal(http.ListenAndServe(":3000", handlers.CORS(originsOk, headersOk, methodsOk)(logging.WebLoggingMiddleware(r))))
}
