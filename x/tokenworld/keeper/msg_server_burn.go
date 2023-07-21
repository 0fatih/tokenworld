package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"tokenworld/x/tokenworld/types"

  "cosmossdk.io/errors"
)

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errors.Wrapf(types.ErrBurn, "invalid receiver address %v", msg.Creator)
	}

	// Take coins from the recipient
	err = k.bk.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, msg.Amount)
	if err != nil {
    return nil, errors.Wrapf(types.ErrReceiveCoins, "taking %v tokens from %s failed", msg.Amount, sender)
	}

	return &types.MsgBurnResponse{}, nil
}
