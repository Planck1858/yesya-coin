package models

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"reflect"
	"strconv"
	"time"
)

type Block struct {
	Index        int       `json:"index" db:"index"`
	Timestamp    time.Time `json:"timestamp" db:"timestamp"`
	Data         string    `json:"data" db:"data"`
	Hash         string    `json:"hash" db:"hash"`
	PreviousHash string    `json:"previousHash" db:"previous_hash"`
}

type Blockchain struct {
	Blocks []Block `json:"blocks"`
}

func NewBlockchain() *Blockchain {
	var bc Blockchain
	genesisBlock := Block{
		Index:        0,
		Timestamp:    time.Date(2020, 2, 17, 00, 00, 0, 0, time.UTC),
		Data:         "My genesisBlock",
		Hash:         "a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447",
		PreviousHash: "",
	}
	bc.Blocks = append(bc.Blocks, genesisBlock)

	return &bc
}

func (bc *Blockchain) GetBlocks() []Block {
	return bc.Blocks
}

func (bc *Blockchain) NewBlock(data string) {
	bc.generateNextBlock(data)
}

func (bc *Blockchain) generateNextBlock(blockData string) {
	previousBlock := bc.GetLatestBlock()
	nextIndex := previousBlock.Index + 1
	nextTimestamp := time.Now()
	nextHash := calculateHash(nextIndex, previousBlock.Hash, nextTimestamp, blockData)
	var newBlock = Block{nextIndex, nextTimestamp, blockData,
		nextHash, previousBlock.Hash}

	if isValidChain(*bc) && isValidBlockStructure(newBlock) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Printf("Error - invalid new block")
	}
}

func calculateHash(index int, previousHash string, timestamp time.Time, data string) string {
	sum := sha256.Sum256([]byte(strconv.Itoa(index) + previousHash + timestamp.String() + data))

	return convByteToStr(sum)
}

func convByteToStr(bs [32]byte) string {
	return hex.EncodeToString(bs[:])
}

func (bc *Blockchain) GetLatestBlock() Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func isValidNewBlock(newBlock Block, previousBlock Block) bool {
	// TODO: check new Hash validity
	if previousBlock.Index+1 != newBlock.Index {
		log.Println("Error - Invalid Index !")
		return false
	} else if previousBlock.Hash != newBlock.PreviousHash {
		log.Println("Error - Invalid previoushash !")
		return false
	}

	return true
}

func isValidBlockStructure(block Block) bool {
	if reflect.TypeOf(block.Index).String() == "int" &&
		reflect.TypeOf(block.Hash).String() == "string" &&
		reflect.TypeOf(block.PreviousHash).String() == "string" &&
		reflect.TypeOf(block.Timestamp).String() == "time.Time" &&
		reflect.TypeOf(block.Data).String() == "string" {
		return true
	}

	log.Println("Error - Invalid Block Structure !")
	return false
}

func isValidChain(bc Blockchain) bool {
	for i := 1; i < len(bc.Blocks); i++ {
		if !isValidNewBlock(bc.Blocks[i], bc.Blocks[i-1]) {
			log.Println("Error - Invalid Chain !")
			return false
		}
	}
	return true
}
