package blog

const (
	ModuleName = "blog"
	StoreKey   = ModuleName

	PostKey    = "post"
	CommentKey = "comment"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
