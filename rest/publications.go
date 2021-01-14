package rest

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tirith-tech/dlc-oracle/logging"
	"github.com/tirith-tech/dlc-oracle/store"

	"github.com/gorilla/mux"
)

// PublicationsResponse returns a list of PublicationResponses
type PublicationsResponse struct {
	Publications []PublicationResponse `json:"publications"`
}

// PublicationsHandler takes a base and quote and returns all related publications
func PublicationsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PublicationsHandler")
	vars := mux.Vars(r)
	name := fmt.Sprintf("%v/%v", vars["base"], vars["quote"])

	publications, err := store.GetAllPublicationsByName(name)
	if err != nil {
		logging.Error.Println("SubscribeHandler - Error getting all publications: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responses := []PublicationResponse{}

	for _, p := range publications {
		publication := PublicationResponse{Value: p.Value, Signature: hex.EncodeToString(p.Signature[:]), Timestamp: p.Timestamp, Name: p.Name}
		responses = append(responses, publication)
	}
	js, err := json.Marshal(responses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
