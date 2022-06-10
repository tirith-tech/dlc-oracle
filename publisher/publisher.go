package publisher

import (
	"time"

	"github.com/tirith-tech/dlc-oracle/crypto"
	"github.com/tirith-tech/dlc-oracle/datasources"
	"github.com/tirith-tech/dlc-oracle/logging"
	"github.com/tirith-tech/dlc-oracle/store"

	dlcoracle "github.com/mit-dci/dlc-oracle-go"
)

var lastPublished = uint64(0)
var dataSources []datasources.Datasource

// Init finds lastPublished timestamp or runs from current time if database is empty
// It then queries for all datasources
// Finally, it launches a go routine of the Process function incremented every second
func Init() {
	var err error
	lastPublished, err = store.GetLastPublicationTimestamp()
	if err != nil {
		lastPublished = uint64(time.Now().Unix())
		logging.Error.Println(err)
		logging.Error.Println("No publications found. Starting from current timestamp:", lastPublished)
	} else {
		logging.Info.Println("Resuming from last publication timestamp:", lastPublished)
	}

	dataSources = datasources.GetAllDatasources()

	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			Process()
		}
	}()
}

// Process iterates over all datasources for each second from the last published record up to present time
// If data is missing, historical prices are queried, otherwise current prices are requested
// Oracle private key (a), signing key (k) and rPoint (public key) are retreived
// Signature created for publication
// rPoint, value, signature, time and name are published to Publications collection
func Process() error {
	timeNow := uint64(time.Now().Unix())
	var valueToPublish uint64 = 0
	for time := lastPublished + 1; time <= timeNow; time++ {
		for _, ds := range dataSources {

			if time%ds.Interval() == 0 && datasources.HasDatasource(ds.Id()) {

				if lastPublished+1 == time {
					logging.Info.Printf("[%v] %v PUBLISH CURRENT [ts: %d]\n", ds.Id(), ds.Name(), time)

					value, err := ds.Value()
					if err != nil {
						logging.Error.Printf("Could not retrieve value for data source %d: %s", ds.Id(), err.Error())
						continue
					}
					valueToPublish = value
				} else {
					logging.Info.Printf("[%v] %v PUBLISH HISTORICAL [ts: %d]\n", ds.Id(), ds.Name(), time)

					value, err := ds.HistoricalValue(time)
					if err != nil {
						logging.Error.Printf("Could not retrieve value for data source %d: %s", ds.Id(), err.Error())
						continue
					}
					valueToPublish = value
				}

				var a [32]byte
				copy(a[:], crypto.RetrieveKey()[:])

				k, err := store.GetK(ds.Id(), time)
				if err != nil {
					logging.Error.Printf("Could not get signing key for data source %d and timestamp %d : %s", ds.Id(), time, err.Error())
					continue
				}

				R, err := store.GetRPoint(ds.Id(), time)
				if err != nil {
					logging.Error.Printf("Could not get pubkey for data source %d and timestamp %d : %s", ds.Id(), time, err.Error())
					continue
				}

				publishedAlready, err := store.IsPublished(R)
				if err != nil {
					logging.Error.Printf("Error determining if this is already published: %s", err.Error())
					continue
				}

				if publishedAlready {
					logging.Info.Printf("Already published for data source %d and timestamp %d", ds.Id(), time)
					continue
				}

				message := dlcoracle.GenerateNumericMessage(valueToPublish)

				signature, err := dlcoracle.ComputeSignature(a, k, message)
				if err != nil {
					logging.Error.Printf("Could not sign the message: %s", err.Error())
					continue
				}

				store.Publish(R, valueToPublish, signature, time, ds.Name())
			}
		}
	}

	lastPublished = timeNow
	return nil
}
