package gomanifests

import (
    "crypto/sha256"
    "hash"
)

type Blob struct{}
func (b *Blob) Layer() Layer {
    return Layer{}
}
func (b *Blob) Digest() hash.Hash {
    return sha256.New()
}