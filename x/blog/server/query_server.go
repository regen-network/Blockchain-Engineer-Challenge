package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/bec/x/blog"
)

var _ blog.QueryServer = serverImpl{}

func (s serverImpl) AllPosts(goCtx context.Context, request *blog.QueryAllPostsRequest) (*blog.QueryAllPostsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(s.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, blog.KeyPrefix(blog.PostKey))

	defer iterator.Close()

	var posts []*blog.Post
	for ; iterator.Valid(); iterator.Next() {
		var msg blog.Post
		err := s.cdc.Unmarshal(iterator.Value(), &msg)
		if err != nil {
			return nil, err
		}

		posts = append(posts, &msg)
	}

	return &blog.QueryAllPostsResponse{
		Posts: posts,
	}, nil
}

func (s serverImpl) AllComments(goCtx context.Context, request *blog.QueryAllCommentsRequest) (*blog.QueryAllCommentsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(s.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, blog.KeyPrefix(blog.CommentKey))

	defer iterator.Close()

	var comments []*blog.Comment
	for ; iterator.Valid(); iterator.Next() {
		var msg blog.Comment
		err := s.cdc.Unmarshal(iterator.Value(), &msg)
		if err != nil {
			return nil, err
		}

		comments = append(comments, &msg)
	}

	return &blog.QueryAllCommentsResponse{
		Comments: comments,
	}, nil
}
