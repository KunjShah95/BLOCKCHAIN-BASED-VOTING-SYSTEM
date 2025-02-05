package handlers

import (
	"encoding/json"
	"net/http"
	"voting-system/blockchain"
)

var bc = blockchain.NewBlockchain()

func VoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var vote blockchain.Vote
		if err := json.NewDecoder(r.Body).Decode(&vote); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		bc.AddVote(vote.Candidate)
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(vote)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}