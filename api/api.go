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