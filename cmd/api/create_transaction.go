package main

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gustvision/backend-interview/pkg/account"

	"github.com/rs/zerolog/log"
)

func (h *handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	logger := log.With().Str("method", "create_transaction").Logger()

	var t account.Transaction

	//TODO: improve the validation
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil || t.ID == "" || t.AccountID == "" || t.Amount < 0 || t.CreatedAt < 0 {
		logger.Error().Err(err).Msg("invalid payload")
		http.Error(w, "invalid payload", http.StatusBadRequest)

		return
	}

	// # Verify if account exists
	a, err := h.account.Fetch(ctx, account.Filter{ID: t.AccountID})
	fmt.Println(a)
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch account")
		http.Error(w, "failed to fetch account", http.StatusInternalServerError)

		return
	}

	err = h.account.InsertTransaction(ctx, t)

	if err != nil {
		logger.Error().Err(err).Msg("failed to create transaction")
		http.Error(w, "failed to fetch account", http.StatusInternalServerError)

		return
	}

	// #Update the account total
	fmt.Println(a.Total)
	fmt.Println(t.Amount)
	newTotal := a.Total - t.Amount
	err = h.account.UpdateAccountTotal(ctx, account.Filter{ID: t.AccountID, Total: newTotal})
	if err != nil {
		logger.Error().Err(err).Msg("failed to update acount")
		http.Error(w, "failed to update account", http.StatusInternalServerError)

		return
	}

	// #Write response
	w.WriteHeader(http.StatusCreated)
	logger.Info().Msg("success")
}
