package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func listAllKeysinBucket(bucket string) {
	svc := s3.New(session.New())

	params := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	}

	resp, _ := svc.ListObjects(params)
	for _, key := range resp.Contents {
		fmt.Println(*key.Key)
	}
}
