package server

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/amaurymartiny/bec/x/blog"
)

func (k serverImpl) createPost(ctx sdk.Context, post blog.MsgPost) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), blog.KeyPrefix(blog.PostKey))
	b := k.cdc.MustMarshalBinaryBare(&post)
	store.Set(blog.KeyPrefix(blog.PostKey), b)
}

func (k serverImpl) getAllPost(ctx sdk.Context) (msgs []blog.MsgPost) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), blog.KeyPrefix(blog.PostKey))
	iterator := sdk.KVStorePrefixIterator(store, blog.KeyPrefix(blog.PostKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg blog.MsgPost
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
