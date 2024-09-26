package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

func main() {
	lambda.Start(handler)
}

func handler(event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	fmt.Println("Hello, world!")

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body: fmt.Sprintf("[%s] Hello %s",
			os.Getenv("ENV_NAME"),
			event.PathParameters["name"]),
	}, nil
}
