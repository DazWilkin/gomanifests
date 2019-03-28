package gomanifests

import ()

type Registry struct {
	owner Party
	root  string
}

func NewRegistry(owner Party, root string) *Registry {
	return &Registry{
		owner: owner,
		root:  root,
	}
}
func (r *Registry) NewAccount(owner Party) *Account {
    return &Account{

    }
}
func (r *Registry) NewImage(repository *Repository, blobs []Blob) *Image {
    return &Image{

    }
}
