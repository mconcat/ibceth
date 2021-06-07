package ibceth

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mconcat/ibc-eth/x/ibceth/keeper"
	"github.com/mconcat/ibc-eth/x/ibceth/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.DefaultGenesis()
}
