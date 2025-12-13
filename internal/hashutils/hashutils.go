package hashutils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Hash(s string) string {
	digest := md5.New()
	digest.Write([]byte(s))
	return hex.EncodeToString(digest.Sum(nil))
}
