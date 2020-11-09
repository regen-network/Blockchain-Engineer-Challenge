package blog

const (
	ModuleName = "blog"
	StoreKey   = ModuleName

	PostKey = "post"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
