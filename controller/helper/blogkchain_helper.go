package helper

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rizki4106/blockhain/model"
)

/* It's a function to create a block. */
func CreateBlock(index int, proof int, prev_hash string) model.Block {
	block := model.Block{Index: index, Proof: proof, Timestamp: time.Now().String(), Prevhash: prev_hash}
	return block
}

/* It's a function to get the last block of the blockchain. */
func GetLastBlock(block []model.Block) model.Block {
	return block[len(block) - 1]
}

/* It's a function to calculate the proof of work. */
func ProofOfWork(prev_proof int) int {
	new_proof := 1
	valid := false

	for !valid {
		calculate := (new_proof * 2) - (prev_proof * 2)
		validation := fmt.Sprintf("%x", sha256.Sum256([]byte(string(calculate))))
		if validation[:4] == "0000" {
			valid = true
		}else{
			new_proof = new_proof + 1
		}
	}

	return new_proof
}

/* It's a function to hash the block. */
func Hash(block model.Block) string {
	toJson, _ := json.Marshal(block)
	toSha256 := sha256.Sum256(toJson)
	toString := fmt.Sprintf("%x", toSha256)
	return toString
}

/* It's a function to check the blockchain is valid or not.
*/
func IsChainValid(chain []model.Block) bool {
	init_block := chain[0]
	init_index := 1

	for init_index < len(chain) {
		block := chain[init_index]

		if block.Prevhash != Hash(init_block) {
			return false
		}

		prev_proof := init_block.Proof
		proof := block.Proof
		calculate := (proof * 2) - (prev_proof * 2)
		validation := fmt.Sprintf("%x", sha256.Sum256([]byte(string(calculate))))

		if validation[:4] != "0000" {
			return false
		}

		init_index = init_index + 1
		init_block = block
	}

	return true
}