package keeper

import (
	"context"
	"tokenworld/x/tokenworld/types"

	errors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Mint coins to the module
	err := k.bk.MintCoins(ctx, types.ModuleName, msg.Amount)
	if err != nil {
		return nil, errors.Wrapf(types.ErrMint, "minting %v failed", msg.Amount)
	}

	receiver, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errors.Wrapf(types.ErrSendMint, "invalid receiver address %v", msg.Creator)
	}

	// Send minted coins to the recipient
	err = k.bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, msg.Amount)
	if err != nil {
		return nil, errors.Wrapf(types.ErrSendMint, "sending %v to %v failed", msg.Amount, sdk.AccAddress(msg.Creator))
	}

	return &types.MsgMintResponse{}, nil
}
