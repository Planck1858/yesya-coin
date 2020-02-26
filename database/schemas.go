package database

var insertBlocks = `INSERT INTO blocks (index, timestamp, data, hash, previous_hash) VALUES ($1, $2, $3, $4, $5)`
var deleteBlocks = `DELETE FROM blocks;`
