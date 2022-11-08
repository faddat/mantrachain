package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/LimeChain/mantrachain/x/vault/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc                codec.BinaryCodec
		storeKey           sdk.StoreKey
		memKey             sdk.StoreKey
		paramstore         paramtypes.Subspace
		ac                 types.AccountKeeper
		bk                 types.BankKeeper
		sk                 types.StakingKeeper
		dk                 types.DistrKeeper
		nftKeeper          types.NFTKeeper
		bridgeKeeper       types.BridgeKeeper
		wasmViewKeeper     types.WasmViewKeeper
		wasmContractKeeper types.WasmContractOpsKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

	ac types.AccountKeeper,
	bk types.BankKeeper,
	sk types.StakingKeeper,
	dk types.DistrKeeper,
	nftKeeper types.NFTKeeper,
	bridgeKeeper types.BridgeKeeper,
	wasmViewKeeper types.WasmViewKeeper,
	wasmContractKeeper types.WasmContractOpsKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		ac:         ac, bk: bk, sk: sk, dk: dk, nftKeeper: nftKeeper, bridgeKeeper: bridgeKeeper,
		wasmViewKeeper:     wasmViewKeeper,
		wasmContractKeeper: wasmContractKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
