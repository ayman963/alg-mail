package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	request := events.APIGatewayProxyRequest{}
	request.Body = "{ \"emailAddress\": \"ayman.hisnawi@gmail.com\"}"

	_, err := Handler(request)

	// assert.Equal(t, response.Headers, expectedResponse.Headers)
	// assert.Contains(t, response.Body, expectedResponse.Body)
	assert.Equal(t, err, nil)

}
