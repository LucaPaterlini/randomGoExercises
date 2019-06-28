// Package diffiehellman implements the generation and calculation of public and private keys.
package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey calculates the private key starting from the p.
func PrivateKey(p *big.Int) (r *big.Int) {
	r = big.NewInt(0)
	r, _ = rand.Int(rand.Reader, r.Sub(p, big.NewInt(3)))
	return r.Add(r, big.NewInt(2))
}

// PublicKey calculate the public key using pair (p,g) and private a.
func PublicKey(a, p *big.Int, g int64) (r *big.Int) {
	r = big.NewInt(0)
	return r.Exp(big.NewInt(g), a, p)
}

// NewPair calculates a new private and public pair
// starting from the primes p,g.
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	a := PrivateKey(p)
	return a, PublicKey(a, p, g)
}

// SecretKey calculate the secret s using Alice private and Bob public.
func SecretKey(a, B, p *big.Int) (r *big.Int) {
	r = big.NewInt(0)
	return r.Exp(B, a, p)
}
