package main

import "testing"

func BenchmarkRandInt10Times(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandInt10Times()
	}
}

func BenchmarkRegexGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RegexGenerate()
	}
}

func BenchmarkCryptoRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CryptoRand()
	}
}

func BenchmarkCryptoRandMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CrytoRandMath()
	}
}
