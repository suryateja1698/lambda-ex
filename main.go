package main

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Even struct {
	Number *int `json:"number"`
}

type EvenResponse struct {
	IsEven bool `json:"is_even"`
}

func main() {
	lambda.Start(IsEven)
}

func IsEven(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.Body == "" {
		return events.APIGatewayProxyResponse{}, errors.New("empty body")
	}

	var body Even
	err := json.Unmarshal([]byte(req.Body), &body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	var status int
	var check bool

	switch *body.Number%2 == 0 {
	case true:
		status = 200
		check = true
	default:
		status = 400
	}

	resp := EvenResponse{
		IsEven: check,
	}

	byteRes, err := json.Marshal(resp)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(byteRes),
	}, nil
}
