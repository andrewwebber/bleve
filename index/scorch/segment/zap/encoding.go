package zap

import (
	"fmt"

	"github.com/golang/snappy"
)

type EncodingProvider interface {
	Encode(dst, src []byte) ([]byte, error)
	Decode(dst, src []byte) ([]byte, error)
}

type snappyProvider struct {
}

func NewSnappy() EncodingProvider {
	return &snappyProvider{}
}

func (s snappyProvider) Encode(dst, src []byte) ([]byte, error) {
	fmt.Println("zap encode")
	return snappy.Encode(dst, src), nil
}

func (s snappyProvider) Decode(dst, src []byte) ([]byte, error) {
	fmt.Println("zap decode")
	return snappy.Decode(dst, src)
}
