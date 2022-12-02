package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// s3Query is our package's type which we are using
// to access dependencies in the handler.
type s3Query struct {
	s3     s3iface.S3API
	params *s3.ListBucketsInput
}

// main() is the entrypoint we must provide for the lambda service.
// In main we create our 'session' and instantiate our type with the
// values required when invoking our handleLambdaEvent.
func main() {
	session := session.Must(session.NewSession())

	q := &s3Query{
		s3:     s3.New(session),
		params: &s3.ListBucketsInput{},
	}

	// Since our handleLambdaEvent func
	// uses pointer receiver semantics,
	// we call it using q. syntax where
	// 'q' is a variable of type s3Query
	lambda.Start(q.handleLambdaEvent)
}

// Because there are only 7 valid signatures for a Lambda handler Golang func, we use
// a pointer receiver function to pass in the client and query parameters (e.g. the dependencies).
func (q *s3Query) handleLambdaEvent(ctx context.Context) (*s3.ListBucketsOutput, error) {

	resp, err := q.s3.ListBuckets(q.params)
	if err != nil {
		log.Print(err)
	}

	return resp, nil
}
