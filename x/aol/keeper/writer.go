package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/medibloc/panacea-core/v2/types/compkey"
	"github.com/medibloc/panacea-core/v2/x/aol/types"
)

// SetWriter set a specific writer in the store
func (k Keeper) SetWriter(ctx sdk.Context, key types.WriterCompositeKey, writer types.Writer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.WriterKeyPrefix)
	b := k.cdc.MustMarshal(&writer)
	store.Set(compkey.MustEncode(&key), b)
}

// GetWriter returns a writer from its id
func (k Keeper) GetWriter(ctx sdk.Context, key types.WriterCompositeKey) types.Writer {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.WriterKeyPrefix)
	var writer types.Writer
	k.cdc.MustUnmarshal(store.Get(compkey.MustEncode(&key)), &writer)
	return writer
}

// HasWriter checks if the writer exists in the store
func (k Keeper) HasWriter(ctx sdk.Context, key types.WriterCompositeKey) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.WriterKeyPrefix)
	return store.Has(compkey.MustEncode(&key))
}

// RemoveWriter removes a writer from the store
func (k Keeper) RemoveWriter(ctx sdk.Context, key types.WriterCompositeKey) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.WriterKeyPrefix)
	store.Delete(compkey.MustEncode(&key))
}

// GetAllWriters returns all writers
func (k Keeper) GetAllWriters(ctx sdk.Context) ([]types.WriterCompositeKey, []types.Writer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.WriterKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	keys := make([]types.WriterCompositeKey, 0)
	values := make([]types.Writer, 0)

	for ; iterator.Valid(); iterator.Next() {
		var key types.WriterCompositeKey
		compkey.MustDecode(iterator.Key(), &key)
		keys = append(keys, key)

		var value types.Writer
		k.cdc.MustUnmarshal(iterator.Value(), &value)
		values = append(values, value)
	}

	return keys, values
}
