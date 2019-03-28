package gomanifests

import ()

type Account struct {
	owner    Party
	registry Registry
	name     string
}

func NewAccount(owner Party, registry Registry, name string) *Account {
	return &Account{
		owner:    owner,
		registry: registry,
		name:     name,
	}
}
func (a *Account) NewRepository(name string) *Repository {
    return &Repository{
        
    }
}