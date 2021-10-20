package types

import (
	sdkerrors "github.com/puneetsingh166/tm-load-test/types/errors"
)

// x/crisis module sentinel errors
var (
	ErrNoSender         = sdkerrors.Register(ModuleName, 2, "sender address is empty")
	ErrUnknownInvariant = sdkerrors.Register(ModuleName, 3, "unknown invariant")
)
