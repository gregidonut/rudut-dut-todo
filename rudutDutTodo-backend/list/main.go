package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gregidonut/rudut-dut-todo/rudutDutTodo-backend/list/handler"
)

func main() {
	lambda.Start(handler.Handler)
}
