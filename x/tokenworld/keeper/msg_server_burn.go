package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"tokenworld/x/tokenworld/types"
)

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgBurnResponse{}, nil
}
