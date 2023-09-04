package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func main() {
	lambda.Start(handler)
}
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Hola Juan Valero! ")
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
	}

	return response, nil
}
