package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	gtypes "github.com/nabakirov/hellochain/x/greeter/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc *codec.Codec
}

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc: cdc,
	}
}

func (k Keeper) GetGreetings(ctx sdk.Context, addr sdk.AccAddress, from sdk.Address) gtypes.GreetingList {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(addr)) {
		return gtypes.GreetingList{}
	}
	bz := store.Get([]byte(addr))
	var list gtypes.GreetingList
	k.cdc.MustUnmarshalBinaryBare(bz, &list)

	if from != nil {
		var fromList gtypes.GreetingList
		for _, g := range list {
			if g.Sender.Equals(from) {
				fromList = append(fromList, g)
			}
		}
		return fromList
	}
	return list
}

func (k Keeper) SetGreeting(ctx sdk.Context, greeting gtypes.Greeting) {
	if greeting.Sender.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	list := k.GetGreetings(ctx, greeting.Recipient, nil)
	list = append(list, greeting)
	store.Set(greeting.Recipient.Bytes(), k.cdc.MustMarshalBinaryBare(list))
}


func (k Keeper) GetGreetingsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}



