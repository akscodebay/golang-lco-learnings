package main

import (
	cryptorand "crypto/rand"
	"fmt"
	"math/big"
	mathrand "math/rand"
)

func main() {
	fmt.Println("Math and Crypto Rand in Golang")
	fmt.Println("Random value from math rand: ", mathrand.Intn(5)+1)
	randomNumbFromCrypto, _ := cryptorand.Int(cryptorand.Reader, big.NewInt(5))
	fmt.Println("Random value from crypto: ", randomNumbFromCrypto.Add(randomNumbFromCrypto, big.NewInt(1)))
}
