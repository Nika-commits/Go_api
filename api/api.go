package api

import (
	"encoding/json"
	"net/http"
)

type BalanceParams struct {
	Username string
}

type BalanceResponse struct {
	Code int
	Balance float64
}

type Error struct {
	Code int
	Message string
}

func writeError(w http.ResponseWriter, messsage string, code int){
	resp := Error{
		Code: code,
		Message: messsage,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}