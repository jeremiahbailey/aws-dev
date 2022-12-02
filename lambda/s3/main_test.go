package main

import (
	"context"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// mockedS3ListResponse represents the mocked
// response we want to receive when testing our function.
type mockedS3ListResponse struct {
	s3iface.S3API
	Resp s3.ListBucketsOutput
}

// This method implements the signature of ListBuckets as defined in the s3iface package.
// This satisfies the s3API interface defined in mockedS3ListResponse.
func (m mockedS3ListResponse) ListBuckets(*s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	return &m.Resp, nil
}
func TestHandleEvent(t *testing.T) {

	// Setup required for our test
	ctx := context.Background()
	expectedBucketName := "mylatestbucket"
	bucket := &s3.Bucket{
		Name: &expectedBucketName,
	}

	// Instantiate our mock type and populate the response field
	// with the ListBucketsOutput we expect.
	m := &mockedS3ListResponse{
		Resp: s3.ListBucketsOutput{Buckets: []*s3.Bucket{bucket}},
	}

	// Instantiate our s3Query type and provide our mock
	// object ('m') to satisfy the s3iface.S3API type.
	q := &s3Query{
		s3:     m,
		params: &s3.ListBucketsInput{},
	}

	// Call our function under test. Note
	// the use of 'q.' syntax this means that
	// we are passing our mock object that is in 'q'.
	got, err := q.handleLambdaEvent(ctx)
	if err != nil {
		log.Println(err)
	}

	for _, v := range got.Buckets {
		if v.Name != &expectedBucketName {
			t.Fatalf("expected %v , got %v", expectedBucketName, got)
		}
	}

}
