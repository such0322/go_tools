package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Sessions struct {
	SessionId   string
	AppId       string
	Platform    string
	Guid        string
	Uiid        string
	ExpDate     time.Time
	InsDate     time.Time
	LastModDate time.Time
}
