package handler

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"udrive-request/model"
)

func Create(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Request:", r.Body)
	var request model.Request

	err := json.Unmarshal([]byte(r.Body), &request)

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, fmt.Errorf("method arguments not valid, %v", err)
	}

	id, err := model.Insert(request)

	var msg string
	var status int

	if err != nil {
		msg = fmt.Sprintf("Error creating ride with parameters: %v", err)
		status = 400
	} else {
		msg = fmt.Sprintf("Ride successfully created! ID: %v", id)
		status = 201
	}

	responseBody := model.ResponseBody{
		Message: &msg,
	}

	jbytes, err := json.Marshal(responseBody)

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(jbytes),
	}

	return response, nil
}
