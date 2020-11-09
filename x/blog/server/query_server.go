package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/types"

	"github.com/amaurymartiny/bec/x/blog"
)

var _ blog.QueryServer = serverImpl{}

func (s serverImpl) AllPost(goCtx context.Context, request *blog.QueryAllPostRequest) (*blog.QueryAllPostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	cid := request.Cid

	var timestamp types.Timestamp
	err := s.anchorTable.GetOne(ctx, cid, &timestamp)
	if err != nil {
		return nil, err
	}

	var signers blog.Signers
	// ignore error because we at least have the timestamp
	_ = s.signersTable.GetOne(ctx, cid, &signers)

	store := ctx.KVStore(s.storeKey)
	content := store.Get(DataKey(cid))

	return &blog.QueryAllPostResponse{
		Timestamp: &timestamp,
		Signers:   signers.Signers,
		Content:   content,
	}, err
}

func (s serverImpl) AllPost(goCtx context.Context, request *blog.QueryAllPostRequest) (*blog.QueryAllPostResponse, error) {
	panic("implement me")
}
