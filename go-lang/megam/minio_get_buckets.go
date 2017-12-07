package main

import (
	"github.com/minio/minio-go"
	"log"
"fmt"
)

func main() {
	endpoint := "apidb.megam.io:8000"
	accessKeyID := "BR902BPU65FXCDRW6IL7"
	secretAccessKey := "pnw9Ko1saN+XH1OSxDWGIkpsmWpf5DPnl7gFtVDV"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

bucketName := "backup"
fmt.Println(minioClient)
exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}


/* to upload a file
	objectName := "abcd.pdf"
	filePath := "/home/vijay/Desktop/abcd.pdf"
	contentType := "application/pdf"
n, err := minioClient.FPutObject(bucketName, objectName, filePath, contentType)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
*/

buckets, err := minioClient.ListBuckets()
    if err != nil {
    fmt.Println(err)
    return
}
for _, bucket := range buckets {
    fmt.Println(bucket)
}



}

