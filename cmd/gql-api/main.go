package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"tryroll/challenge/cmd/gql-api/helpers"
	"tryroll/challenge/cmd/gql-api/internal/graph/generated"
	pb "tryroll/challenge/pkg/grpc"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = "8888"

type Resolver struct {
	tryrollClient pb.TryRollServiceClient
}

//	func (r *Resolver) Mutation() generated.MutationResolver {
//		return &mutationResolver{
//		}
//	}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{
		tryrollClient: r.tryrollClient,
	}
}

type queryResolver struct {
	tryrollClient pb.TryRollServiceClient
}

func (q *queryResolver) GetERC20Tx(
	ctx context.Context,
	input generated.GetERC20TxReq,
) (*generated.GetERC20TxRes, error) {
	res, err := q.tryrollClient.GetERC20Tx(ctx, &pb.GetERC20TxReq{
		Amount: helpers.SafeDerefString(input.Amount),
		To:     helpers.SafeDerefString(input.To),
		From:   helpers.SafeDerefString(input.From),
		Limit:  helpers.SafeDerefString(input.Limit),
		Offset: helpers.SafeDerefString(input.Offset),
	})
	if err != nil {
		//todo: sanitize error return
		return nil, err
	}

	return &generated.GetERC20TxRes{
		Txs: helpers.ConvertERC20TxsGRPCToGQL(res.GetTxs()),
	}, nil
}

func main() {

	conn, err := grpc.Dial("service:8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(fmt.Errorf("failed to connect to tryroll-service GRPC: %w", err))
	}

	trClient := pb.NewTryRollServiceClient(conn)

	gqlConfig := generated.Config{
		Resolvers: &Resolver{
			tryrollClient: trClient,
		},
	}

	handler := handler.NewDefaultServer(generated.NewExecutableSchema(gqlConfig))

	r := chi.NewRouter()
	r.Use(
		cors.New(cors.Options{
			AllowedOrigins:   strings.Split("*", ","),
			AllowedMethods:   []string{"POST"},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: true,
			Debug:            false,
		}).Handler,
	)

	r.Handle("/playground", playground.Handler("GraphQL playground", "/gql"))
	r.Handle("/gql", handler)

	log.Printf("GQL API listening on http://localhost:%s/gql", port)
	log.Printf("connect to http://localhost:%s/playground for GraphQL playground", port)

	log.Fatal(http.ListenAndServe(":"+port, r))
}
