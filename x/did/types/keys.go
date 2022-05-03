package types

const (
	// ModuleName defines the module name
	ModuleName = "did"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	DidMethod = "mns"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	DidKey          = "did:"
	DidCountKey     = "did-count:"
	DidNamespaceKey = "did-namespace:"
)
