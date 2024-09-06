package client

import "github.com/fbsobreira/gotron-sdk/pkg/proto/api"

func (g *GrpcClient) GetTransactionListFromPending() (*api.TransactionIdList, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	return g.Client.GetTransactionListFromPending(
		ctx,
		new(api.EmptyMessage),
	)
}
