package main

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type input struct {
	email string `json: "emailAddress"`
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
		}, err
	}
	svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("alg-newsletter"),
		Item: map[string]*dynamodb.AttributeValue{
			"email": &dynamodb.AttributeValue{
				S: &mail.email,
			},
		},
	})
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(Handler)
}
