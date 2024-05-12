package util

type ServiceHash uint32

const (
	fnv_offset_basis ServiceHash = 0x811C9DC5
	fnv_prime        ServiceHash = 0x01000193
)

// returns Fowler–Noll–Vo 1A 32bit hash of string name
func HashServiceName(name string) (hash ServiceHash) {
	hash = fnv_offset_basis
	for i := 0; i < len(name); i++ {
		hash ^= ServiceHash(name[i])
		hash *= fnv_prime
	}

	return

}
