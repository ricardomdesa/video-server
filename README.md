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
