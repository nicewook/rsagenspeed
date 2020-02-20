package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		_, _ = rsa.GenerateKey(rand.Reader, 4096)
		fmt.Println("since", i, ": ", time.Since(start))
		start = time.Now()
	}
	for i := 0; i < 10; i++ {
		_, _ = rsa.GenerateKey(rand.Reader, 2048)
		fmt.Println("since", i, ": ", time.Since(start))
		start = time.Now()
	}
}
