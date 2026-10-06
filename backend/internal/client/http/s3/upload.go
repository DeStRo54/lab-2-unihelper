package s3

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"mime/multipart"
)

func (s3Client S3Client) Upload(ctx context.Context, fileHeader *multipart.FileHeader, keyName string) error {
	if s3Client.client == nil {
		return ErrNotConfigured
	}
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = s3Client.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s3Client.config.Bucketname),
		Key:    aws.String(keyName),
		Body:   file,
	})
	if err != nil {
		return err
	}
	return nil
}
