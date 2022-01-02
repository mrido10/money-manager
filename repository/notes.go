package repository

import (
	"time"
)

type Notes struct {
	NotesID         *int64    `json:"notesID"`
	UserId          string    `json:"userID"`
	WalletID        string    `json:"walletID"`
	NotesTypeID     string    `json:"notesTypeID"`
	NotesCategoryID string    `json:"notesCategoryID"`
	NotesName       string    `json:"notesName"`
	Amount          int64     `json:"amount"`
	FormAccount     string    `json:"fromAccount"`
	ToAccount       string    `json:"toAccount"`
	Description     string    `json:"description"`
	NotesDate       time.Time `json:"notesDate"`
}
