package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("测试组")
	name := os.Args[1]
	flagSetAddBlock := flag.NewFlagSet("add", flag.ExitOnError)
	string := flagSetAddBlock.String("addBlock", "ccc", "addblockmore")
	flagSetAddBlock.Parse(os.Args[2:])

	fmt.Println(name)
	fmt.Println(*string)

}
