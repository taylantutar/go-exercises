package util

import (
    "math/rand"
    "time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
    rand.Seed(time.Now().UnixNano())
}

func GenerateShortCode() string {
    length := 8
    code := make([]byte, length)
    for i := range code {
        code[i] = charset[rand.Intn(len(charset))]
    }
    return string(code)
}
