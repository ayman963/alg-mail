package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/sirupsen/logrus"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type input struct {
	Email string `json:"emailAddress"`
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var mail input
	mySession := session.New()
	svc := dynamodb.New(mySession, aws.NewConfig().WithRegion("eu-west-1"))
	err := json.Unmarshal([]byte(request.Body), &mail)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Access-Allow-Control-Origin":  "http://www.ausbildung-leicht-gemacht.de",
				"Access-Control-Allow-Headers": "application/x-www-form-urlencoded",
				"Access-Control-Allow-Methods": "POST",
			},
		}, err
	}
	_, err = svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("alg-newsletter"),
		Item: map[string]*dynamodb.AttributeValue{
			"email": &dynamodb.AttributeValue{
				S: &mail.Email,
			},
		},
	})
	if err != nil {
		logrus.Error(err.Error())
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Access-Allow-Control-Origin":  "http://www.ausbildung-leicht-gemacht.de",
				"Access-Control-Allow-Headers": "application/x-www-form-urlencoded",
				"Access-Control-Allow-Methods": "POST"},
		}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Allow-Control-Origin":  "http://www.ausbildung-leicht-gemacht.de",
			"Access-Control-Allow-Headers": "application/x-www-form-urlencoded",
			"Access-Control-Allow-Methods": "POST"},
	}, nil

}

func main() {
	lambda.Start(Handler)
}
