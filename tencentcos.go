package main
import (
   "context"
   "github.com/tencentyun/cos-go-sdk-v5"
   "net/http"
   "net/url"
   "time"
   "path"
   "log"
)
func Tencentcos(c *Conf, file string)  {
	u, _ := url.Parse(c.Tencentcos.BucketUrl)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  c.Tencentcos.SecretId,
			SecretKey: c.Tencentcos.SecretKey,
		},
	})
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	fileName := path.Base(file)
	objectName := year + month + "/" + fileName
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			XCosStorageClass: c.Tencentcos.StorageClass,
		},
	}
	_, err := client.Object.PutFromFile(context.Background(),objectName,file,opt)
	if err != nil {
		log.Fatal(err)
	}
}

