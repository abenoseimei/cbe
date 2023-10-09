package main
import (
	"context"
	"io"
	"log"
	"os"
	"time"
	"cloud.google.com/go/storage"
	"path"
)
type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}
var uploader *ClientUploader
func Googlecs(c *Conf, file string) {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", c.Googlecs.Credential)
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	uploader = &ClientUploader{
		cl:         client,
		bucketName: c.Googlecs.BucketName,
		projectID:  c.Googlecs.ProjectID,
		uploadPath: year + month + "/",
	}
	fileName := path.Base(file)
	uploader.UploadFile(file,fileName)
}
func (c *ClientUploader) UploadFile(file string, object string) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object).NewWriter(ctx)
	if _, err := io.Copy(wc, f); err != nil {
		log.Fatal(err)
	}
	if err := wc.Close(); err != nil {
		log.Fatal(err)
	}
	return nil
}
