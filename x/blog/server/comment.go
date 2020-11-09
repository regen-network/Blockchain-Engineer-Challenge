package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/amaurymartiny/bec/x/blog"
)

func (s serverImpl) CreateComment(ctx sdk.Context, comment blog.MsgComment) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), blog.KeyPrefix(blog.CommentKey))
	b := k.cdc.MustMarshalBinaryBare(&comment)
	store.Set(blog.KeyPrefix(blog.CommentKey), b)
}

func (s serverImpl) GetAllComment(ctx sdk.Context) (msgs []blog.MsgComment) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), blog.KeyPrefix(blog.CommentKey))
	iterator := sdk.KVStorePrefixIterator(store, blog.KeyPrefix(blog.CommentKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg blog.MsgComment
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
