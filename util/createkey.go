package main

import (
	"log"
	"os"
	"os/user"
	"path"

	dlcoracle "github.com/mit-dci/dlc-oracle-go"
	"github.com/tirith-tech/dlc-oracle/logging"
)

func main() {
	// Create app folder in user's homedir
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dataDir := path.Join(usr.HomeDir, ".dlcoracle")
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		os.Mkdir(dataDir, 0700)
	}

	// Read or create a keyfile
	keyFilePath := path.Join(dataDir, "privkey.hex")

	key := new([32]byte)
	pw := []byte{}
	err = dlcoracle.SaveKeyToFileArg(keyFilePath, key, pw)
	if err != nil {
		logging.Error.Fatal(err)
	}
}
