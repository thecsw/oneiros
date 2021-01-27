package main

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"strconv"
	"time"

	regen "github.com/zach-klippenstein/goregen"
)

const (
	Pattern = `^[0-9]{16}$`
)

var (
	Big10 = big.NewInt(10)
)

func RandInt10Times() string {
	rand.Seed(time.Now().UnixNano())
	res := ""
	for i := 0; i < 16; i++ {
		res += strconv.FormatInt(rand.Int63n(10), 10)
	}
	return res
}

func RegexGenerate() string {
	res, _ := regen.Generate(Pattern)
	return res
}

func CryptoRand() string {
	res := ""
	for i := 0; i < 16; i++ {
		n, _ := crand.Int(crand.Reader, Big10)
		res += n.String()
	}
	return res
}

func CrytoRandMath() string {
	res := big.NewInt(0)
	for i := 0; i < 16; i++ {
		n, _ := crand.Int(crand.Reader, Big10)
		res.Add(res.Mul(res, Big10), n)
	}
	return res.String()
}
