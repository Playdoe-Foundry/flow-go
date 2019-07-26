package collect

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	collectSvc "github.com/dapperlabs/bamboo-node/pkg/grpc/services/collect"
)

type Controller struct {
	dal *DAL
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Ping(context.Context, *collectSvc.PingRequest) (*collectSvc.PingResponse, error) {
	return &collectSvc.PingResponse{
		Address: []byte("pong!"),
	}, nil
}

// SubmitTransaction accepts an incoming transaction from a user agent or peer node.
//
// This function will return an error in the follow cases:
//
// The request is malformed or incomplete.
// The transaction has an invalid or missing signature.
//
// The submitted transaction will be stored for future inclusion in a collection
// if it belongs to this node's cluster, but otherwise it will be forwarded to the
// correct cluster.
func (c *Controller) SubmitTransaction(
	ctx context.Context,
	req *collectSvc.SubmitTransactionRequest,
) (*collectSvc.SubmitTransactionResponse, error) {
	tx := req.GetTransaction()

	if tx == nil {
		return nil, status.Error(codes.InvalidArgument, "")
	}

	// TODO: validate transaction contents
	// https://github.com/dapperlabs/bamboo-node/issues/170

	// TODO: validate transaction signature
	// https://github.com/dapperlabs/bamboo-node/issues/171

	// TODO: store transaction
	// https://github.com/dapperlabs/bamboo-node/issues/169

	return nil, status.Error(codes.Unimplemented, "")
}

func (c *Controller) SubmitCollection(context.Context, *collectSvc.SubmitCollectionRequest) (*collectSvc.SubmitCollectionResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (c *Controller) GetTransaction(context.Context, *collectSvc.GetTransactionRequest) (*collectSvc.GetTransactionResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (c *Controller) GetCollection(context.Context, *collectSvc.GetCollectionRequest) (*collectSvc.GetCollectionResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
