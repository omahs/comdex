package keeper

import (
	"context"

	"github.com/comdex-official/comdex/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServiceServer = (*queryServer)(nil)

type queryServer struct {
	Keeper
}

func NewQueryServiceServer(k Keeper) types.QueryServiceServer {
	return &queryServer{
		Keeper: k,
	}
}

func (q *queryServer) QueryParams(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	var (
		ctx    = sdk.UnwrapSDKContext(c)
		params = q.GetParams(ctx)
	)

	return &types.QueryParamsResponse{
		Params: params,
	}, nil
}