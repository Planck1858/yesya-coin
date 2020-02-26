# yesya-coin
This project is a small implementation of blockchain technology on Golang

## Initialization
```
go mod tidy
go run main.go
```

## Usage  
Just enter this in URL path:
```
/blocks (GET) - gets all the blocks of blockchain
/mineBlock (POST) - ganerates a new block with data from the body of http request. 
```
*Blocks are generated both in memory and in the connected database (PostgreSQL)*
