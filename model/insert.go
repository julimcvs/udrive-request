package model

import (
	"udrive-request/db"
)

type Response struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

func Insert(request Request) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := "INSERT INTO tb_request (user_id, origin, destination, scheduled_time, price, status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"

	err = conn.QueryRow(sql, request.UserId, request.Origin, request.Destination, request.ScheduledTime, request.Price, PENDING).Scan(&id)

	return
}
