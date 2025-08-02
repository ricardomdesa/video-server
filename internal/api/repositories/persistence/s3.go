package persistence

import (
	context "context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Repository struct {
	session *session.Session
	bucket  string
}

func NewS3Repository(sess *session.Session, bucket string) *S3Repository {
	return &S3Repository{
		session: sess,
		bucket:  bucket,
	}
}

func (s *S3Repository) GetMediaData(ctx context.Context, key string) ([]byte, error) {
	// Implement the logic to retrieve media data from S3 using the session and bucket.
	// This is a placeholder for the actual implementation.
	downloader := s3manager.NewDownloader(s.session)
	buf := &aws.WriteAtBuffer{}
	_, err := downloader.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}	