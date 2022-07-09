#!/bin/bash
echo "building..."
GOARCH=amd64 GOOS=linux go build main.go

echo "compressing..."
file_name=v
current_time=$(date "+%Y.%m.%d-%H.%M.%S")
new_fileName=$file_name.$current_time
zip $new_fileName.zip main
mv $new_fileName.zip ./build
echo "done"