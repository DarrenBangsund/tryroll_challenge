package grpc

import (
	"context"
	"fmt"
	"net"
	"tryroll/challenge/cmd/tryroll-service/internal/app"
	pb "tryroll/challenge/pkg/grpc"

	"google.golang.org/grpc"
)

// TryRollGRPCServer is the grpc implementation
type TryRollGRPCServer struct {
	pb.UnimplementedTryRollServiceServer
	app *app.App
}

// NewTryRollGRPCServer returns a new grpc server with the configured app
func NewTryRollGRPCServer(app *app.App) *TryRollGRPCServer {
	return &TryRollGRPCServer{
		app: app,
	}
}

// NewGRPCServerConn configures a listener and returns a new grpc server and listener
func NewGRPCServerConn(addr string) (*grpc.Server, net.Listener, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8002))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to listen: %w", err)
	}

	return grpc.NewServer(), lis, nil
}

// GetERC20Tx implements the grpc method to return objects filtered by the given filter parameters.
// this method supports pagination
func (t *TryRollGRPCServer) GetERC20Tx(
	ctx context.Context,
	req *pb.GetERC20TxReq,
) (*pb.GetERC20TxRes, error) {
	txs, err := t.app.GetERC20TokenTx(
		req.GetAmount(),
		req.GetTo(),
		req.GetFrom(),
		req.GetLimit(),
		req.GetOffset(),
	)
	if err != nil {
		return nil, err
	}

	pbTxs := []*pb.ERC20Tx{}

	for _, tx := range txs {
		pbTxs = append(pbTxs, &pb.ERC20Tx{
			To:     tx.To,
			From:   tx.From,
			Hash:   tx.Hash,
			Amount: tx.Value.String(),
		})
	}

	return &pb.GetERC20TxRes{
		Txs: pbTxs,
	}, nil
}
