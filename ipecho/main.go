package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler)
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received headers: ", request.Headers)

	parts := strings.Split(request.Headers["X-Forwarded-For"], ",")

	return events.APIGatewayProxyResponse{Body: parts[0], StatusCode: 200}, nil
}
