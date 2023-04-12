package helpers

import (
	"tryroll/challenge/cmd/gql-api/internal/graph/generated"
	"tryroll/challenge/pkg/grpc"
)

// ConvertERC20TxsGRPCToGQL will convert a list of grpc ERC20Tx objects to gql objects
func ConvertERC20TxsGRPCToGQL(input []*grpc.ERC20Tx) []*generated.ERC20Tx {
	gqltxs := []*generated.ERC20Tx{}

	for _, tx := range input {
		gqltxs = append(gqltxs, &generated.ERC20Tx{
			To:     tx.GetTo(),
			From:   tx.GetFrom(),
			Amount: tx.GetAmount(),
			Hash:   tx.GetHash(),
		})
	}

	return gqltxs
}
