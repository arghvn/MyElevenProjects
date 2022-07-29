package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {}

func main() {
	lambda.Start(HandleLambdaEvent)
}
