package model

type Block struct {
	Index int `json:"index"`
	Proof int `json:"proof"`
	Timestamp string `json:"timestamp"`
	Prevhash string `json:"previous_hash"`
	Data []Data `json:"data"`
}

type Nodes struct {
	URL string `json:"url"`
}

type Data struct {

}