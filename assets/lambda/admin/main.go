package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type myEvent struct {
	Path string `json:"path"`
}

func handler(ctx context.Context, e myEvent) (interface{}, error) {
	return fmt.Sprintf("someone asked for %s on the admin side!", e.Path), nil
}

func main() {
	lambda.Start(handler)
}
