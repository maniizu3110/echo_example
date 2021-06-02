package middlewares

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/storage"
)
//開発環境では一回切っておく
//GcpConfig Gcpのcledentialsを設定する
func GcpConfig() {
	ctx := context.Background()
	
	projectID := "next-echo-app"

	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Sets the name for the new bucket.
	bucketName := "my-new-bucket"

	// Creates a Bucket instance.
	bucket := client.Bucket(bucketName)

	// Creates the new bucket.
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	if err := bucket.Create(ctx, projectID, nil); err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
	}

	fmt.Printf("Bucket %v created.\n", bucketName)
}
//TODO:リクエストを数回繰り返して、無駄にバケットが作られたりしないか確認