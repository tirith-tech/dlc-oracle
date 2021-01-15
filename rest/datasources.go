package rest

import (
	"encoding/json"
	"net/http"

	"github.com/tirith-tech/dlc-oracle/datasources"
	"github.com/tirith-tech/dlc-oracle/logging"
	"github.com/tirith-tech/dlc-oracle/store"
)

// SeriesResponse with Name, Data
type SeriesResponse struct {
	Name string     `json:"name"`
	Data [][]uint64 `json:"data"`
}

// DataSourceResponse with Name, Description, ID, CurrentValue, Series and ValueError
type DataSourceResponse struct {
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	ID           uint64           `json:"id"`
	CurrentValue uint64           `json:"currentValue"`
	Series       []SeriesResponse `json:"series"`
	ValueError   string           `json:"valueError,omitempty"`
}

// ListDataSourcesHandler handles request for all datasources
func ListDataSourcesHandler(w http.ResponseWriter, r *http.Request) {

	var ds = datasources.GetAllDatasources()

	response := []DataSourceResponse{}
	for _, src := range ds {

		publications, err := store.GetAllPublicationsByName(src.Name())
		if err != nil {
			logging.Error.Println("SubscribeHandler - Error getting all publications: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		series := make([][]uint64, 0)

		for _, p := range publications {
			if p.Timestamp%86400 == 0 { // Limit to midnight only
				tick := []uint64{p.Timestamp * 1000, p.Value}
				series = append(series, tick)
			}
		}

		// Only send 30 most recent values
		trimmedSeries := make([][]uint64, 0)
		if len(series) > 30 {
			trimmedSeries = series[len(series)-30:]
		} else {
			trimmedSeries = series
		}

		value, err := src.Value()

		seriesResponse := SeriesResponse{Name: src.Name(), Data: trimmedSeries}
		seriesSlice := []SeriesResponse{}
		seriesSlice = append(seriesSlice, seriesResponse)

		jsonSrc := DataSourceResponse{
			Name:         src.Name(),
			Description:  src.Description(),
			ID:           src.Id(),
			CurrentValue: value,
			Series:       seriesSlice,
		}

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
