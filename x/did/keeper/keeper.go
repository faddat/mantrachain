package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/MANTRA-Finance/aumega/x/did/types"
)

// UnmarshalFn is a generic function to unmarshal bytes
type UnmarshalFn func(value []byte) (interface{}, bool)

// MarshalFn is a generic function to marshal bytes
type MarshalFn func(value interface{}) []byte

type Keeper struct {
	cdc      codec.Codec
	storeKey storetypes.StoreKey
	memKey   storetypes.StoreKey
	gk       types.GuardKeeper
}

func NewKeeper(cdc codec.Codec, storeKey, memKey storetypes.StoreKey,
	gk types.GuardKeeper) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		gk:       gk,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Set sets a value in the db with a prefixed key
func (k Keeper) Set(ctx sdk.Context,
	key []byte,
	prefix []byte,
	i interface{},
	marshal MarshalFn,
) {
	store := ctx.KVStore(k.storeKey)
	store.Set(append(prefix, key...), marshal(i))
}

// Delete - deletes a value form the store
func (k Keeper) Delete(
	ctx sdk.Context,
	key []byte,
	prefix []byte,
) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(append(prefix, key...))
}

func (k Keeper) Has(
	ctx sdk.Context,
	key []byte,
	prefix []byte,
) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(append(prefix, key...))
}

// Get gets an item from the store by bytes
func (k Keeper) Get(
	ctx sdk.Context,
	key []byte,
	prefix []byte,
	unmarshal UnmarshalFn,
) (i interface{}, found bool) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(append(prefix, key...))

	return unmarshal(value)
}

// GetAll values from with a prefix from the store
func (k Keeper) GetAll(
	ctx sdk.Context,
	prefix []byte,
) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, prefix)
}
