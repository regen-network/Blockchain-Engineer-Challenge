package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/bec/x/blog"
	uuid "github.com/uuid6/uuid6go-proto"
)

var _ blog.MsgServer = serverImpl{}

func (s serverImpl) CreatePost(goCtx context.Context, request *blog.MsgCreatePost) (*blog.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := prefix.NewStore(ctx.KVStore(s.storeKey), blog.KeyPrefix(blog.PostKey))

	key := []byte(request.Slug)
	if store.Has(key) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "duplicate slug %s found", request.Slug)
	}

	post := blog.Post{
		Author: request.Author,
		Slug:   request.Slug,
		Title:  request.Title,
		Body:   request.Body,
	}

	bz, err := s.cdc.Marshal(&post)
	if err != nil {
		return nil, err
	}

	store.Set(key, bz)

	return &blog.MsgCreatePostResponse{}, nil
}

func (s serverImpl) CreateComment(goCtx context.Context, request *blog.MsgCreateComment) (*blog.MsgCreateCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storePost := prefix.NewStore(ctx.KVStore(s.storeKey), blog.KeyPrefix(blog.PostKey))

	// We add comment to particular post thus we check the slug of the post exist in KVStore
	keyPost := []byte(request.Slug)
	if !storePost.Has(keyPost) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Post %s isn't written yet", request.Slug)
	}

	comment := blog.Comment{
		Slug:   request.Slug,
		Author: request.Author,
		Body:   request.Body,
	}

	bz, err := s.cdc.Marshal(&comment)
	if err != nil {
		return nil, err
	}

	// Comment has key is comprised by `comment` + `slug` + `_` + `uuidV7`.
	keyComment := blog.KeyPrefix(request.Slug + "_" + getUuid())
	//  Prefix `comment` is added by KVStore here.
	storeComment := prefix.NewStore(ctx.KVStore(s.storeKey), blog.KeyPrefix(blog.CommentKey))

	storeComment.Set(keyComment, bz)

	return &blog.MsgCreateCommentResponse{}, nil
}

// One post can have many comments. We are going to specify them with uuid
func getUuid() string {
	var gen uuid.UUIDv7Generator
	gen.SubsecondPrecisionLength = 12

	id := gen.Next()
	return id.ToString()
}
