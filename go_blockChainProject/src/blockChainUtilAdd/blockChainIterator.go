package blockChainUtilAdd

import (
	"github.com/boltdb/bolt"
	"math/big"
)

type BlockChainIterator struct {
	DB            *bolt.DB
	lateBlockHash []byte
}

func initBlockChainIter(bc *BlockChain) *BlockChainIterator {
	return &BlockChainIterator{bc.DB, bc.lateBlockHash}
}

func (bt *BlockChainIterator) Next() *Block {
	var BigHashInt big.Int
	BigHashInt.SetBytes(bt.lateBlockHash)
	if BigHashInt.Cmp(big.NewInt(0)) == 0 {
		return nil
	}

	var block *Block
	bt.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BOLTNAME))

		if b != nil {
			blockByte := b.Get(bt.lateBlockHash)
			block = DeSerialization(blockByte)
			bt.lateBlockHash = block.PreHash
		}
		return nil
	})
	return block
}
