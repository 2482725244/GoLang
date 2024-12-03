package blockChainUtilAdd

import (
	"flag"
	"fmt"
	"os"
)

type BlockChainCMD struct {
	Bc *BlockChain
}

func (bcmd *BlockChainCMD) addBlock(data string) {
	if IsHaveDB() == false {
		fmt.Println("-------->还没有建立创世区块")
		os.Exit(1)
	} else {
		bc := GetDBAndBlockChain()
		bcmd.Bc = bc

		bcmd.Bc.AddBlock([]*Transaction{})
		fmt.Println("addBlock-->success")
	}
}

func (bcmd *BlockChainCMD) printBlockChain() {
	if IsHaveDB() == false {
		fmt.Println("-------->还没有建立创世区块")
		os.Exit(1)
	} else {
		bc := GetDBAndBlockChain()
		bcmd.Bc = bc

		bcmd.Bc.BlockChainPrint()
	}

}

func IsHaveDB() bool {

	if _, err := os.Stat(DBNAME); os.IsNotExist(err) {
		return false
	}
	return true
}

func (bcmd *BlockChainCMD) pintUsage() {

	fmt.Println()
	fmt.Println("Usage")
	fmt.Println("-------->addBlock -data \"String Content\"")
	fmt.Println("-------->printBlocks")
	fmt.Println("-------->creatGenesisBlock")
	fmt.Println()

}

func InitBlockChainCMD() *BlockChainCMD {
	return &BlockChainCMD{}
}

func (bcmd *BlockChainCMD) createGenesis() {
	if IsHaveDB() == false {
		//创建数据库
		bc := InitBlockChain()
		bcmd.Bc = bc

	} else {
		fmt.Println("------>创世区块已存在")
		os.Exit(1)
	}
}

func (bcmd *BlockChainCMD) Run() {

	addBlockFlagSet := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printBlocksSet := flag.NewFlagSet("printBlocks", flag.ExitOnError)
	createGenesisSet := flag.NewFlagSet("createGenesis", flag.ExitOnError)

	data := addBlockFlagSet.String("data", "调用函数添加新区快", "addBlock-->")

	if len(os.Args) < 2 {
		bcmd.pintUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "addBlock":
		addBlockFlagSet.Parse(os.Args[2:])
		bcmd.addBlock(*data)

	case "createGenesisBlock":
		createGenesisSet.Parse(os.Args[2:])
		bcmd.createGenesis()
		fmt.Println("-------->createGenesisSuccess!")

	case "printBlocks":
		printBlocksSet.Parse(os.Args[2:])
		bcmd.printBlockChain()

	default:
		bcmd.pintUsage()
		os.Exit(1)
	}

}
