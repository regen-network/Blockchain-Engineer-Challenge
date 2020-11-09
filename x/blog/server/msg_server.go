package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"

	"github.com/amaurymartiny/bec/x/blog"
)

var _ blog.MsgServer = serverImpl{}

func (s serverImpl) CreatePost(goCtx context.Context, request *blog.MsgCreatePostRequest) (*blog.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := prefix.NewStore(ctx.KVStore(s.storeKey), blog.KeyPrefix(blog.PostKey))

	id := uuid.New().String()
	post := blog.Post{
		Id:     id,
		Author: request.Author,
		Title:  request.Title,
		Body:   request.Body,
	}

	bz, err := s.cdc.MarshalBinaryBare(&post)
	if err != nil {
		return nil, err
	}

	store.Set(blog.KeyPrefix(blog.PostKey), bz)

	return &blog.MsgCreatePostResponse{Id: id}, nil
}
