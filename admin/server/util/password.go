package util

import (
	"crypto/md5"
	"github.com/anaskhan96/go-password-encoder"
	"math/rand"
	"time"
)

var options = password.Options{SaltLen: 10, Iterations: 10000, KeyLen: 50, HashFunction: md5.New}

func GetSaltAndEncodedPassword(pwd string) (string, string) {
	salt, encodedPwd := password.Encode(pwd, &options)
	return salt, encodedPwd
}

func VerifyRawPassword(rawPwd, encodedPwd, salt string) bool {
	return password.Verify(rawPwd, salt, encodedPwd, &options)
}

func GeneratePassword(length int) []byte {
	baseStr := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	baseSymbol := "!@#$%^&"

	pwdStr := baseStr + baseSymbol
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63()))
	bytes := make([]byte, length)
	l := len(pwdStr)

	for i := 0; i < length; i++ {
		bytes[i] = pwdStr[r.Intn(l)]
	}

	return bytes
}
