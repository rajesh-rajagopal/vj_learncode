package main

import (
	"github.com/minio/minio-go"
	"log"
        "fmt"
)

func main() {
	endpoint := "188.240.231.84:8085"
	accessKeyID := "6YSBQTQHF3HP52062XY5"
	secretAccessKey := "4h5opVTnxqS5JJ4upiMbhqbgeRvBcCow/0qZPe7k"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called mymusic.
	// bucketName := "100"
	//location := "us-east-1"

/*       // Create a done channel to control 'ListObjects' go routine.
doneCh := make(chan struct{})

// Indicate to our routine to exit cleanly upon return.
defer close(doneCh)

isRecursive := true
objectCh := minioClient.ListObjects(bucketName, "", isRecursive, doneCh)
for object := range objectCh {
    if object.Err != nil {
        fmt.Println(object.Err)
        return
    }
    fmt.Println(object)
}
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
