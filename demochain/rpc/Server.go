package rpc

import (
	"net/http"
	"github.com/ghost0036/go-hello/demochain/core"
	"encoding/json"
	"io"
	"fmt"
)

var blockchain *core.Blockchain

func Run(bc *core.Blockchain) {
	blockchain = bc
	http.HandleFunc("/blockchain/get",blockchainGetHandler)
	http.HandleFunc("/blockchain/write",blockchainWriterHandler)
	fmt.Println("Server listening on port: 8888...")
	http.ListenAndServe(":8888",nil)
}

func blockchainGetHandler(w http.ResponseWriter,r *http.Request)  {
	bytes,err := json.Marshal(blockchain)
	if err != nil {
		http.Error(w,"",http.StatusInternalServerError)
		return
	}

	io.WriteString(w,string(bytes))
}

func blockchainWriterHandler(w http.ResponseWriter,r *http.Request)  {
	blockchain.Print()
	vars := r.URL.Query()
	blockchain.SendData(vars["data"][0])
	blockchain.Print()
	io.WriteString(w,"Sucess!!")
}