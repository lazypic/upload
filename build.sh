GOOS=linux GOARCH=amd64 go build -o ./bin/upload upload.go
GOOS=windows GOARCH=amd64 go build -o ./bin/upload.exe upload.go
GOOS=darwin GOARCH=amd64 go build -o ./bin/upload upload.go
