package main

import "testing"

func TestCheckPrime(t *testing.T) {
	isPrime(-1)
	isPrime(0)
	isPrime(1)
	isPrime(10)
	isPrime(12)
}
