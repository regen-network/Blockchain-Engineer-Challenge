package server

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

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

func (s serverImpl) PostComments(goCtx context.Context, request *blog.QueryPostCommentsRequest) (*blog.QueryPostCommentsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	slug := request.Slug
	store := ctx.KVStore(s.storeKey)
	iteratorComments := sdk.KVStorePrefixIterator(store, blog.KeyPrefix(blog.CommentKey))

	defer iteratorComments.Close()

	var comments []*blog.Comment
	// we itirate over keys with prefix CommentKey ("comment")
	for ; iteratorComments.Valid(); iteratorComments.Next() {
		key := string(iteratorComments.Key())
		// Drop "_" and uuid of the key
		// Now slugFromComment = comment+slug
		slugFromComment := strings.Split(key, "_")[0]
		if slugFromComment == "" {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "can not extract slug from the key: %s", key)
		}

		if blog.CommentKey+slug == slugFromComment {
			var msg blog.Comment
			err := s.cdc.Unmarshal(iteratorComments.Value(), &msg)
			if err != nil {
				return nil, err
			}
			comments = append(comments, &msg)
		}
	}

	return &blog.QueryPostCommentsResponse{
		Comments: comments,
	}, nil
}
