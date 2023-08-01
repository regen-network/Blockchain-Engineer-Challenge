package blog

const (
	ModuleName = "blog"
	StoreKey   = ModuleName

	PostKey        = "post"
	PostCommentKey = "post_comment"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
