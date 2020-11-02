package routes

import (
	"encoding/json"
	"net/http"

	"github.com/tirith-tech/dlc-oracle/datasources"
)

// DataSourceResponse with Name, Description, ID, CurrentValue and ValueError
type DataSourceResponse struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	ID           uint64 `json:"id"`
	CurrentValue uint64 `json:"currentValue"`
	ValueError   string `json:"valueError,omitempty"`
}

// ListDataSourcesHandler handles request for all datasources
func ListDataSourcesHandler(w http.ResponseWriter, r *http.Request) {

	var ds = datasources.GetAllDatasources()

	response := []DataSourceResponse{}
	for _, src := range ds {
		value, err := src.Value()

		jsonSrc := DataSourceResponse{
			Name:         src.Name(),
			Description:  src.Description(),
			ID:           src.Id(),
			CurrentValue: value}

		if err != nil {
			jsonSrc.ValueError = err.Error()
		}
		response = append(response, jsonSrc)
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
