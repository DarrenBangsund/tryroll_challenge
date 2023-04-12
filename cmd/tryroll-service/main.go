package main

import (
	"fmt"
	"log"
	"tryroll/challenge/cmd/tryroll-service/internal/app"
	ipb "tryroll/challenge/cmd/tryroll-service/internal/grpc"
	"tryroll/challenge/cmd/tryroll-service/internal/odm"
	"tryroll/challenge/pkg/etherscan"
	pb "tryroll/challenge/pkg/grpc"
)

// TODO: put these in env var
const (
	// EtherscanMock is a boolean defined to mock the Etherscan pkg for testing
	EtherscanMock = false
	// EtherscanURL is the base url for the Etherscan pkg to make requests against (not used when EtherscanMock is true)
	EtherscanURL = "https://api.etherscan.io/api"
	// EtherscanKey is the API key for etherscan (not used when EtherscanMock is true)
	EtherscanKey = ""
)

func main() {
	es := etherscan.NewEtherscan(&etherscan.EtherscanConfig{
		BaseURL: EtherscanURL,
		APIKey:  EtherscanKey,
		Mock:    EtherscanMock,
	})

	log.Printf("Fetching erc20 tx from Etherscan...")
	transactions, err := es.Account.GetERC20TokenTx(
		"0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2",
		"0x4e83362442b8d1bec281594cea3050c8eb01311c",
		"0",
		"100",
		"1",
	)
	if err != nil {
		panic(err)
	}
	log.Printf("Done")

	log.Printf("Creating ODM...")
	odm := odm.NewODM()

	log.Printf("Creating App...")
	app := app.NewApp(odm)

	log.Printf("Adding %d records to the datastore...", len(transactions))
	// insert initial transactions
	for _, t := range transactions {
		err := app.AddERC20TokenTx(t.To, t.From, t.Value, t.Hash)
		if err != nil {
			panic(err)
		}
	}

	log.Printf("Creating GRPC server...")
	grpcServer, listener, err := ipb.NewGRPCServerConn(fmt.Sprintf(":%d", 8002))

	pb.RegisterTryRollServiceServer(
		grpcServer,
		ipb.NewTryRollGRPCServer(app),
	)
	log.Printf("Server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
