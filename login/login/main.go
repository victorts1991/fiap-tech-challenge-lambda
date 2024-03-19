package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

//"os"

type LoginBody struct {
	Cpf    string `json:"cpf"`
	ApiUrl string `json:"api_url"`
}

func handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var login LoginBody
	err := json.Unmarshal([]byte(event.Body), &login)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "body invalid",
		}, err
	}

	if &login == nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "body invalid",
		}, err
	}

	cpf := login.Cpf
	if cpf == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Cpf invalido",
		}, nil
	}

	res, err := http.Get(fmt.Sprintf("%s/login/%s", login.ApiUrl, cpf))
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
