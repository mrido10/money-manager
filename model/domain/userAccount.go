package domain

import "time"

type UserAccount struct {
	UserID   string    `json:"userID"`
	UserName string    `json:"userName"`
	JoinDate time.Time `json:"joinDate"`
}
