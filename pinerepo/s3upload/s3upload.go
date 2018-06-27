package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	f := flag.String("filename", "", "a string")
	b := flag.String("bucket", "ejunkey", "a string")
	p := flag.String("prefix", "", "a string")
	flag.Parse()
	fmt.Println("filename:", *f)
	fmt.Println("bucket:", *b)
	fmt.Println("prefix:", *p)
	// if len(os.Args) != 4 {
	// 	fmt.Println("usage: %s  <bucket>  <filename>\n", filepath.Base(os.Args[0]))
	// 	os.Exit(1)

	// }
	// bucket := os.Args[1]
	// filename := os.Args[2]

	file, err := os.Open(*f)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	conf := aws.Config{Region: aws.String("us-east-2")}
	sess := session.New(&conf)
	svc := s3manager.NewUploader(sess)

	fmt.Println("Uploading file to S3...")
	result, err := svc.Upload(&s3manager.UploadInput{
		Bucket: aws.String(*b),
		Key:    aws.String(filepath.Base(*f)),
		Body:   file,
	})
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully uploaded %s to %s\n", *f, result.Location)
}

