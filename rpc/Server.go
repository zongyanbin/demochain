package main

import (
	"demochain/core"
	"encoding/json"
	"io"
	"net/http"
)
var blockchain *core.Blockchain
func run()  {
	http.HandleFunc("/blockchain/get",blockchainGetHandler) // 获取链上数据
	http.HandleFunc("/blockchain/write",blockchainWriteHandler) // 写数据到链上
	http.ListenAndServe("localhost:3333",nil)
}

func blockchainGetHandler(w http.ResponseWriter,r *http.Request)  {
	bytes, error := json.Marshal(blockchain)
	if error != nil{
		http.Error(w,error.Error(),http.StatusInternalServerError)
		return
	}
	io.WriteString(w,string(bytes))
}
func blockchainWriteHandler(w http.ResponseWriter,r*http.Request)  {
	blockData := r.URL.Query().Get("data")
	blockchain.SendData(blockData)
	blockchainGetHandler(w,r)
}

func main()  {
	blockchain = core.NewBlockchain()
	run()
}