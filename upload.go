package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func uploadS3(s *session.Session, bucket, key, fileDir string) error {
	file, err := os.Open(fileDir)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(key),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return err
}

func main() {
	regionPtr := flag.String("region", "ap-northeast-2", "AWS region name.") // 현재 Lazypic은 서울리전을 사용중이다.
	bucketPtr := flag.String("bucket", "lazypic", "S3 bucket name.")
	projectPtr := flag.String("project", "", "project name")
	episodePtr := flag.Int("episode", 0, "episode number")
	scenePtr := flag.Int("scene", 0, "scene number")
	cutPtr := flag.Int("cut", 0, "cut number")
	filePtr := flag.String("file", "", "project name")
	flag.Parse()
	if *projectPtr == "" || *filePtr == "" {
		fmt.Println("Upload file for AWS S3")
		fmt.Println("Copyright 2018, Lazypic, All rights reserved.")
		flag.PrintDefaults()
		return
	}
	path, err := filepath.Abs(*filePtr)
	if err != nil {
		log.Fatal(err)
	}
	ext := filepath.Ext(path)
	key := fmt.Sprintf("%s/%d/%d_%d%s", *projectPtr, *episodePtr, *scenePtr, *cutPtr, ext)
	s := session.New(&aws.Config{Region: aws.String(*regionPtr)})
	err = uploadS3(s, *bucketPtr, key, path)
	if err != nil {
		log.Fatal(err)
	}
}
