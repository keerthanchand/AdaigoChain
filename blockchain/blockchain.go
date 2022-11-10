package blockchain

import("github.com/dgraph-io/badger")

const (
	dbPath="./temp/blocks"
)
type BlockChain struct{
	LastHash []byte
	Database *badger.DB
}


func InitBlockchain() *BlockChain{
	var lastHash []byte
	opts:=badger.DefaultOptions
	opts.Dir=dbPath
	opts.ValueDir = dbPath
	db *DB
	db, err : badger.Open(opts)
	CheckError(err)
	
	return &BlockChain{[]*Block{Genesis()}}
}

func (chain *BlockChain) AddBlock(data string){
	previousBlock:= chain.Blocks[len(chain.Blocks) -1]
	newBlock := CreateBlock(data, previousBlock.Hash, previousBlock.BlockHeight+1)
	chain.Blocks = append(chain.Blocks, newBlock)
}


