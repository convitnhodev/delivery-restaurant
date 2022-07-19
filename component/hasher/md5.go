package hasher

import (
	"crypto/md5"
	"encoding/hex"
)

type md5Hash struct{}

func NewMD5Hash() *md5Hash {
	return &md5Hash{}
}

func (h *md5Hash) Hash(data string) string {
	// hash
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
