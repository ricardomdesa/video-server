


Speed up video in 1.5x

```bash
ffmpeg -i 5-Percorrendo\ Arrays.mp4 -filter_complex "[0:v]setpts=0.6667*PTS[v];[0:a]atempo=1.5[a]" -map "[v]" -map "[a]" video5_1_5x.mp4
```

Convert video to hls format
```bash
ffmpeg -i video5_1_5x.mp4 -profile:v baseline -level 3.0 -s 1920x1080 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls 5index.m3u8
```


https://github.com/u2takey/ffmpeg-go