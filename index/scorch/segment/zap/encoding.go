package zap

import (
	"github.com/golang/snappy"
)

type encodingRegistry map[string]EncodingProvider

var registry encodingRegistry

func init() {
	registry = make(encodingRegistry)
	registry["snappy"] = NewSnappy()
}

type EncodingProvider interface {
	Encode(dst, src []byte) ([]byte, error)
	Decode(dst, src []byte) ([]byte, error)
}

type snappyProvider struct {
}

func New(name string) EncodingProvider {
	if p, ok := registry[name]; ok {
		return p
	}

	return registry["snappy"]
}

func NewSnappy() EncodingProvider {
	return &snappyProvider{}
}

func (s snappyProvider) Encode(dst, src []byte) ([]byte, error) {
	return snappy.Encode(dst, src), nil
}

func (s snappyProvider) Decode(dst, src []byte) ([]byte, error) {
	return snappy.Decode(dst, src)
}
