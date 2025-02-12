package types

import (
	"fmt"
	"github.com/tendermint/crypto/blake2b"
)

func (b *Bundle) ID() string {
	//TODO lazy compute and cache this

	hash, err := blake2b.New256(nil)

	if err != nil {
		// not expecting this
		panic(fmt.Sprintf("failed to create blake2b hash: %s", err))
	}
	bytes, err := b.Marshal()
	if err != nil {
		panic(fmt.Sprintf("failed to marshal bundle: %s", err))
	}

	hash.Write(bytes)

	sum := hash.Sum(nil)
	return fmt.Sprintf("%x", sum) //hexencode for readability
}
