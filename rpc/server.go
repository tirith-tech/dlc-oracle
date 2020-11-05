package rpc

import (
	"context"
	"encoding/hex"
	"fmt"
	"net"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/tirith-tech/dlc-oracle/crypto"
	"github.com/tirith-tech/dlc-oracle/datasources"
	"github.com/tirith-tech/dlc-oracle/logging"
	"github.com/tirith-tech/dlc-oracle/rpc/protobuf"
	"github.com/tirith-tech/dlc-oracle/store"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) PubKey(context.Context, *emptypb.Empty) (*protobuf.PubKeyResponse, error) {
	logging.Info.Println("PubKey was invoked")
	A, err := crypto.GetPubKey()
	if err != nil {
		logging.Error.Println("gRPC - PubKey: ", err)
		return nil, err
	}

	key := hex.EncodeToString(A[:])

	res := &protobuf.PubKeyResponse{
		Pubkey: key,
	}
	return res, nil
}

func (*server) DataSources(empty *emptypb.Empty, stream protobuf.OracleService_DataSourcesServer) error {
	logging.Info.Println("DataSources was invoked")
	datasources := datasources.GetAllDatasources()

	for _, ds := range datasources {
		value, err := ds.Value()
		if err != nil {
			logging.Error.Println("gRPC - DataSource: ", err)
			return err
		}
		res := &protobuf.DataSourcesResponse{
			Name:         ds.Name(),
			Description:  ds.Description(),
			Id:           ds.Id(),
			CurrentValue: value,
			ValueError:   "",
		}
		stream.Send(res)
	}

	return nil
}

func (*server) DataSource(ctx context.Context, req *protobuf.DataSourceRequest) (*protobuf.DataSourceResponse, error) {
	logging.Info.Printf("DataSource was invoked with %v\n", req)
	id := req.GetId()
	ds, err := datasources.GetDatasource(id)
	if err != nil {
		logging.Error.Println("gRPC - DataSource: ", err)
		return nil, err
	}
	value, err := ds.Value()
	if err != nil {
		logging.Error.Println("gRPC - DataSource: ", err)
		return nil, err
	}
	res := &protobuf.DataSourceResponse{
		CurrentValue: value,
		ValueError:   "",
	}
	return res, nil
}

func (*server) RPoint(ctx context.Context, req *protobuf.RPointRequest) (*protobuf.RPointResponse, error) {
	logging.Info.Printf("RPoint was invoked with %v\n", req)
	id := req.GetId()
	timestamp := req.GetTimestamp()
	rPoint, err := store.GetRPoint(id, uint64(timestamp))
	if err != nil {
		logging.Error.Println("gRPC - RPoint: ", err)
		return nil, err
	}
	res := &protobuf.RPointResponse{
		RPoint: hex.EncodeToString(rPoint[:]),
	}
	return res, nil
}

func (*server) Publication(ctx context.Context, req *protobuf.PublicationRequest) (*protobuf.PublicationResponse, error) {
	logging.Info.Printf("Publication was invoked with %v\n", req)
	passedR, err := hex.DecodeString(req.GetRPoint())
	if err != nil {
		logging.Error.Println("gRPC - Publication: ", err)
		return nil, err
	}

	var R [33]byte
	copy(R[:], passedR[:])

	value, signature, timestamp, name, err := store.GetPublication(R)
	if err != nil {
		logging.Error.Println("gRPC - Publication: ", err)
		return nil, err
	}
	res := &protobuf.PublicationResponse{
		Value:     value,
		Signature: hex.EncodeToString(signature[:]),
		Timestamp: timestamp,
		Name:      name,
	}
	return res, nil
}

func (*server) Publications(req *protobuf.PublicationsRequest, stream protobuf.OracleService_PublicationsServer) error {
	logging.Info.Printf("Publications was invoked with %v\n", req)
	name := fmt.Sprintf("%v/%v", req.GetBase(), req.GetQuote())

	publications, err := store.GetAllPublicationsByName(name)
	if err != nil {
		logging.Error.Println("gRPC - Publications: ", err)
		return err
	}

	for _, p := range publications {
		res := &protobuf.PublicationResponse{
			Value:     p.Value,
			Signature: hex.EncodeToString(p.Signature[:]),
			Timestamp: p.Timestamp,
			Name:      p.Name,
		}
		stream.Send(res)
	}

	return nil
}

// Init gRPC Server
func Init() {
	logging.Info.Println("Booting up gRPC Server...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		logging.Error.Fatalf("Failed to listen: %v", err)
	}

	tls := false
	opts := []grpc.ServerOption{}

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
		if sslErr != nil {
			logging.Error.Fatalf("Failed loading certificates: %v", sslErr)
			return
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	// protobuf.RegisterOracleServiceServer(s, &server{})
	protobuf.RegisterOracleServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		logging.Error.Fatalf("Failed to serve: %v", err)
	}
}
