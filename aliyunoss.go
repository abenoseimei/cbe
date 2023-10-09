package main
import (
   "github.com/aliyun/aliyun-oss-go-sdk/oss"
   "time"
   "path"
   "log"
)
func Aliyunoss(c *Conf, file string) {
	endpoint := c.Aliyunoss.Endpoint
	accessKeyId := c.Aliyunoss.AccessKeyId
	accessKeySecret := c.Aliyunoss.AccessKeySecret
	bucketName := c.Aliyunoss.BucketName
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	localFileName := file
	fileName := path.Base(localFileName)
	objectName := year + month + "/" + fileName
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		log.Fatal(err)
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		log.Fatal(err)
	}
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		log.Fatal(err)
	}
}
