package server

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"mime/multipart"
	"os"
)

func Oss(FilePath *multipart.FileHeader) {
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New("oss-cn-shanghai.aliyuncs.com", "LTAI5tFghQXd7Vf8T3WA5Y6H", "jdFjKMdbcQ33vTVUkYQkbPSeHPcoMZ")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket("wanan2")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	open, err := FilePath.Open()
	if err != nil {
		return
	}
	all, err := io.ReadAll(open)
	if err != nil {
		return
	}
	// 将Byte数组上传至exampledir目录下的exampleobject.txt文件。
	err = bucket.PutObject("exampledir/exampleobject.txt", bytes.NewReader([]byte(all)))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
