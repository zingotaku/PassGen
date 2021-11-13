package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func md5sum(s string) string {
	h := md5.New()

	io.WriteString(h, s)

	return hex.EncodeToString(h.Sum(nil)[:16])
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" +
	"*$%&()"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(1000 * time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

var randomInt int = rand.Int()
var passwd string = (StringWithCharset(20, charset)) + string(randomInt)

var encryptedpasswd string = md5sum(passwd)

func main() {
	fmt.Println("Password: ", encryptedpasswd)
}
