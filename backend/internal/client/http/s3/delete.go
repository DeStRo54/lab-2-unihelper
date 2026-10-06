package s3

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s3Client S3Client) Delete(ctx context.Context, key string) error {
	if s3Client.client == nil {
		return ErrNotConfigured
	}
	_, err := s3Client.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s3Client.config.Bucketname),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	return nil
}
