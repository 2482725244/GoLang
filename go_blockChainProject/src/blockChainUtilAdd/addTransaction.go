package blockChainUtilAdd

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type TransInput struct {
	TransHash    []byte
	OutPutIndex  int64
	TransAddress string
}

type TransOutput struct {
	Money        int64
	TransAddress string
}

type Transaction struct {
	TransHash []byte
	TxInputs  []*TransInput
	TxOutputs []*TransOutput
}

func NewGenesisTransaction() *Transaction {
	transaction := &Transaction{TxInputs: []*TransInput{&TransInput{nil, -1, "This GenesisBlock"}},
		TxOutputs: []*TransOutput{&TransOutput{10, "江南"}}}
	transaction.TransHash = transaction.SetTransHash()
	return transaction
}

func (tx *Transaction) SetTransHash() []byte {
	tranceHash := tx.Serialization()
	Hash := sha256.Sum256(tranceHash)

	return Hash[:]
}
func (tx *Transaction) Serialization() []byte {

	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(tx)

	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}
