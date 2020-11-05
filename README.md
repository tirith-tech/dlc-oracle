<img src="logo.png">

# Tirith DLC Oracle

This project serves as an oracle to settle Discreet Log Contracts.

This oracle is REST-accessible. It automatically generates R-points on request and keeps published values stored for later retrieval.

The data sources publish every 5 minutes.

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

# REST Endpoints

| Resource                                                     | Description                                                                                                                       |
| :----------------------------------------------------------- | :-------------------------------------------------------------------------------------------------------------------------------- |
| [`localhost:3000/api/pubkey`]                                | Returns the public key of the oracle                                                                                              |
| [`localhost:3000/api/datasources`]                           | Returns an array of data sources the oracle publishes                                                                             |
| [`localhost:3000/api/rpoint/{s}/{t}`]                        | Returns the public one-time-signing key for datasource with ID **s** at the unix timestamp **t**.                                 |
| [`localhost:3000/api/publication/{R}`]                       | Returns the value, signature, timestamp, and name published for data source point **R** (if published). R is hex encoded [33]byte |
| [`localhost:3000/api/publications/tradepair/{base}/{quote}`] | Returns the value, signature, timestamp, and name of all published records for a given pair of **base** and **quote**             |

# Swagger API Documentation

If you don't already have and use the very excellent [Postman API Platform](https://www.postman.com/downloads/) for API development, you can get started playing with the API easily using Swagger. From the root directory of this project run:

```
make serve-swagger
```

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

Forked from [https://github.com/mit-dci/dlc-oracle-go-samples](mit-dci/dlc-oracle-go-samples)
