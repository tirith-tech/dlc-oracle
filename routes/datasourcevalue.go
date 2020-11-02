package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/tirith-tech/dlc-oracle/datasources"
)

// DataSourceValueResponse for CurrentValue of datasource
type DataSourceValueResponse struct {
	CurrentValue uint64 `json:"currentValue"`
	ValueError   string `json:"valueError,omitempty"`
}

// DataSourceValueHandler takes a datasourceID and returns a json response
func DataSourceValueHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	response := DataSourceValueResponse{}
	datasourceID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response.ValueError = err.Error()
	}
	ds, err := datasources.GetDatasource(datasourceID)

	response.CurrentValue, err = ds.Value()
	if err != nil {
		response.ValueError = err.Error()
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
