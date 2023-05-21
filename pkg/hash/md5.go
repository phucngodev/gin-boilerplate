package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(src string) string {
	return getResult(src)
}

func MD5WithSalt(src string, salt string) string {
	str := src + "#" + salt
	return getResult(str)
}

func getResult(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}
