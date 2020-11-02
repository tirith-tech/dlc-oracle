package routes

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/tirith-tech/dlc-oracle/crypto"
	"github.com/tirith-tech/dlc-oracle/datasources"
	"github.com/tirith-tech/dlc-oracle/logging"
	"github.com/tirith-tech/dlc-oracle/store"

	"github.com/gorilla/mux"
)

// RPointResponse R string
type RPointResponse struct {
	R string
}

// RPointHandler takes a datasourceID and timestamp and then creates and returns an R value
func RPointHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	datasourceID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		logging.Error.Println("RPointPubKeyHandler - Invalid Datasource: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !datasources.HasDatasource(datasourceID) {
		logging.Error.Println("RPointPubKeyHandler - Invalid Datasource: ", datasourceID)
		http.Error(w, fmt.Sprintf("Invalid datasource %d", datasourceID), http.StatusInternalServerError)
		return
	}

	ds, err := datasources.GetDatasource((datasourceID))
	if err != nil {
		logging.Error.Println("RPointPubKeyHandler - Unable to Retreive Datasource: ", datasourceID)
		http.Error(w, fmt.Sprintf("Invalid datasource %d", datasourceID), http.StatusInternalServerError)
	}

	timestamp, err := strconv.ParseUint(vars["timestamp"], 10, 64)
	if err != nil {
		logging.Error.Println("RPointPubKeyHandler - Invalid Timestamp: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if timestamp%ds.Interval() != 0 {
		message := fmt.Sprintf("Invalid Timestamp: Must be divisible evenly by datasource interval [%v seconds].", ds.Interval())
		logging.Error.Println("RPointPubKeyHandler - Invalid Timestamp (Not an acceptable interval)")
		http.Error(w, message, http.StatusNotAcceptable)
		return
	}

	rPoint, err := store.GetRPoint(datasourceID, timestamp)
	if err != nil {
		logging.Error.Println("RPointPubKeyHandler", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := RPointResponse{
		R: hex.EncodeToString(rPoint[:]),
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// PubKeyResponse A string
type PubKeyResponse struct {
	A string
}

// PubKeyHandler returns DLC Oracle public key
func PubKeyHandler(w http.ResponseWriter, r *http.Request) {
	A, err := crypto.GetPubKey()
	if err != nil {
		logging.Error.Println("PubKeyHandler", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := PubKeyResponse{
		A: hex.EncodeToString(A[:]),
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
