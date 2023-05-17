package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"udrive-request/configs"
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

	postBody, _ := json.Marshal(map[string]int64{
		// TODO Calcular distância
		"distance": 10,
	})

	responseBody := bytes.NewBuffer(postBody)

	config := configs.GetGateway()
	url := fmt.Sprintf("%s/calculate", config.Url)
	res, err := http.Post(url, "application/json", responseBody)
	fmt.Println("Price calculated")

	if err != nil {
		log.Printf("Error calculating request price: %v", err)
	}
	defer res.Body.Close()
	var response Response
	err = json.NewDecoder(res.Body).Decode(&response)
	if response.Status != 200 || err != nil {
		log.Print("Error calculating request price")
	}

	sql := "INSERT INTO tb_request (user_id, driver_id, origin, destination, scheduled_time, price, status) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"

	err = conn.QueryRow(sql, request.UserId, request.DriverId, request.Origin, request.Destination, request.ScheduledTime, 10, PENDING).Scan(&id)
	fmt.Println("Inserted on database")

	return
}
