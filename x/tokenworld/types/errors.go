package types

// DONTCOVER

import (
	errors "cosmossdk.io/errors"
)

// x/tokenworld module sentinel errors
var (
	ErrMint = errors.Register(ModuleName, 1100, "mint coins failed")
	ErrSendMint = errors.Register(ModuleName, 1101, "send minted coins failed")
)
