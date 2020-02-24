package models

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	index        int       `json:"index"`
	hash         string    `json:"hash"`
	previousHash string    `json:"previous_hash"`
	timestamp    time.Time `json:"timestamp"`
	data         string    `json:"data"`
}

type Blockchain struct {
	Blocks []Block `json:"blocks"`
}

func NewBlockchain() *Blockchain {
	var bc Blockchain
	genesisBlock := Block{
		index:        0,
		hash:         "a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447",
		previousHash: "",
		timestamp:    time.Date(2020, 2, 17, 12, 12, 0, 0, time.UTC),
		data:         "My genesisBlock",
	}
	bc.Blocks = append(bc.Blocks, genesisBlock)

	return &bc
}

func (bc *Blockchain) GetBlockchain() []Block {
	return bc.Blocks
}

func (bc *Blockchain) NewBlock(data string) {
	bc.generateNextBlock(data)
}

func (bc *Blockchain) generateNextBlock(blockData string) {
	previousBlock := bc.getLatestBlock()
	nextIndex := previousBlock.index + 1
	nextTimestamp := time.Now()
	nextHash := calculateHash(nextIndex, previousBlock.hash, nextTimestamp, blockData)
	var newBlock = Block{nextIndex, nextHash, previousBlock.hash, nextTimestamp, blockData}

	bc.Blocks = append(bc.Blocks, newBlock)
}

func calculateHash(index int, previousHash string, timestamp time.Time, data string) string {
	sum := sha256.Sum256([]byte(strconv.Itoa(index) + previousHash + timestamp.String() + data))
	return convByteToStr(sum)
}

func convByteToStr(bs [32]byte) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = byte(v)
	}
	return string(b)
}

func (bc *Blockchain) getLatestBlock() Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func isValidNewBlock(newBlock Block, previousBlock Block) bool {
	// TODO: check new hash validity
	if previousBlock.index+1 != newBlock.index {
		fmt.Println("Invalid index !")
		return false
	} else if previousBlock.hash != newBlock.previousHash {
		fmt.Println("Invalid previoushash !")
		return false
	}
	return true
}

//func isValidBlockStructure(block Block) bool {
//	if reflect.TypeOf(block.index).String() == "int" &&
//		reflect.TypeOf(block.hash).String() == "string" &&
//		reflect.TypeOf(block.previousHash).String() == "string" &&
//		reflect.TypeOf(block.timestamp).String() == "time.Time" &&
//		reflect.TypeOf(block.data).String() == "string" {
//		return true
//	}
//	fmt.Println("Invalid Block Structure !")
//	return false
//}
//
//func isValidChain(bc Blockchain) bool {
//	for i := 1; i < len(bc.Blocks); i++ {
//		if !isValidNewBlock(bc.Blocks[i], bc.Blocks[i-1]) {
//			fmt.Println("Invalid Chain !")
//			return false
//		}
//	}
//	return true
//}
