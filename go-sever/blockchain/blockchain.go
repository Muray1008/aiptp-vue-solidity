package blockchain

import (
    "fmt"
    "time"
	"strings"
	"strconv"
	"crypto/sha256"
	"encoding/hex"
)

// Block data structure
type Block struct {
    Id 			int 	`json:"id"`
    Nonce  		int		`json:"nonce"`
	Timestamp	string 	`json:"timestamp"`
    Data 		string 	`json:"data"`
    PrevHash	string 	`json:"prevHash"`
    Hash		string 	`json:"hash"`
}

// Retrieve these objects from db instance
var blockchain []Block

// Init a new blockchain
func InitBlockchain(data string) {
	prevHash := "0000000000000000000000000000000000000000000000000000000000000000"
	hash, nonce := findHash(prevHash)

	newBlock := createNewBlock(0, nonce, data, prevHash, hash)

    blockchain = append(blockchain, newBlock)
}

// Get timestamp
func GetTimestamp() string {
    return time.Now().String()
}

// Find a new valid hash
func findHash(prevHash string) (string, int) {

	condition := true
	nonce := 0
	hash := ""

	for condition {
		nonce++

		proof := []byte(strconv.Itoa(nonce*2) + prevHash)

		sha256Proof := sha256.Sum256(proof)
		hash = hex.EncodeToString(sha256Proof[:])

		if strings.HasPrefix(hash, "0000") { // Check hash condition
			condition = false
		}

	}
	fmt.Println("\nHash found: ", hash, " at hounce ", nonce)
	return hash, nonce
}

// Get lash block
func getLastBlock() Block {
	return blockchain[len(blockchain)-1]
}

// Create a new block
func createNewBlock(id int, nonce int, data string, prevHash string, hash string) Block {
    block := Block{id, nonce, GetTimestamp(), data, prevHash, hash}
    return block
}

// Mine a new block
func MineBlock(data string) *Block {
    var lastBlock Block

    // Check lenght chain
	if (len(blockchain) == 0) {
		fmt.Println("Init new chain...")
		InitBlockchain(data)
		lastBlock = blockchain[len(blockchain)-1]

	} else {
	    // Get previous block info
		lastBlock := getLastBlock();

        // Get previous hash
		prevHash := lastBlock.Hash;

		// New block
		hash, nonce := findHash(prevHash)
		newId := lastBlock.Id + 1
		newBlock := createNewBlock(newId, nonce, data, prevHash, hash)
		blockchain = append(blockchain, newBlock)

		lastBlock = newBlock
	}

	return &lastBlock
}

// Get the blockchain
func GetBlockchain() []Block {
	return blockchain
}
