package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func main() {
	fmt.Println("Password: ", passwd)
}

