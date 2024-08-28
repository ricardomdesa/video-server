#!/bin/bash
mkdir -p out
for i in *.mp4;
  do name=`echo "$i" | cut -d'.' -f1`
  mkdir -p out/$name
  echo "$name"
  #ffmpeg -i "$i" -profile:v baseline -level 3.0 -s 1920x1080 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls "out/${name}/index.m3u8"
  ffmpeg -i "$i" -codec: copy -start_number 0 -hls_time 10 -hls_list_size 0 -f hls "out/${name}/index.m3u8"
done
