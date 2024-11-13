package utils

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/minio/minio-go"
)

type bucketclient struct {
	client      *minio.Client
	objectName  string
	filePath    string
	contentType string
}

//封装client
func NewClient() *bucketclient {

	endpoint := "43.138.181.174:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		fmt.Println(err)
	}

	buckets, err := minioClient.ListBuckets()
	if err != nil {
		fmt.Println(err)
	}
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}
	return &bucketclient{
		client: minioClient,
	}
}

//封装ObjectName
func (b *bucketclient) GetObjectName() string {
	return b.objectName
}
func (b *bucketclient) SetObjectName(objectName string) {
	b.objectName = objectName
}

//封装filePath
func (b *bucketclient) GetfilePath() string {
	return b.filePath
}
func (b *bucketclient) SetfilePath(filePath string) {
	b.filePath = filePath
}

//封装contentType
func (b *bucketclient) GetcontentType() string {
	return b.contentType
}
func (b *bucketclient) SetcontentType(contentType string) {
	b.contentType = contentType
}

//上传文件夹内部文件
//UpLoadFile 将整个目录都上传,遍历所有文件夹及子文件夹上传
func (b *bucketclient) UpLoadFile(filepath string, bucketName string) {
	file, err := ioutil.ReadDir(filepath)
	// var bc bucketclient
	if err != nil {
		log.Fatalln(err)
	}
	for _, fi := range file {
		if fi.Name() == ".vscode" {
			continue

		}
		if !fi.IsDir() {
			stirngpath := filepath + "/" + fi.Name()
			str := stirngpath[3:]
			n, err := b.client.FPutObject(bucketName, str, stirngpath, minio.PutObjectOptions{
				ContentType: b.contentType,
			})
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Successfully uploaded bytes: ", n)
		} else if fi.IsDir() {
			b.UpLoadFile(filepath+"/"+fi.Name(), bucketName)

		}

	}
}

//下载桶文件到本地
func (b *bucketclient) DownloadFile(bucketName string, path string) {
	err := b.client.FGetObject(bucketName, b.objectName, path, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}

}

func Bucket() {
	b := NewClient()
	//创建一个桶
	bucketName := "flt-file"
	location := "us-east-1"
	err1 := b.client.MakeBucket(bucketName, location)
	//检测桶是否存在
	if err1 != nil {
		isIn, err := b.client.BucketExists(bucketName)
		if err == nil && isIn {
			log.Println("桶已经创建成功")
		} else {
			log.Fatalln(err)
		}
	}
	//创建完桶之后进行上传
	//上传文件
	// b.SetObjectName("k8sfile")
	b.SetcontentType("application/txt")
	b.UpLoadFile("./BackupFile", bucketName)
}
func Download() {
	b := NewClient()
	//创建一个桶
	path := "/home/xin/minio/g"
	bucketName := "k8s-file2"
	b.objectName = "GORUN/b_pv/pv.txt"
	b.DownloadFile(bucketName, path)

}
