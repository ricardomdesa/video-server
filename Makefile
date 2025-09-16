.PHONY: all clean uploadVideos listVideos test downloadVideos

lsvideos:
	@echo "Listing videos in S3 bucket..."
	@aws s3 ls s3://videos-bucket-fullc --recursive --summarize --human-readable --profile localstack

createbucket:
	@aws s3 mb s3://videos-bucket-fullc --profile localstack

deletebucket:
	@aws s3 rb s3://videos-bucket-fullc --force --profile localstack

lsbucket:
	@aws s3 ls --profile localstack

test:
	go test -v ./...

runapi:
	@echo "Running API..."
	go run ./cmd/api/main.go
