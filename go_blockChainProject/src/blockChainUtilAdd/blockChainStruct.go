package blockChainUtilAdd

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"time"
)

type Block struct {
	Timestamp int64
	Nonce     int64
	Height    int64
	ThisHash  []byte
	PreHash   []byte
	Txs       []*Transaction
}

const DBNAME = "BlockChains.db"
const BOLTNAME = "FBlockChain"

type BlockChain struct {
	DB            *bolt.DB
	lateBlockHash []byte
}

func (b *Block) Serialization() []byte {

	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)

	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeSerialization(blockSerBytes []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockSerBytes))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}

// Deprecated
func (b *Block) SetHash() {
	info := bytes.Join([][]byte{ToHexInt(b.Timestamp), b.PreHash, b.TransToHash()}, []byte{})
	hash := sha256.Sum256(info)
	b.ThisHash = hash[:]
}

func (b *Block) TransToHash() []byte {
	var txsJoin [][]byte
	var txsHash []byte

	for _, tx := range b.Txs {
		txsJoin = append(txsJoin, tx.TransHash)
	}
	txsHash = bytes.Join(txsJoin, []byte{})
	lateHash := sha256.Sum256(txsHash)
	return lateHash[:]
}

func ToHexInt(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

func (bc *BlockChain) CreateBlock(preHash []byte, txs []*Transaction, height int64) *Block {
	block := &Block{
		Timestamp: time.Now().Unix(),
		PreHash:   preHash,
		Txs:       txs,
		Height:    height,
	}
	//block.SetHash()

	//pow := pow.CreatePow(block)
	pow := CreatePow(block)
	hash, nonce := pow.Run()
	block.Nonce = nonce
	block.ThisHash = hash
	return block
}

func (bc *BlockChain) createWorldBock() *Block {
	transaction := NewGenesisTransaction()
	return bc.CreateBlock([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]*Transaction{transaction}, 1)
}

func InitBlockChain() *BlockChain {
	blockChain := &BlockChain{}
	var block *Block

	db, err := bolt.Open(DBNAME, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {

		b, err := tx.CreateBucket([]byte(BOLTNAME))
		if err != nil {
			log.Panic(err)
		}

		block = blockChain.createWorldBock()

		err = b.Put(block.ThisHash, block.Serialization())
		if err != nil {
			log.Panic(err)
		}
		err = b.Put([]byte("l"), block.ThisHash)
		if err != nil {
			log.Panic(err)
		}

		return nil
	})

	blockChain.DB = db
	blockChain.lateBlockHash = block.ThisHash
	return blockChain
}

func GetDBAndBlockChain() *BlockChain {
	blockChain := &BlockChain{}

	db, err := bolt.Open(DBNAME, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BOLTNAME))
		lateHash := b.Get([]byte("l"))
		blockChain.lateBlockHash = lateHash
		return nil
	})

	blockChain.DB = db
	return blockChain
}

func (bc *BlockChain) AddBlock(txs []*Transaction) {

	err := bc.DB.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(BOLTNAME))
		if b != nil {

			blockByte := b.Get(bc.lateBlockHash)
			block := DeSerialization(blockByte)
			newBlock := bc.CreateBlock(block.ThisHash, txs, block.Height+1)
			b.Put(newBlock.ThisHash, newBlock.Serialization())
			b.Put([]byte("l"), newBlock.ThisHash)
			bc.lateBlockHash = newBlock.ThisHash
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

//	func (bc *BlockChain) BlockChainPrint() {
//		bc.DB.View(func(tx *bolt.Tx) error {
//			b := tx.Bucket([]byte(BOLTNAME))
//
//			if b != nil {
//
//				for {
//					var bigHashInt big.Int
//
//					blockByte := b.Get(bc.lateBlockHash)
//					block := DeSerialization(blockByte)
//
//					fmt.Println("区块", block.Height)
//					fmt.Printf("创建时间戳为:%v\n", block.Timestamp)
//					fmt.Printf("父区块Hash为:0X%x\n", block.PreHash)
//					fmt.Printf("自身区块Hash为:0X%x\n", block.ThisHash)
//					fmt.Printf("区块数据为:%s\n", block.ThisData)
//					fmt.Printf("测试nonce为:%d\n", block.Nonce)
//					fmt.Printf("区块高度为:%d\n", block.Height)
//					fmt.Println()
//
//					bigHashInt.SetBytes(block.PreHash)
//					if bigHashInt.Cmp(big.NewInt(0)) == 0 {
//						break
//					}
//					bc.lateBlockHash = block.PreHash
//				}
//			}
//
//			return nil
//		})
//
// }
func (bc *BlockChain) BlockChainPrint() {
	bt := initBlockChainIter(bc)
	block := bt.Next()

	for block != nil {
		fmt.Println("区块", block.Height)
		fmt.Printf("创建时间戳为:%v\n", block.Timestamp)
		fmt.Printf("父区块Hash为:0X%x\n", block.PreHash)
		fmt.Printf("自身区块Hash为:0X%x\n", block.ThisHash)
		fmt.Printf("测试nonce为:%d\n", block.Nonce)
		fmt.Printf("区块高度为:%d\n", block.Height)
		fmt.Println("交易记录{")
		for key, tx := range block.Txs {
			fmt.Printf("\t交易记录%d\n", key+1)
			fmt.Printf("\t交易哈希为：%x\n", tx.TransHash)
			fmt.Println("\tInputs内容为[")
			for key, input := range tx.TxInputs {
				fmt.Printf("\t\t\t第%d个input内容为:{交易哈希：%x,位于索引为：%d,用户名为：%s}", key+1, input.TransHash, input.OutPutIndex, input.TransAddress)
			}
			fmt.Println("\n\t\t\t]")
			fmt.Println("\tOutputs内容为[")
			for key, output := range tx.TxOutputs {
				fmt.Printf("\t\t\t第%d个output内容为:{获得金额：%d,用户名：%s}", key+1, output.Money, output.TransAddress)
			}
			fmt.Println("\n\t\t\t]")
		}
		fmt.Println("\t}")
		fmt.Println()
		block = bt.Next()
	}
}
