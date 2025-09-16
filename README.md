# Video streaming server

[link video streaming-server](https://www.rohitmundra.com/video-streaming-server)

# Startup server
Run
```bash
go build cmd/api/main.go && ./main
```

The videos must be in assets/media folder and converted to m3u8 HLS format

## Speed up video in 1.5x.

```bash
ffmpeg -i 5-video.mp4 -filter_complex "[0:v]setpts=0.6667*PTS[v];[0:a]atempo=1.5[a]" -map "[v]" -map "[a]" video5_1_5x.mp4
```

## Convert video to hls format
```bash
ffmpeg -i filename.mp4 -codec: copy -start_number 0 -hls_time 10 -hls_list_size 0 -f hls filename.m3u8
```

---

# Kubernetes - minikube

## create namespace

```bash

kubectl create namespace video-str

```


## secret

```bash
kubectl create secret <secret_type> <secret_name> -n <namespace_name>

kubectl get secrets -n <namespace>

```

secret.yaml
kubectl apply -f secret.yaml
```

apiVersion: v1
kind: Secret
metadata:
  name: secret-env
type: Opaque
data:
  secret_key: <base64-encoded-secret-value>

```

---

# Localstack

S3 Bucket is running in localstack

~/projects/localstack

```bash

localstack start

aws --endpoint-url=http://localhost:4566 s3 mb s3://videos-bucket-fullc --profile localstack;

aws --endpoint-url=http://localhost:4566 s3 ls s3://videos-bucket-fullc --profile localstack;


```

## Creating bucket and feeding

Use the Makefile commands to create the bucket
`make createbucket`

Check if the bucket was created:
`make lsbuckets`

---

### Feeding

Run the upload_videos_s3 binary

```bash
cd cmd/upload_videos_s3
go run main.go
```



