package main

import (
	"fmt"
	"strings"

	"github.com/keerthanchand/adaigo-blockchain/blockchain"
)

func main() {
	adaigoChain := blockchain.InitBlockchain()
	adaigoChain.AddBlock("First block in adaigochain")
	adaigoChain.AddBlock("second block in adaigochain")
	adaigoChain.AddBlock("third block in adaigochain")
	adaigoChain.AddBlock("Okay, lets add some data {take data in please take data in some more data bro please take somemore data in iknow there is some place inside }")

	for _, block := range adaigoChain.Blocks {
		fmt.Println(strings.Repeat("-", 84))
		fmt.Printf("Previous Hash : %x \n", block.PreviousHash)
		fmt.Printf("Block Height  : %x \n", block.BlockHeight)
		fmt.Printf("Data          : %x \n", block.Data)
		fmt.Printf("Nonce         : %x \n", block.Nonce)
		fmt.Printf("Block Hash    : %x \n", block.Hash)
		fmt.Println(strings.Repeat("-", 84))
		fmt.Println(strings.Repeat(" ", 42) + "|")
		fmt.Println(strings.Repeat(" ", 42) + "|")
	}
}
