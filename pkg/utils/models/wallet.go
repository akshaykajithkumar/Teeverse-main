package models

import "Teeverse/pkg/domain"

type Wallet struct {
	Balance int                    `json:"balance"`
	History []domain.WalletHistory `json:"history"`
}
