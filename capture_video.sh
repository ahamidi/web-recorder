#! /usr/bin/env bash

../node_modules/phantomjs2/bin/phantomjs capture_video.js $1 $2 | ffmpeg -r 10 -y -c:v png -f image2pipe -i - -c:v libx264 -pix_fmt yuv420p -movflags +faststart output.mp4
