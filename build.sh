# OS별로 빌드함.
GOOS=linux GOARCH=amd64 go build -o ./bin/linux/upload upload.go
GOOS=windows GOARCH=amd64 go build -o ./bin/windows/upload.exe upload.go
GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin/upload upload.go

# Github Release에 업로드 하기위해 압축
cd ./bin/linux/ && tar -zcvf ../upload_linux.tgz . && cd -
cd ./bin/windows/ && tar -zcvf ../upload_windows.tgz . && cd -
cd ./bin/darwin/ && tar -zcvf ../upload_darwin.tgz . && cd -
