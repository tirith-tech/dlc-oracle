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

	app := &cli.App{
		Name:  "Tirith DLC Oracle",
		Usage: "The Beacons of Minas Tirith! The Beacons are lit!",
		Commands: []*cli.Command{
			{
				Name:  "rest",
				Usage: "Run Oracle as RESTful API",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "password",
						Aliases: []string{"p"},
						Usage:   "Pass password to process",
					},
				},
				Action: func(c *cli.Context) error {
					log.Print(c.String("password"))
					loadKey([]byte(c.String("password")))
					services()
					// REST API
					rest.Init()
					return nil
				},
			},
			{
				Name:  "rpc",
				Usage: "Run Oracle as gRPC API",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "password",
						Aliases: []string{"p"},
						Usage:   "Pass password to process",
					},
				},
				Action: func(c *cli.Context) error {
					log.Print(c.String("password"))
					loadKey([]byte(c.String("password")))
					services()
					// gRPC Server
					rpc.Init()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logging.Error.Fatal(err)
	}
}

func loadKey(password []byte) {
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

	key, err := dlcoracle.LoadKeyFromFileArg(keyFilePath, password)

	if err != nil {
		logging.Error.Fatal("Could not open or create keyfile:", err)
		os.Exit(1)
	}
	crypto.StoreKey(key)
	// Safely terminate in case of an interrupt signal
	memguard.CatchInterrupt()

	// Purge the session when we return
	defer memguard.Purge()
}

func services() {
	store.Init()
	logging.Info.Println("Connecting to MongoDB...")

	publisher.Init()
	logging.Info.Println("Started publisher...")
}
