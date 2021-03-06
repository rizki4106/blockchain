package router

import (
	"github.com/gorilla/mux"
	"github.com/rizki4106/blockhain/controller"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/create-block", controller.Mining).Methods("POST")
	r.HandleFunc("/get-block", controller.GetAllBlock).Methods("GET")
	r.HandleFunc("/validate-block", controller.ValidateChain).Methods("GET")
	r.HandleFunc("/add-node", controller.ConnectNode).Methods("POST")
	r.HandleFunc("/verify-chain", controller.CheckTheLongestChain).Methods("GET")
	return r
}