package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"udrive-request/configs"
	"udrive-request/handler"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	lambda.Start(handler.Create)
}
