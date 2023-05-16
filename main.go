package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"udrive-request/handler"
)

func main() {
	lambda.Start(handler.Create)
}
