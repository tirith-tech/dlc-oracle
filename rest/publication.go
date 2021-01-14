package rest

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/tirith-tech/dlc-oracle/logging"
	"github.com/tirith-tech/dlc-oracle/store"

	"github.com/gorilla/mux"
)

// PublicationResponse for marshalling JSON response with Value, Signature, Timestamp, and Name
type PublicationResponse struct {
	Value     uint64 `json:"value"`
	Signature string `json:"signature"`
	Timestamp uint64 `json:"timestamp"`
	Name      string `json:"name"`
}

// PublicationHandler takes an "R" value and returns a publication
func PublicationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	passedR, err := hex.DecodeString(vars["R"])
	if err != nil {
		logging.Error.Println("SubscribeHandler - Error parsing R: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var R [33]byte
	copy(R[:], passedR[:])

	value, signature, timestamp, name, err := store.GetPublication(R)
	if err != nil {
		logging.Error.Println("SubscribeHandler - Error getting publication: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := PublicationResponse{Value: value, Signature: hex.EncodeToString(signature[:]), Timestamp: timestamp, Name: name}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

// PublicationByNameAndTimestampHandler takes a name and timestamp value and returns a publication
func PublicationByNameAndTimestampHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := fmt.Sprintf("%v/%v", vars["base"], vars["quote"])
	timestamp, err := strconv.ParseUint(vars["timestamp"], 10, 64)
	if err != nil {
		logging.Error.Println("PublicationByTimestampHandler - Invalid Datasource: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	value, signature, timestamp, name, err := store.GetPublicationByNameAndTimestamp(name, timestamp)
	if err != nil {
		logging.Error.Println("SubscribeHandler - Error getting publication: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := PublicationResponse{Value: value, Signature: hex.EncodeToString(signature[:]), Timestamp: timestamp, Name: name}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
