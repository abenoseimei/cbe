package main
import (
	"flag"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)
type Conf struct {
	Aliyunoss struct {
		Endpoint		string	`yaml:"endpoint"`
		AccessKeyId		string	`yaml:"accessKeyId"`
		AccessKeySecret string	`yaml:"accessKeySecret"`
		BucketName		string	`yaml:"bucketName"`
	}
	Tencentcos struct {
		BucketUrl		string	`yaml:"bucketUrl"`
		SecretId		string	`yaml:"secretId"`
		SecretKey		string	`yaml:"secretKey"`
		StorageClass	string	`yaml:"storageClass"`
	}
	Googlecs struct {
		ProjectID		string	`yaml:"projectID"`
		BucketName		string	`yaml:"bucketName"`
		Credential		string	`yaml:"credential"`
	}
	Googledrive struct {
		Credential		string	`yaml:"credential"`
		Parent			string	`yaml:"parent"`
	}
}
func readConf(file string) *Conf {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	var conf Conf
	err = yaml.Unmarshal(buf, &conf)
	if err != nil {
		log.Fatal(err)
	}
	return &conf
}
func main() {
	var mode		string
	var confFile	string
	var localFile	string 
	flag.StringVar(&mode, "m", "", "")
	flag.StringVar(&confFile, "c", "cbe.yaml", "")
	flag.StringVar(&localFile, "f", "", "")
	flag.Parse()
	if localFile == "" && mode != "googledrivelist" {
		log.Fatal("No local file specified")
	}
	conf := readConf(confFile)
	if mode == "aliyunoss" {
		Aliyunoss(conf, localFile)
	} else if mode == "tencentcos" {
		Tencentcos(conf, localFile)
	} else if mode == "googlecs" {
		Googlecs(conf, localFile)
	} else if mode == "googledrive" {
		Googledrive(conf, localFile, "")
	} else if mode == "googledrivelist" {
		Googledrive(conf, localFile, "list")
	} else {
		log.Fatal("Unsupported mode: ", mode)
	}
}
