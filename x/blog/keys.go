package blog

import (
	"crypto/md5"
)

const (
	ModuleName = "blog"
	StoreKey   = ModuleName

	PostKey          = "post"
	CommentKeyPrefix = "comment"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func CommentAllKey(postSlug string) []byte {
	prefixBytes := []byte(CommentKeyPrefix)
	postSlugBytes := append([]byte(postSlug), byte('/'))
	return append(prefixBytes, postSlugBytes...)
}

func CommentKey(postSlug, author, body string) []byte {
	key := CommentAllKey(postSlug)
	commentBytes := append(bodyKey(author, body), byte('/'))
	return append(key, commentBytes...)
}

func bodyKey(author, body string) []byte {
	hasher := md5.New()
	hasher.Write([]byte(author + body))
	return hasher.Sum(nil)
}
