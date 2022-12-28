package utils

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

const (
	salt = "hjqrhjqw124617ajfhajs"
)

func FuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func GenerateHash(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
