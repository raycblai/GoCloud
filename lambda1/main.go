package main

import (
	"fmt"
	
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	FName string `json:"What is your name?"`
	City string `json:"Where do you live?"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func hello(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("Hello %s, I know you live in %s !!", event.FName, event.City)}, nil
}

func main() {
	lambda.Start(hello)
}
