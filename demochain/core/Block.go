package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index int64
	Timestamp int64
	PreBlockHash string
	Hash  string

	Data string
}

func calculateHash(b Block) string  {
	blockData := string(b.Index)  + string(b.Timestamp)  + b.PreBlockHash + b.Data
	hashByte := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashByte[:])
}


func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.PreBlockHash = preBlock.Hash
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

func GenerateGenesisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock,"Genesis Block")
}