package types

const (
	// ModuleName defines the module name
	ModuleName = "cernwebfestpostquantumcryptography"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cernwebfestpostquantumcryptography"
)

var (
	ParamsKey = []byte("p_cernwebfestpostquantumcryptography")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
