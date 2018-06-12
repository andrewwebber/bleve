package zap

import (
	"github.com/blevesearch/bleve/index/scorch/segment"
	"github.com/golang/snappy"
)

type encodingRegistry map[string]segment.EncodingProvider

var registry encodingRegistry

func init() {
	registry = make(encodingRegistry)
	registry["snappy"] = NewSnappy()
}

func RegisterEncodingProvider(name string, provider segment.EncodingProvider) {
	registry[name] = provider
}

type snappyProvider struct {
}

func New(name string) segment.EncodingProvider {
	if p, ok := registry[name]; ok {
		return p
	}

	return registry["snappy"]
}

func NewSnappy() segment.EncodingProvider {
	return &snappyProvider{}
}

func (s snappyProvider) Encode(dst, src []byte) ([]byte, error) {
	return snappy.Encode(dst, src), nil
}

func (s snappyProvider) Decode(dst, src []byte) ([]byte, error) {
	return snappy.Decode(dst, src)
}
