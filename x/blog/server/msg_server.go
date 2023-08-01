package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/bec/x/blog"
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

func (s serverImpl) CreatePostComment(goCtx context.Context, request *blog.MsgCreatePostComment) (*blog.MsgCreatePostCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	commentStore := prefix.NewStore(ctx.KVStore(s.storeKey), blog.KeyPrefix(blog.PostCommentKey))
	postStore := prefix.NewStore(ctx.KVStore(s.storeKey), blog.KeyPrefix(blog.PostKey))

	// check if the post exists
	if !postStore.Has([]byte(request.Slug)) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "slug %s is found n 'post' storage", request.Slug)
	}

	comment := blog.PostComment{
		Author: request.Author,
		Slug:   request.Slug,
		Body:   request.Body,
	}

	bComment, err := s.cdc.Marshal(&comment)
	if err != nil {
		return nil, err
	}

	// For key we can use body + author + slug
	// explain: if the author write several similar message for one post comment will have no value
	// In my opinion it's best solution for comment key
	// Also we can try to generate some key, like UUID or something and for sure it's also will be working solution
	key := []byte(request.Body + request.Author + request.Slug)

	commentStore.Set(key, bComment)

	return &blog.MsgCreatePostCommentResponse{}, nil
}
