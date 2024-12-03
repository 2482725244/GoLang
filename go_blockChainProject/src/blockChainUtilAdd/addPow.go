package blockChainUtilAdd

import (
	"bytes"
	"crypto/sha256"
	"math/big"
)

const TARGETBIT = 16

type Pow struct {
	block  *Block
	target *big.Int
}

func CreatePow(block *Block) *Pow {

	Target := big.NewInt(1)
	Target = Target.Lsh(Target, 32*8-TARGETBIT)

	return &Pow{
		block:  block,
		target: Target,
	}
}

func (pow *Pow) Run() ([]byte, int64) {

	var nonce int64 = 0
	var hash [32]byte
	var hashInt big.Int

	for {
		info := pow.combine(nonce)
		hash = sha256.Sum256(info)
		hashInt.SetBytes(hash[:])
		if pow.target.Cmp(&hashInt) == 1 {
			break
		}
		nonce++
	}

	return hash[:], nonce
}

func (pow *Pow) combine(nonce int64) []byte {
	info := bytes.Join([][]byte{ToHexInt(pow.block.Timestamp), ToHexInt(nonce),
		pow.block.PreHash, pow.block.TransToHash()}, []byte{})
	return info
}
