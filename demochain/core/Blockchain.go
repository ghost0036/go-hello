package core

import (
	"log"
	"fmt"
)

type Blockchain struct {
	Blocks []*Block
}


func NewBlockchain() *Blockchain{
	genesisBlock := GenerateGenesisBlock()
	blockchain := Blockchain{}
	blockchain.AppendBlock(&genesisBlock)
	return &blockchain
}


func (bc *Blockchain)SendData(data string)  {
	preBlock := bc.Blocks[len(bc.Blocks) - 1]
	newBlock := GenerateNewBlock(*preBlock,data)
	bc.AppendBlock(&newBlock)
}

func (bc *Blockchain)AppendBlock(newBlock *Block)  {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks,newBlock)
		return
	}
	if isValidate(*newBlock,*bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	}else {
		log.Fatal("invalid block")
	}

}

func (bc *Blockchain)Print(){
	for _,block := range bc.Blocks {
		fmt.Printf("Index: %d \n",block.Index)
		fmt.Printf("PreBlockHash: %s\n",block.PreBlockHash)
		fmt.Printf("Timestamp: %d\n",block.Timestamp)
		fmt.Printf("CurrentBlockHash:%s\n",block.Hash)
		fmt.Printf("Data: %s\n",block.Data)
		fmt.Println()
	}
}

func isValidate(newBlock,oldBlock Block) bool  {
	if newBlock.Index != oldBlock.Index + 1 {
		return false
	}

	if newBlock.PreBlockHash != oldBlock.Hash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}
