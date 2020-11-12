package datasources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type pair struct {
	ID       uint64 `json:"id"`
	Provider string `json:"provider"`
	Base     string `json:"base"`
	Interval uint64 `json:"interval"`
	RoundTo  uint64 `json:"roundTo"`
}

// GetAllDatasources returns all available datasources
func GetAllDatasources() []Datasource {
	var datasources []Datasource

	pairs := getPairs()

	for _, pair := range pairs {
		if pair.Provider == "Crypto Compare" {
			datasources = append(datasources, &CryptoCompare{
				ID:       uint64(pair.ID),
				Provider: pair.Provider,
				Base:     pair.Base,
				Quote:    "BTC",
				interval: uint64(pair.Interval),
				roundTo:  uint64(pair.RoundTo),
				prices:   make(map[uint64]float64),
			})
		}

	}
	return datasources
}

// GetDatasource returns one datasource by ID
func GetDatasource(id uint64) (Datasource, error) {
	if !HasDatasource(id) {
		return nil, fmt.Errorf("Data source with ID %d not known", id)
	} else {
		pairs := getPairs()

		for _, pair := range pairs {
			if pair.ID == id {
				if pair.Provider == "Crypto Compare" {
					return &CryptoCompare{
						ID:       id,
						Provider: pair.Provider,
						Base:     pair.Base,
						Quote:    "BTC",
						interval: uint64(pair.Interval),
						roundTo:  uint64(pair.RoundTo),
						prices:   make(map[uint64]float64),
					}, nil
				}
			}
		}
	}
	return nil, nil
}

// HasDatasource return boolean response for a given datasource ID
func HasDatasource(id uint64) bool {
	pairs := getPairs()
	return (id <= uint64(len(pairs)))
}

func getPairs() []pair {
	// Get pairs from JSON
	// Pair IDs must not be changed once node(s) launched
	file, err := ioutil.ReadFile("datasources/btcpairs.json")
	if err != nil {
		fmt.Println(err)
	}

	var pairs []pair

	err = json.Unmarshal([]byte(file), &pairs)
	if err != nil {
		fmt.Println(err)
	}

	return pairs
}
