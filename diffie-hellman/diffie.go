// Package diffiehellman implements a Diffie-Hellman-Merkle key exchange
package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

// PrivateKey generates a private key
func PrivateKey(p *big.Int) *big.Int {
	prKey := new(big.Int)
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	upperLimit := new(big.Int).Sub(p, big.NewInt(2))
	return prKey.Rand(r, upperLimit).Add(prKey, big.NewInt(2))

}

// PublicKey A = g**private mod p
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

// SecretKey s = A**b mod p
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}

// NewPair produces working keys
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}
