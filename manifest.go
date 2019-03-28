package gomanifests

import (
	"hash"
)

type Manifest struct {
	schemaVersion SchemaVersion
	mediaType     string
	config        Config
	layers        []Layer
}
type SchemaVersion uint8
type Config struct {
	mediaType string
	size      int32
	digest    hash.Hash
}
type Layer struct {
	mediaType string
	size      int32
	digest    hash.Hash
}
