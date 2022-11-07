package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strings"
)

type BlockChain struct{
	blocks []*Block
}


type Block struct{
	Hash [] byte
	Data [] byte
	PreviousHash[] byte
}

func (b *Block) CalculateBlockHash(){
	info:= bytes.Join([][]byte{b.Data, b.PreviousHash}, []byte{})
	hash:=sha256.Sum256(info)
	b.Hash=hash[:]
}

func CreateBlock(data string, previousHash []byte) *Block{
	block:=&Block{[]byte{},[]byte(data),previousHash}
	block.CalculateBlockHash()
	return block
}

func (chain *BlockChain) AddBlock(data string){
	previousBlock:= chain.blocks[len(chain.blocks) -1]
	newBlock := CreateBlock(data, previousBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

func Genesis() *Block{
	return CreateBlock("Genesis block", []byte{})
}

func InitBlockchain() *BlockChain{
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	adaigoChain:=InitBlockchain()
	adaigoChain.AddBlock("First block in adaigochain")
	adaigoChain.AddBlock("second block in adaigochain")
	adaigoChain.AddBlock("third block in adaigochain")
	adaigoChain.AddBlock("Okay, lets add some data")

	for _,block := range adaigoChain.blocks{
		fmt.Println(strings.Repeat("-", 84))
		fmt.Printf("Previous Hash : %x \n", block.PreviousHash)
		fmt.Printf("Data          : %x \n", block.Data)
		fmt.Printf("Block Hash    : %x \n", block.Hash)
		fmt.Println(strings.Repeat("-", 84))
		fmt.Println(strings.Repeat(" ", 42)+"|")
		fmt.Println(strings.Repeat(" ", 42)+"|")
	}
}
