package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)
//"os"

func handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	/*var (
		apiUrl = os.Getenv("http://ae9cc1af00cdb488ea524a1da64bf434-730275616.us-east-2.elb.amazonaws.com:3000")
	)*/

	//apiUrl = os.Getenv("<API_URL>")
	var apiUrl = "http://ae9cc1af00cdb488ea524a1da64bf434-730275616.us-east-2.elb.amazonaws.com:3000"

	cpf := event.PathParameters["cpf"]
	if cpf == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Cpf invalido",
		}, nil
	}

	res, err := http.Get(fmt.Sprintf("%s/login/%s", apiUrl, cpf))
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, err
	}
	token := res.Header.Get("Authorization")

	mapRes := map[string]string{"token": token}
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
			"Authorization": token,
		},
		Body: string(body),
	}, nil
}

func main() {
	lambda.Start(handler)
}
