package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/tirith-tech/dlc-oracle/rpc/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	fmt.Println("Booting up a client...")

	tls := false
	opts := grpc.WithInsecure()
	if tls {
		certFile := "ssl/ca.crt" // Certificate Authority Trust certificate
		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
			return
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := protobuf.NewOracleServiceClient(cc)

	getPubKey(c)
	getDataSources(c)
	getDataSource(c)
	getRPoint(c)
	getPublication(c)
	getPublications(c)
}

func getPubKey(c protobuf.OracleServiceClient) {
	fmt.Println("getPubKey Unary RPC call...")
	req := &emptypb.Empty{}
	res, err := c.PubKey(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling PubKey RPC: %v", err)
	}
	log.Printf("Response from getPubKey: %v", res.Pubkey)
}

func getDataSources(c protobuf.OracleServiceClient) {
	fmt.Println("getDataSources Streaming RPC call...")
	req := &emptypb.Empty{}
	resStream, err := c.DataSources(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling getDataSources RPC: %v", err)
	}

	for {
		ds, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("%v [%v] %v [current value in sats: %v]", ds.GetName(), ds.GetId(), ds.GetDescription(), ds.GetCurrentValue())
	}
}

func getDataSource(c protobuf.OracleServiceClient) {
	fmt.Println("getDataSource Unary RPC call...")
	req := &protobuf.DataSourceRequest{
		Id: 1, // Get current price for datasource=1
	}
	res, err := c.DataSource(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling getDataSource RPC: %v", err)
	}
	log.Printf("Response from getDataSource: %v", res.CurrentValue)
}

func getRPoint(c protobuf.OracleServiceClient) {
	fmt.Println("getRPoint Unary RPC call...")
	req := &protobuf.RPointRequest{
		Id:        1, // Get current price for datasource=1
		Timestamp: 1604606100,
	}
	res, err := c.RPoint(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling getRPoint RPC: %v", err)
	}
	log.Printf("Response from getRPoint: %v", res.RPoint)
}

func getPublication(c protobuf.OracleServiceClient) {
	fmt.Println("getPublication Unary RPC call...")
	req := &protobuf.PublicationRequest{
		RPoint: "03410f8f338e2f3e8b1848af11a6d2ce49d489486cff45ce390a0f4b919cefb9ed",
	}
	pub, err := c.Publication(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling getPublication RPC: %v", err)
	}
	log.Printf("%v [ts %v] [value in sats: %v]  [signature: %v]", pub.GetName(), pub.GetTimestamp(), pub.GetValue(), pub.GetSignature())
}

func getPublications(c protobuf.OracleServiceClient) {
	fmt.Println("getPublications Streaming RPC call...")
	req := &protobuf.PublicationsRequest{
		Base:  "LTC",
		Quote: "BTC",
	}
	resStream, err := c.Publications(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling getPublication RPC: %v", err)
	}

	for {
		pub, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("%v [ts %v] [value in sats: %v]  [signature: %v]", pub.GetName(), pub.GetTimestamp(), pub.GetValue(), pub.GetSignature())
	}
}
