package main

import (
	"log"
	"os"
	"os/user"
	"path"

	"github.com/tirith-tech/dlc-oracle/crypto"
	"github.com/tirith-tech/dlc-oracle/logging"
	"github.com/tirith-tech/dlc-oracle/publisher"
	"github.com/tirith-tech/dlc-oracle/rest"
	"github.com/tirith-tech/dlc-oracle/rpc"
	"github.com/tirith-tech/dlc-oracle/store"
	"github.com/urfave/cli/v2"

	"github.com/awnumar/memguard"
	dlcoracle "github.com/mit-dci/dlc-oracle-go"
)

func main() {
	logging.Init(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	logging.Info.Println("Tirith Discreet Log Oracle starting...")

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

	key, err := dlcoracle.ReadKeyFile(keyFilePath)

	if err != nil {
		logging.Error.Fatal("Could not open or create keyfile:", err)
		os.Exit(1)
	}
	crypto.StoreKey(key)
	// Safely terminate in case of an interrupt signal
	memguard.CatchInterrupt()

	// Purge the session when we return
	defer memguard.Purge()

	app := &cli.App{
		Name:  "Tirith DLC Oracle",
		Usage: "The Beacons of Minas Tirith! The Beacons are lit!",
		Commands: []*cli.Command{
			{
				Name:  "rest",
				Usage: "Run Oracle as RESTful API",
				Action: func(c *cli.Context) error {
					services()
					// REST API
					rest.Init()
					return nil
				},
			},
			{
				Name:  "rpc",
				Usage: "Run Oracle as gRPC API",
				Action: func(c *cli.Context) error {
					services()
					// gRPC Server
					rpc.Init()
					return nil
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		logging.Error.Fatal(err)
	}
}

func services() {
	store.Init()
	logging.Info.Println("Connecting to MongoDB...")

	publisher.Init()
	logging.Info.Println("Started publisher...")
}
