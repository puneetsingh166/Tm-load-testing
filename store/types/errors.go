package types

import (
	sdkerrors "github.com/puneetsingh166/tm-load-test/types/errors"
)

const StoreCodespace = "store"

var (
	ErrInvalidProof = sdkerrors.Register(StoreCodespace, 2, "invalid proof")
)
