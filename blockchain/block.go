package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
	BlockHeight  int
	Nonce        int
}

// func (b *Block) CalculateBlockHash(){
// 	info:= bytes.Join([][]byte{b.Data, b.PreviousHash, {byte(b.BlockHeight)}}, []byte{})
// 	hash:=sha256.Sum256(info)
// 	b.Hash=hash[:]
// }

func CreateBlock(data string, previousHash []byte, blockHeight int) *Block {
	block := &Block{[]byte{}, []byte(data), previousHash, blockHeight, 0}
	// block.CalculateBlockHash()
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis block", []byte{}, 1000000)
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(b)
	CheckError(err)
	return res.Bytes()
}

func Deserializer(data []byte) *Block {
	var block Block
	deocder := gob.NewDecoder(bytes.NewReader(data))
	err := deocder.Decode(&block)
	CheckError(err)

	return &block
}

func CheckError(err error){
	if(err!=nil){
		log.Panic(err)
	}
}