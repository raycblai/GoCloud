package main

import (
    "fmt"
    "github.com/aws/aws-lambda-go/lambda"
    //"github.com/aws/aws-lambda-go/events"
    )
    
func main() {
    lambda.Start(Handler)
}

func Handler(event InputEvent) (string, error){
    // fmt.Println("Function Invoked !!!")
    fmt.Println(event.firstName)
    fmt.Println(event.lastName)
    return "It works !!",nil
}

type InputEvent struct {
    firstName string `json:"firstname"`
    lastName string `json:"lastname"`
}