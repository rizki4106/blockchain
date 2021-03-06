package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rizki4106/blockhain/controller/helper"
	"github.com/rizki4106/blockhain/model"
)

var blockhain []model.Block
var nodes []model.Nodes

type Response struct {
	TotalItem int `json:"total_item"`
	Message string `json:"message"`
	Block []model.Block `json:"block"`
}
/* This function is used to create a genesis block. */
func init(){
	genesis_block := model.Block{Index: 1, Proof: 1, Prevhash: "0", Timestamp: time.Now().String()}
	blockhain = append(blockhain, genesis_block)
}

/* This function is used to create a new block. */
func Mining(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var mined_block []model.Block

	prev_block := helper.GetLastBlock(blockhain)
	prev_proof := prev_block.Proof

	proof_of_work := helper.ProofOfWork(prev_proof)
	prev_hash := helper.Hash(prev_block)
	new_block := helper.CreateBlock(len(blockhain) + 1, proof_of_work, prev_hash)

	blockhain = append(blockhain, new_block)
	mined_block = append(mined_block, new_block)

	response := Response{len(blockhain), "New block has been created", mined_block}
	json.NewEncoder(w).Encode(response)
}

/* This function is used to get all block in blockhain. */
func GetAllBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{len(blockhain), "", blockhain}
	json.NewEncoder(w).Encode(response)
}

/* This function is used to validate the blockchain. */
func ValidateChain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response []byte
	valid := helper.IsChainValid(blockhain)

	if valid {
		response = []byte(`{"message": "All blockchain looks good"}`)
	}else{
		response = []byte(`{"message": "Blockhain not valid"}`)
	}
	w.Write(response)
}

func ConnectNode(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var node model.Nodes
	json.NewDecoder(r.Body).Decode(&node)
	nodes = append(nodes, node)

	json.NewEncoder(w).Encode([]map[string]string{{"message": "Success"}})
}

func CheckTheLongestChain(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	
	for _, node := range nodes {
		var response Response
		fmt.Println(node.URL)
		req, _ := http.Get(node.URL + "/get-block")
		res, _ := ioutil.ReadAll(req.Body)
		json.Unmarshal(res, &response)

		if len(response.Block) > len(blockhain) {
			blockhain = response.Block
		}
	}

	json.NewEncoder(w).Encode(blockhain)
}