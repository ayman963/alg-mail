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

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var mail input

	switch request.HTTPMethod {
	case http.MethodPost:
		mySession := session.New()
		svc := dynamodb.New(mySession, aws.NewConfig().WithRegion("eu-west-1"))
		err := json.Unmarshal([]byte(request.Body), &mail)
		if err != nil {
			return events.APIGatewayProxyResponse{

				StatusCode: 500,
				Headers: map[string]string{
					"Access-Allow-Control-Origin":  "http://www.ausbildung-leicht-gemacht.de",
					"Access-Control-Allow-Headers": "Accept, Content-Type, application/x-www-form-urlencoded",
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
					"Access-Control-Allow-Headers": "Accept, Content-Type, application/x-www-form-urlencoded",
					"Access-Control-Allow-Methods": "POST"},
			}, err
		}
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers: map[string]string{
				"Access-Allow-Control-Origin": "http://www.ausbildung-leicht-gemacht.de",
			},
		}, nil
	case http.MethodOptions:
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers: map[string]string{
				"Access-Allow-Control-Origin":  "http://www.ausbildung-leicht-gemacht.de",
				"Access-Control-Allow-Headers": "Accept, Content-Type, application/x-www-form-urlencoded",
				"Access-Control-Allow-Methods": "POST",
			},
		}, nil

	}

	return events.APIGatewayProxyResponse{
		StatusCode: 404,
		Headers: map[string]string{
			"Access-Allow-Control-Origin": "http://www.ausbildung-leicht-gemacht.de",
		},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
