package keeper

import (
	"github.com/mconcat/ibc-eth/x/ibceth/types"
)

var _ types.QueryServer = Keeper{}
