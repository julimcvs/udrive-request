package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type Request struct {
	ID            *int64           `json:"id"`
	UserId        int64            `json:"userId"`
	DriverId      int64            `json:"driverId"`
	Origin        string           `json:"origin"`
	Destination   string           `json:"destination"`
	Distance      int64            `json:"distance"`
	ScheduledTime time.Time        `json:"time"`
	Price         *decimal.Decimal `json:"price"`
	Status        *Status          `json:"status"`
}

type Status string

type ResponseBody struct {
	Status  *int    `json:"status"`
	Message *string `json:"message"`
}

const (
	PENDING  Status = "PENDING"
	ACCPETED Status = "ACCEPTED"
	FINISHED Status = "FINISHED"
)
