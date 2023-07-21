package types

// DONTCOVER

import (
	errors "cosmossdk.io/errors"
)

// x/tokenworld module sentinel errors
var (
	ErrMint         = errors.Register(ModuleName, 1100, "mint coins failed")
	ErrSendMint     = errors.Register(ModuleName, 1101, "send minted coins failed")
	ErrBurn         = errors.Register(ModuleName, 1102, "burn coins failed")
	ErrReceiveCoins = errors.Register(ModuleName, 1103, "receive coins failed")
)
