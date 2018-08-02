package main

import (
	"flag"
	"fmt"
)

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
	fmt.Println(*projectPtr)
	fmt.Println(*episodePtr)
	fmt.Println(*scenePtr)
	fmt.Println(*cutPtr)
	fmt.Println(*filePtr)
	fmt.Println(flag.Args())
}
