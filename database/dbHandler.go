package database

import (
	"fmt"
	"github.com/Planck1858/yesya-coin/config"
	"github.com/Planck1858/yesya-coin/models"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var db *sqlx.DB

func Init(database string, bc *models.Blockchain) *sqlx.DB {
	var err error
	conf := config.GetConfig("./config/config.json")
	var (
		host     = conf.Host
		port     = conf.Port
		user     = conf.User
		password = conf.Password
		dbname   = conf.DbName
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	log.Printf("Database connect info: %s", psqlInfo)

	db, err = sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error - connecting to database: %s", err.Error())
		return nil
	}

	log.Printf("Successfully connected to %v, database %v", host, database)

	_, err = db.Exec(deleteBlocks)
	if err != nil {
		log.Printf("Error - wrong schema on 'deleteBlocks'")
		return nil
	}

	// init geneziz block
	db.MustExec(insertBlocks, bc.Blocks[0].Index, bc.Blocks[0].Timestamp,
		bc.Blocks[0].Data, bc.Blocks[0].Hash, bc.Blocks[0].PreviousHash)
	return db
}

func SelectAllBlocks(db *sqlx.DB) []models.Block {
	var blocks []models.Block
	err := db.Select(&blocks, "SELECT * FROM blocks")
	if err != nil {
		log.Printf("Error - select posts: %s", err.Error())
		return nil
	}

	return blocks
}

func NewBlock(db *sqlx.DB, bc *models.Blockchain) {
	var (
		index        = bc.GetLatestBlock().Index
		timestamp    = bc.GetLatestBlock().Timestamp
		data         = bc.GetLatestBlock().Data
		hash         = bc.GetLatestBlock().Hash
		previousHash = bc.GetLatestBlock().PreviousHash
	)

	db.MustExec(insertBlocks, index, timestamp, data, hash, previousHash)
}
