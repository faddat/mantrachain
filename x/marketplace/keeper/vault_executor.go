package keeper

import (
	"github.com/LimeChain/mantrachain/x/marketplace/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type VaultExecutor struct {
	ctx         sdk.Context
	vaultKeeper types.VaultKeeper
}

func NewVaultExecutor(ctx sdk.Context, vaultKeeper types.VaultKeeper) *VaultExecutor {
	return &VaultExecutor{
		ctx:         ctx,
		vaultKeeper: vaultKeeper,
	}
}

func (c *VaultExecutor) UpsertNftStake(
	marketplaceIndex []byte,
	collectionIndex []byte,
	nftIndex []byte,
	creator sdk.AccAddress,
	amount sdk.Coin,
	delegate bool,
	stakingChain string,
	stakingValidator string,
) (bool, error) {
	return c.vaultKeeper.UpsertNftStake(c.ctx, marketplaceIndex, collectionIndex, nftIndex, creator, amount, delegate, stakingChain, stakingValidator)
}
