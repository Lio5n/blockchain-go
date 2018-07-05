package main

import (
	"net/http"
	"core"
	"encoding/json"
	"io"
)

var blockchain *core.Blockchain

func run()  {
	http.HandleFunc("/blockchain/get",blockchainGetHandler)
	http.HandleFunc("/blockchain/write",blockchainWriteHandler)
	http.ListenAndServe("localhost:8888",nil)
}

func blockchainGetHandler(w http.ResponseWriter,r *http.Request)  {
	bytes, error := json.Marshal(blockchain)
	if error != nil {
		http.Error(w,error.Error(),http.StatusInternalServerError)
		return
	}
	io.WriteString(w,string(bytes))
}

func blockchainWriteHandler(w http.ResponseWriter,r *http.Request)  {
	blockData := r.URL.Query().Get("data")
	blockchain.SendData(blockData)
	blockchainGetHandler(w,r)
}

func main()  {
	// http://localhost:8888/blockchain/get
	// http://localhost:8888/blockchain/write?data=send 1 btc to jack
	blockchain = core.NewBlockchain()
	run()
}