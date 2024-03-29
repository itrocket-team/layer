package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tellor-io/layer/utils"
	"github.com/tellor-io/layer/x/oracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetUserTipTotal(goCtx context.Context, req *types.QueryGetUserTipTotalRequest) (*types.QueryGetUserTipTotalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	tipper := sdk.MustAccAddressFromBech32(req.Tipper)

	var totalTips types.UserTipTotal
	if len(req.QueryData) == 0 {
		// if query data is not provide, return total tips for the tipper on all queries
		totalTips, err := k.GetUserTips(ctx, tipper)
		if err != nil {
			return nil, err
		}
		return &types.QueryGetUserTipTotalResponse{TotalTips: &totalTips}, nil
	}
	// if query data is provided, return total tips for the tipper on the specific query
	queryId := utils.QueryIDFromData(req.QueryData)

	tip, err := k.Tips.Get(ctx, collections.Join(queryId, tipper.Bytes()))
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return &types.QueryGetUserTipTotalResponse{TotalTips: &totalTips}, nil
		}
		return nil, err
	}
	totalTips.Total = tip

	return &types.QueryGetUserTipTotalResponse{TotalTips: &totalTips}, nil
}
