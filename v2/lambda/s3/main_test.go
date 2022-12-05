package main

import (
	"context"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// mockedS3ListResponse represents the mocked
// response we want to receive when testing our function.
type mockedS3ListResponse struct {
}

// This method implements the signature of ListBuckets as defined here: https://github.com/aws/aws-sdk-go-v2/blob/main/service/s3/api_op_ListBuckets.go
func (m mockedS3ListResponse) ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error) {
	name := "mylatestbucket"
	resp := &s3.ListBucketsOutput{
		Buckets: []types.Bucket{
			{
				Name: &name,
			},
		},
	}

	return resp, nil
}
func TestHandleEvent(t *testing.T) {

	// Setup required for our test
	ctx := context.Background()
	expectedBucketName := "mylatestbucket"

	var mockAPI mockedS3ListResponse
	// Call our function under test and pass it our mock object.
	got, err := handleLambdaEvent(ctx, mockAPI)
	if err != nil {
		log.Println(err)
	}

	for _, v := range got.Buckets {
		if *v.Name != expectedBucketName {
			t.Fatalf("expected %v, got %v", expectedBucketName, *v.Name)
		}
	}

}
