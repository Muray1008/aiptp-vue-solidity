package main

import (
    blockchain "./blockchain"
    "fmt"
	"encoding/json"
    "net/http"
)

// New block request
type blockRequest struct {
    Data string
}

// Get Blockchain response
type BlockchainResponse struct {
	Blockchain 	[]blockchain.Block  `json:"blockchain"`
}

// New block response
type NewBlockResponse struct {
	Message 	string  `json:"message"`
}

// Get enable corse header
func enableCors(w *http.ResponseWriter) {
	//(*w).Header().Set("Access-Control-Allow-Origin", "*")

	  (*w).Header().Set("Content-Type", "text/html; charset=ascii")
      (*w).Header().Set("Access-Control-Allow-Origin", "*")
      (*w).Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")

}

// Add a new block
func newBlock(w http.ResponseWriter, req *http.Request) {
    enableCors(&w)
    if req.Method == http.MethodOptions {
        return
    }
	fmt.Println("\nAdding a new block...")

	decoder := json.NewDecoder(req.Body)
    var br blockRequest
    err := decoder.Decode(&br)
    if err != nil {
        panic(err)
    }

    // Mine a new block
    block := blockchain.MineBlock(br.Data)
	fmt.Println("\nNew block: ", block)

	newBlockResponse := NewBlockResponse{"Created"}

	response, err := json.Marshal(newBlockResponse)
	if err != nil {
	    fmt.Println("\nJSON marshal error")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
	    fmt.Println("\nNew block added: ", response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

// Get blockchain list
func getBlockchain(w http.ResponseWriter, req *http.Request) {
	fmt.Println("\nGetting Blockchain...")
	enableCors(&w)
    list := blockchain.GetBlockchain()
	var blockchainBlockchainSlice []blockchain.Block

    // Create a response structure
	for _, block := range list {
		emp := &blockchain.Block{Id: block.Id, Nonce: block.Nonce, Timestamp: block.Timestamp, Data: block.Data, PrevHash: block.PrevHash, Hash: block.Hash}
		blockchainBlockchainSlice = append(blockchainBlockchainSlice, *emp)
	}

	blockchainResponse := BlockchainResponse{blockchainBlockchainSlice}

	response, err := json.Marshal(blockchainResponse)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

func main() {
    // Define endpoints
    http.HandleFunc("/newBlock", newBlock)
    http.HandleFunc("/getBlockchain", getBlockchain)

   http.ListenAndServe(":8090", nil)
}
