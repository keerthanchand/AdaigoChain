package blockchain


type BlockChain struct{
	Blocks []*Block
}


type Block struct{
	Hash [] byte
	Data [] byte
	PreviousHash[] byte
	BlockHeight int
	Nonce int
}

// func (b *Block) CalculateBlockHash(){
// 	info:= bytes.Join([][]byte{b.Data, b.PreviousHash, {byte(b.BlockHeight)}}, []byte{})
// 	hash:=sha256.Sum256(info)
// 	b.Hash=hash[:]
// }

func CreateBlock(data string, previousHash []byte, blockHeight int) *Block{
	block:=&Block{[]byte{},[]byte(data),previousHash, blockHeight, 0}
	// block.CalculateBlockHash()
	pow:= NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func (chain *BlockChain) AddBlock(data string){
	previousBlock:= chain.Blocks[len(chain.Blocks) -1]
	newBlock := CreateBlock(data, previousBlock.Hash, previousBlock.BlockHeight+1)
	chain.Blocks = append(chain.Blocks, newBlock)
}

func Genesis() *Block{
	return CreateBlock("Genesis block", []byte{}, 1000000)
}

func InitBlockchain() *BlockChain{
	return &BlockChain{[]*Block{Genesis()}}
}
