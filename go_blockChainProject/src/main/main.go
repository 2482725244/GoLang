package main

import (
	"fmt"
	"src/blockChainUtilAdd"
	"time"
)

func main() {

	fmt.Println("Hello BlockChian")
	time.Sleep(time.Second)

	bcmd := blockChainUtilAdd.InitBlockChainCMD()
	bcmd.Run()

}

//func main() {
//
//	fmt.Println("Hello BlockChian")
//	time.Sleep(time.Second)
//
//	bc := block.InitBlockChain()
//	defer bc.DB.Close()
//
//	bc.AddBlock([]byte("Java程序员的自我修养"))
//	bc.AddBlock([]byte("软件工程导论"))
//	bc.AddBlock([]byte("软件程序的运维"))
//	bc.AddBlock([]byte("c语言程序基础"))
//	bc.AddBlock([]byte("智能合约solidity"))
//	bc.AddBlock([]byte("c语言程序基础副本"))
//	bc.AddBlock([]byte("智能合约solidity副本"))
//
//	bc.BlockChainPrint()
//
//}
