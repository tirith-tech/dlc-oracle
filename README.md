# DLC Oracle

This project serves as an oracle to settle Bitcoin Discreet Log Contracts.

This oracle is REST-accessible. It automatically generates R-points on request and keeps published values stored for later retrieval.

The data sources publish every 24 hours at midnight GMT.
<br><br>

# Installation

First, install [Go](https://golang.org/doc/install).

Second, you will need to install [MongoDB](https://docs.mongodb.com/manual/installation/). Community Edition is sufficient. For an easier time working with the database, you may also want to install [Robo3T](https://robomongo.org/download).
<br><br>
Next, clone the repo:

```
go get github.com/tirith-tech/dlc-oracle
```

Then, you will need to install the dependencies from the root directory:

```
go mod download
```

To run as RESTful API:

```
go run main.go rest
```

to run as gRPC API:

```
go run main.go rpc
```

<br><br>

# REST Endpoints

| Resource                                                        | Description                                                                                                                       |
| :-------------------------------------------------------------- | :-------------------------------------------------------------------------------------------------------------------------------- |
| [`localhost:3000/api/pubkey`]                                   | Returns the public key of the oracle                                                                                              |
| [`localhost:3000/api/datasources`]                              | Returns an array of data sources the oracle publishes                                                                             |
| [`localhost:3000/api/rpoint/{s}/{t}`]                           | Returns the public one-time-signing key for datasource with ID **s** at the unix timestamp **t**.                                 |
| [`localhost:3000/api/pub/rpoint/{R}`]                           | Returns the value, signature, timestamp, and name published for data source point **R** (if published). R is hex encoded [33]byte |
| [`localhost:3000/api/pub/tradepair/{base}/{quote}/{timestamp}`] | Returns (value, signature, timestamp, name) of published record for a given pair of **base**/**quote** and **timestamp**          |
| [`localhost:3000/api/pubs/tradepair/{base}/{quote}`]            | Returns the value, signature, timestamp, and name of all published records for a given pair of **base** and **quote**             |

<br><br>

# Swagger API Documentation

If you don't already have and use the very excellent [Insomnia API Platform](https://insomnia.rest/) for API development, you can get started playing with the API easily using Swagger. From the root directory of this project run:

```
make serve-swagger
```

<br><br>

# Docker

You will need [Docker](https://www.docker.com/) installed on your machine.

1. First, set a docker secret as your Oracle Private Key encryption password:

```
 echo "YOUR_PASSWORD" | docker secret create oracle_pw -
```

2. Next, run:

```
docker-compose up
```

You will find all of the endpoints at [`localhost:3000`] the same as compiling and running from Go locally.
<br><br>

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

Forked from [https://github.com/mit-dci/dlc-oracle-go-samples](https://github.com/mit-dci/dlc-oracle-go-samples)
