package store

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/binary"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	"github.com/tirith-tech/dlc-oracle/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collectionKeys *mongo.Collection
var collectionPubs *mongo.Collection
var ctx = context.TODO()

// Init initializes connection to MongoDB and makes/retreives two collections
// "Keys" which contains publication index keys (from timestamp and publicationID) and private keys
// "Publications" which contains Value, Signature, Timestamp, and Name of publication
func Init() {
	// Set MongoDB options and create client
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logging.Error.Print(err)
	}

	// Test DB connection
	err = client.Ping(ctx, nil)
	if err != nil {
		logging.Error.Print(err)
	}

	// Get collections
	collectionKeys = client.Database("Tirith").Collection("Keys")
	collectionPubs = client.Database("Tirith").Collection("Publications")
}

// Publication with Value, Signature, Timestamp, and Name of publication
type Publication struct {
	Value     uint64
	Signature []byte
	Timestamp uint64
	Name      string
}

// Key with private key
type Key struct {
	PrivKey []byte
}

// GetRPoint takes datasourceID and timestamp and then creates and returns rPoint
func GetRPoint(datasourceID, timestamp uint64) ([33]byte, error) {
	var pubKey [33]byte

	privKey, err := GetK(datasourceID, timestamp)
	if err != nil {
		logging.Error.Print(err)
		return pubKey, err
	}

	_, pk := btcec.PrivKeyFromBytes(btcec.S256(), privKey[:])

	copy(pubKey[:], pk.SerializeCompressed())
	return pubKey, nil
}

// GetK takes datasourceID and timestamp
// datasourceID and timestamp are used to create the storage key for the Keys collection
// Creates privKey if not found and then inserts new record in Keys collection
// Finally, returns privKey
func GetK(datasourceID, timestamp uint64) ([32]byte, error) {
	var privKey [32]byte

	key := makeStorageKey(datasourceID, timestamp)

	var result Key
	filter := bson.M{"key": key}
	err := collectionKeys.FindOne(ctx, filter).Decode(&result)

	if err != nil { // No record found for key
		_, err := rand.Read(privKey[:])
		if err != nil {
			return [32]byte{},

				err
		}

		_, err = collectionKeys.InsertOne(ctx, bson.M{"key": key, "privKey": privKey[:]})
		if err != nil {
			return [32]byte{},

				err
		}
	} else { // privKey found
		copy(privKey[:], result.PrivKey[:])
	}
	return privKey, nil
}

// Publish takes in rPoint, value, signature, time and name and returns nil if successfully published
func Publish(rPoint [33]byte, value uint64, signature [32]byte, time uint64, name string) error {
	var result Publication
	filter := bson.M{"rPoint": rPoint[:]}
	err := collectionPubs.FindOne(ctx, filter).Decode(&result)
	if err == nil { // Should stop if a record exists
		return fmt.Errorf("There is already a value published for this rpoint")
	}

	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, value)
	buf.Write(signature[:])

	_, err = collectionPubs.InsertOne(ctx, bson.M{"rPoint": rPoint[:], "value": value, "signature": buf.Bytes(), "timestamp": time, "name": name})
	if err != nil {
		fmt.Println(err)
		return err
	}
	logging.Info.Printf("rPoint published %v\n", rPoint)

	return nil
}

// IsPublished takes in rPoint and returns true, nil if published or false, nil if unpublished
func IsPublished(rPoint [33]byte) (bool, error) {
	var result Publication
	filter := bson.M{"rPoint": rPoint[:]}
	err := collectionPubs.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return false, nil
	}

	return true, nil
}

// GetPublication takes rPoint and returns value, signature, timestamp and name of publication if found
func GetPublication(rPoint [33]byte) (uint64, []byte, uint64, string, error) {
	var result Publication
	filter := bson.M{"rPoint": rPoint}
	err := collectionPubs.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return 0, []byte{}, 0, "", fmt.Errorf("Publication not found")
	}

	logging.Info.Printf("Publication found at rPoint %v with Signature %v and Value %v\n", rPoint, result.Signature, result.Value)

	return result.Value, result.Signature, result.Timestamp, result.Name, nil
}

// GetLastPublicationTimestamp finds most recent timestamp for all publications
// This function is used to gather historical data in the event an Oracle is down and rebooted and datapoints are missing
func GetLastPublicationTimestamp() (uint64, error) {
	filter := bson.D{{}}
	publications, err := filterPublications(filter)
	if err != nil {
		return 0, err
	}
	return publications[len(publications)-1].Timestamp, nil
}

// GetAllPublicationsByName takes in a name string and returns all matching publications
func GetAllPublicationsByName(name string) ([]*Publication, error) {
	filter := bson.M{"name": name}
	publications, err := filterPublications(filter)
	if err != nil {
		return nil, err
	}
	return publications, nil
}

// filterPublications is a non-exported helper-function to filter publications by some search value
func filterPublications(filter interface{}) ([]*Publication, error) {
	// A slice of Publication for storing the decoded documents
	var publications []*Publication

	cur, err := collectionPubs.Find(ctx, filter)
	if err != nil {
		return publications, err
	}

	for cur.Next(ctx) {
		var p Publication
		err := cur.Decode(&p)
		if err != nil {
			return publications, err
		}

		publications = append(publications, &p)
	}

	if err := cur.Err(); err != nil {
		return publications, err
	}

	// Once exhausted, close the cursor
	cur.Close(ctx)

	if len(publications) == 0 {
		return publications, mongo.ErrNoDocuments
	}

	return publications, nil
}

// makeStorageKey is a helper-function which creates a storage key for the Keys collection
func makeStorageKey(datasourceID uint64, timestamp uint64) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, timestamp)
	binary.Write(&buf, binary.BigEndian, datasourceID)
	return buf.Bytes()
}
