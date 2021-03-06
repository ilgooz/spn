package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/profile module sentinel errors
var (
	ErrCoordAlreadyExist = sdkerrors.Register(ModuleName, 1, "coordinator address already exist")

	// this line is used by starport scaffolding # ibc/errors
)
