package hash

import "crypto/sha256"

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a Adapter) Hash(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}
