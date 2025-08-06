package main

import (
	"encoding/json"
	"net/http"
)

type StripePayload struct {
	Currency string `json:"currency"`
	Amount   int64  `json:"amount"`
}

type JsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
	Content string `json:"content"`
	ID      int    `json:"id"`
}

func (app *application) getPaymentIntent(w http.ResponseWriter, r *http.Request) {
	j := JsonResponse{
		OK: true,
	}

	out, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		app.errorLog.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
