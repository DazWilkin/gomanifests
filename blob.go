package gomanifests

import ()

type Blob struct{}
func (b *Blob) Layer() Layer {
    return Layer{}
}