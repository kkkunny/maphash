package maphash

import "unsafe"

// Hasher2 hashes values of type K.
// Uses runtime AES-based hashing.
type Hasher2[K any] struct {
	hash hashfn
	seed uintptr
}

// NewHasher2 creates a new Hasher2[K] with a random seed.
func NewHasher2[K any]() Hasher2[K] {
	return Hasher2[K]{
		hash: getRuntimeHasher2[K](),
		seed: newHashSeed(),
	}
}

// NewSeed2 returns a copy of |h| with a new hash seed.
func NewSeed2[K comparable](h Hasher2[K]) Hasher2[K] {
	return Hasher2[K]{
		hash: h.hash,
		seed: newHashSeed(),
	}
}

// Hash hashes |key|.
func (h Hasher2[K]) Hash(key K) uint64 {
	return uint64(h.Hash2(key))
}

// Hash2 hashes |key| as more flexible uintptr.
func (h Hasher2[K]) Hash2(key K) uintptr {
	// promise to the compiler that pointer
	// |p| does not escape the stack.
	p := noescape(unsafe.Pointer(&key))
	return h.hash(p, h.seed)
}
