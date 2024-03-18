package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	timeout = time.Millisecond * 5
)



func handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var api_url string = "<API_URL>"

	log.Println("authenticating user2 - ")
	randNum := rand.Int63()

	mapRes := map[string]int64{"random_number - <API_URL>": randNum}
	body, err := json.Marshal(mapRes)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "coundnt marshal the body",
		}, err
	}


	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "test-123",
		},
		Body: string(body),
	}, nil
}

func main() {
	lambda.Start(handler)
}
