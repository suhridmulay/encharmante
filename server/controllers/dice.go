package controllers

import (
	"encharmante/server/services"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

func HandleDiceRoll(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.RequestURI)
	if err != nil {
		HandleError(err, "Validate URL", w)
	}
	queryParams := u.Query()
	sides, err := strconv.Atoi(queryParams.Get("sides"))
	if err != nil {
		HandleError(err, "Enter a valid integer (>0) for sides", w)
		return
	}
	if sides < 0 {
		HandleError(err, "Enter a valid integer (>0) for sides", w)
		return
	}
	bonus, err := strconv.Atoi(queryParams.Get("bonus"))
	if err != nil {
		HandleError(err, "Enter a valid integer (>0) for sides", w)
		return
	}
	diceroll := services.Roll(sides, bonus)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(diceroll)
}
