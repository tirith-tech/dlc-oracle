package datasources

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/tirith-tech/dlc-oracle/logging"
)

// MinAPICryptoCompareBTCResponse for current price from Crypto Compare
type MinAPICryptoCompareBTCResponse struct {
	Value float64 `json:"BTC"`
}

// MinAPICryptoCompareHistoricalBTCResponse nested struct for historical prices from Crypto Compare
type MinAPICryptoCompareHistoricalBTCResponse struct {
	Response   string `json:"Response"`
	HasWarning bool   `json:"HasWarning"`
	Type       int    `json:"Type"`
	Data       struct {
		Aggregated bool   `json:"Aggregated"`
		TimeFrom   int    `json:"TimeFrom"`
		TimeTo     int    `json:"TimeTo"`
		Data       []data `json:"Data"`
	} `json:"Data"`
}

type data struct {
	Time             uint64  `json:"time"`
	High             float64 `json:"high"`
	Low              float64 `json:"low"`
	Open             float64 `json:"open"`
	VolumeFrom       float64 `json:"volumefrom`
	VolumeTo         float64 `json:"volumeto"`
	Close            float64 `json:"close"`
	ConversionType   string  `json:"conversionType"`
	ConversionSymbol string  `json:"conversionSymbol"`
}

// CryptoCompare structure with historical prices map (used as cache for historical calls)
type CryptoCompare struct {
	ID       uint64
	Provider string
	Base     string
	Quote    string
	interval uint64
	roundTo  uint64
	prices   map[uint64]float64
}

// Id receiver function returns ID
func (ds *CryptoCompare) Id() uint64 {
	return ds.ID
}

// Name receiver function returns Base name of pair traded against BTC
func (ds *CryptoCompare) Name() string {
	return fmt.Sprintf("%v/%v", ds.Base, ds.Quote)
}

// Description returns string description of pair
func (ds *CryptoCompare) Description() string {
	return fmt.Sprintf("Publishes the value of %v denominated in 1/100000000th units of %v (satoshis) in multitudes of %v", ds.Base, ds.Quote, ds.roundTo)
}

// Interval returns the time interval between published data in seconds
func (ds *CryptoCompare) Interval() uint64 {
	return ds.interval
}

// Value returns the current value of asset priced in satoshis
func (ds *CryptoCompare) Value() (uint64, error) {
	url := fmt.Sprintf("https://min-api.cryptocompare.com/data/price?fsym=%v&tsyms=%v", ds.Base, ds.Quote)
	resp, err := ds.getData(url)
	if err != nil {
		logging.Error.Println(err)
		return 0, err
	}

	defer resp.Body.Close()

	var record MinAPICryptoCompareBTCResponse

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		message := fmt.Sprintf("[%v] %v.Value - Json decode failed:", ds.ID, ds.Name())
		logging.Error.Println(message, err)
		return 0, err
	}

	satoshiValue := ds.satoshisRounded(record.Value)
	logging.Info.Printf("[%v] %v CURRENT [sats: %v]\n", ds.ID, ds.Name(), satoshiValue)
	return satoshiValue, nil
}

// HistoricalValue returns the historical value of an asset priced in satoshis at a given timestamp
func (ds *CryptoCompare) HistoricalValue(timestamp uint64) (uint64, error) {
	// Check to see if timestamp key exists in prices map, and if so, return value
	if ds.prices[timestamp] != 0 {
		satoshiValue := ds.satoshisRounded(ds.prices[timestamp])
		logging.Info.Printf("[%v] %v HISTORICAL [sats: %v]\n", ds.ID, ds.Name(), satoshiValue)
		return satoshiValue, nil // Price available at this timestamp
	}

	// Timestamp missing from prices map, populate map from CryptoCompare
	recordRange := uint64(2000 * 60) // 2000 minute candles
	fromTimestamp := timestamp + recordRange
	logging.Info.Printf("[%v] %v - %v - FETCHING HISTORICAL [ts: %v]\n", ds.ID, ds.Name(), ds.Provider, fromTimestamp)
	url := fmt.Sprintf("https://min-api.cryptocompare.com/data/v2/histominute?fsym=%v&tsym=%v&limit=2000&toTs=%v", ds.Base, ds.Quote, fromTimestamp)
	resp, err := ds.getData(url)
	if err != nil {
		logging.Error.Println(err)
		return 0, err
	}
	defer resp.Body.Close()

	var record MinAPICryptoCompareHistoricalBTCResponse

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		message := fmt.Sprintf("[%v] %v.Value - Json decode failed:", ds.ID, ds.Name())
		logging.Error.Println(message, err)
		return 0, err
	}

	// Check to be sure that oldest historical record is not more recent than requested timestamp value
	oldest := record.Data.Data[0].Time
	if oldest > timestamp {
		return 0, fmt.Errorf("[%v] %v - %v - NO DATA - OLDEST [ts: %v] > NEEDED [ts: %v]", ds.ID, ds.Name(), ds.Provider, oldest, timestamp)
	}

	// Add timestamps/closes to map
	for _, v := range record.Data.Data {
		ds.prices[v.Time] = v.Close
	}

	if ds.prices[timestamp] != 0 {
		satoshiValue := ds.satoshisRounded(ds.prices[timestamp])
		logging.Info.Printf("[%v] %v HISTORICAL [sats: %v]\n", ds.ID, ds.Name(), satoshiValue)
		return satoshiValue, nil // Price available at this timestamp
	} else {
		return 0, nil
	}
}

// Helper function to get data from a url
func (ds *CryptoCompare) getData(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		message := fmt.Sprintf("[%v] %v.Value - NewRequest:", ds.ID, ds.Name())
		logging.Error.Println(message, err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		message := fmt.Sprintf("[%v] %v.Value - Do: ", ds.ID, ds.Name())
		logging.Error.Println(message, err)
		return resp, err
	}

	return resp, nil
}

// Helper function to convert BTC to satoshis by roundTo
func (ds *CryptoCompare) satoshisRounded(price float64) uint64 {
	place := float64(100000000 / ds.roundTo)
	return uint64(math.Floor((price*place)+0.5)) * ds.roundTo
}
