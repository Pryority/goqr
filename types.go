package main

import (
	"math/rand"
	"time"
)

type TransferRequest struct {
	Recipient int `json:"recipient"`
	Amount    int `json:"amount"`
}

type TransferResponse struct {
	Recipient int       `json:"recipient"`
	Amount    int       `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstName, lastName string, nextID int) *Account {
	return &Account{
		ID:        nextID,
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000)),
		CreatedAt: time.Now().UTC(),
	}
}
