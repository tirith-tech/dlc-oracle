// Package docs Tirith DLC Oracle.
//
// Documentation of Tirith DLC Oracle Rest API.
//
//     Schemes: http, https
//		 Host: localhost:3000
//     BasePath: /api
//     Version: 0.0.1
//		 License: MIT http://opensource.org/licenses/MIT
//     Contact Robert William Allen<robertwilliamallen@gmail.com> https//robertwilliamallen.github.io
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - basic
//
//    SecurityDefinitions:
//    basic:
//      type: basic
//
// swagger:meta
package docs

import "github.com/tirith-tech/dlc-oracle/routes"

// swagger:route GET /pubkey PubKey idPubkeyEndpoint
// Pubkey returns the public key of the DLC Oracle.
// responses:
//   200: PubKeyResponse

// Returns the Oracle public key.
// swagger:response PubKeyResponse
type pubKeyResponseWrapper struct {
	// in:body
	Body routes.PubKeyResponse
}

// swagger:route GET /datasources Datasources idDatasourcesEndpoint
// DataSources returns a list of available datasources.
// responses:
//   200: DataSourceResponse

// Returns a JSON list of datasources.
// swagger:response DataSourceResponse
type datasourceResponseWrapper struct {
	// in:body
	Body routes.DataSourceResponse
}

// swagger:route GET /datasource/{id}/value Datasources idDatasourceValueEndpoint
// DataSourceValue returns the value for a datasource by ID.
// responses:
//   200: DataSourceValueResponse

// Returns the value for a datasource by ID.
// swagger:response DataSourceValueResponse
type datasourceValueResponseWrapper struct {
	// in:body
	Body routes.DataSourceValueResponse
}

// swagger:parameters idDatasourceValueEndpoint
type datasourceValueParam struct {
	// required:true
	// in:path
	ID int `json:"id"`
}

// swagger:route GET /rpoint/{id}/{timestamp} rPoint idRPointEndpoint
// RPointResponse returns the Oracle R point for a given ID and unix timestamp.
// responses:
//   200: RPointResponse
//   404: description: No datasource available for ID
//	 406: description: Timestamp must be divisible by datasource publishing interval
//	 500: description: Server Error

// Returns an rPoint for a given ID and timestamp.
// swagger:response RPointResponse
type rPointResponseWrapper struct {
	// in:body
	Body routes.RPointResponse
}

// swagger:parameters idRPointEndpoint
type rPointParams struct {
	// required:true
	// in:path
	ID int `json:"id"`
	// required:true
	// in:path
	Timestamp int `json:"timestamp"`
}

// swagger:route GET /publication/{R} Publications idPublicationEndpoint
// PublicationResponse returns the value, signature, timestamp, and name for corresponding rPoint.
// responses:
//   200: PublicationResponse
//	 500: description: Server Error

// Returns value, signature, timestamp, and name for corresponding rPoint.
// swagger:response PublicationResponse
type publicationResponseWrapper struct {
	// in:body
	Body routes.PublicationResponse
}

// swagger:parameters idPublicationEndpoint
type publicationResponseParam struct {
	// required:true
	// in:path
	R string `json:"R"`
}

// swagger:route GET /publications/tradepair/{base}/{quote} Publications idPublicationsEndpoint
// PublicationsResponse returns a list of publications with value, signature, timestamp, and name for corresponding rPoint.
// responses:
//   200: PublicationsResponse
//	 500: description: Server Error

// Returns value, signature, timestamp, and name for corresponding rPoint.
// swagger:response PublicationsResponse
type publicationsResponseWrapper struct {
	// in:body
	Body routes.PublicationsResponse
}

// swagger:parameters idPublicationsEndpoint
type publicationsResponseParams struct {
	// required:true
	// in:path
	Base string `json:"base"`
	// required:true
	// in:path
	Quote string `json:"quote"`
}
