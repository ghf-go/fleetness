package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

// 生产md5
func Md5(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}
func Sha1(data string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(data)))
}
func Sha256(data string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
}
func Sha512(data string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(data)))
}
func HmacSha1(key, data string) string {
	h := hmac.New(sha1.New, []byte(key))
	return hex.EncodeToString(h.Sum([]byte(data)))
}
func HmacSha256(key, data string) string {
	h := hmac.New(sha256.New, []byte(key))
	return hex.EncodeToString(h.Sum([]byte(data)))
}
func HmacSha512(key, data string) string {
	h := hmac.New(sha512.New, []byte(key))
	return hex.EncodeToString(h.Sum([]byte(data)))
}
