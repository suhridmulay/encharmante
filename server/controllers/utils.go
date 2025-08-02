package controllers

import (
	"encoding/json"
	"net/http"
)

type ErrorDescriptor struct {
	Reason   string `json:"reason"`
	Recovery string `json:"recovery"`
}

func HandleError(err error, recovery string, w http.ResponseWriter) {
	descriptor := ErrorDescriptor{
		Reason:   err.Error(),
		Recovery: recovery,
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(descriptor)
}
