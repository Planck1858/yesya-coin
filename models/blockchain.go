package models

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type Block struct {
	index        int
	hash         string
	previousHash string
	timestamp    time.Time
	data         string
}

type Blockchain []Block

func NewBlockchain() Blockchain {
	b := make(Blockchain, 0)
	genesisBlock := Block{
		index:        0,
		hash:         "a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447",
		previousHash: nil,
		timestamp:    time.Date(2020, 2, 17, 12, 12, 0, 0, time.UTC),
		data:         "My genesisBlock",
	}
	b[0] = genesisBlock
	return b
}

func (b *Blockchain) generateNextBlock(blockData string) Block {
	previousBlock := b.getLatestBlock()
	nextIndex := previousBlock.index + 1
	nextTimestamp := time.Now()
	nextHash := calculateHash(nextIndex, previousBlock.hash, nextTimestamp, blockData)
	newBlock := Block{
		index:        nextIndex,
		hash:         nextHash,
		previousHash: previousBlock.hash,
		timestamp:    nextTimestamp,
		data:         blockData,
	}
	return newBlock
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

func (b Blockchain) getLatestBlock() Block {
	return b[len(b)-1]
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

func isValidBlockStructure(block Block) bool {
	if reflect.TypeOf(block.index).String() == "int" &&
		reflect.TypeOf(block.hash).String() == "string" &&
		reflect.TypeOf(block.previousHash).String() == "string" &&
		reflect.TypeOf(block.timestamp).String() == "time.Time" &&
		reflect.TypeOf(block.data).String() == "string" {
		return true
	}
	fmt.Println("Invalid Block Structure !")
	return false
}

func isValidChain(bc Blockchain) bool {
	for i := 1; i < len(bc); i++ {
		if !isValidNewBlock(bc[i], bc[i - 1]) {
			fmt.Println("Invalid Chain !")
			return false
		}
	}
	return true
}
