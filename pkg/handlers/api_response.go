package handlers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func apiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	stringBody, _ := json.Marshal(body)

	resp := events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "apllication/json"},
		StatusCode: status,
		Body:       string(stringBody),
	}

	return &resp, nil
}
