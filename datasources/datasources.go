package datasources

import (
	"fmt"
)

// Datasource interface
type Datasource interface {
	Id() uint64
	Name() string
	Description() string
	Value() (uint64, error)
	HistoricalValue(uint64) (uint64, error)
	Interval() uint64
}

// Base currencies for Base/Quote pairs: USD/BTC, EUR/BTC, ETH/BTC, etc.
// Order of slice of string cannot be changed without breaking IDs for pairs.
// New bases should be added to the end of slice.
var bases = []string{"USD", "EUR", "JPY", "ETH", "BCH", "LTC", "XRP"}

// GetAllDatasources returns all available datasources
func GetAllDatasources() []Datasource {
	var datasources []Datasource

	for i, val := range bases {
		datasources = append(datasources, &BTC{
			ID:     uint64(i + 1),
			Base:   val,
			Quote:  "BTC",
			prices: make(map[uint64]float64),
		})
	}
	return datasources
}

// GetDatasource returns one datasource by ID
func GetDatasource(id uint64) (Datasource, error) {
	if !HasDatasource(id) {
		return nil, fmt.Errorf("Data source with ID %d not known", id)
	} else {
		return &BTC{
			ID:     id,
			Base:   bases[id-1],
			Quote:  "BTC",
			prices: make(map[uint64]float64),
		}, nil
	}
}

// HasDatasource return boolean response for a given datasource ID
func HasDatasource(id uint64) bool {
	return (id <= uint64(len(bases)))
}
