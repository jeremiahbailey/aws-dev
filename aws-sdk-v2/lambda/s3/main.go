package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// s3ListBucketAPI satisfies the ListBuckets method for the s3 service.
type s3ListBucketsAPI interface {
	ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
}

// main() is the entrypoint we must provide for the lambda service.
func main() {

	lambda.Start(handleLambdaEvent)
}

// We provide our s3ListBucketAPI interface so that we are not passing a concrete type which
// allows us to perform mocks.
func handleLambdaEvent(ctx context.Context, api s3ListBucketsAPI) (*s3.ListBucketsOutput, error) {
	resp, err := api.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		log.Print(err)
	}
	return resp, nil
}
