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
	fmt.Println(*regionPtr)
	fmt.Println(*bucketPtr)
	fmt.Println(*filePtr)
	key := fmt.Sprintf("%s/%d/%d_%d%s", *projectPtr, *episodePtr, *scenePtr, *cutPtr, ".blend")
	fmt.Println(flag.Args())
	fmt.Println(key)
	s := session.New(&aws.Config{Region: aws.String(*regionPtr)})
	err := uploadS3(s, *bucketPtr, key, *filePtr)
	if err != nil {
		log.Fatal(err)
	}
}
