package component

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

var TencentCOSClent *cos.Client

// InitBucket 初始化微信配置
func InitBucket(bucketURL string, secretID string, secretKey string) {
	u, _ := url.Parse(bucketURL)
	b := &cos.BaseURL{BucketURL: u}

	TencentCOSClent = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretID,
			SecretKey: secretKey,
		},
	})
}

// PutObject  上传文件
func PutObject(path string, r io.Reader) error {
	_, err := TencentCOSClent.Object.Put(context.Background(), path, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// ListObject 对象列表
func ListObject() {
	opt := &cos.BucketGetOptions{
		Prefix:  "",
		MaxKeys: 3,
	}
	v, _, err := TencentCOSClent.Bucket.Get(context.Background(), opt)
	if err != nil {
		fmt.Println("ere:", err)
		panic(err)
	}
	for _, c := range v.Contents {
		fmt.Printf("%s, %d\n", c.Key, c.Size)
	}
}
