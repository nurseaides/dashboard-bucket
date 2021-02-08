package component

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

var tencentCOSClent *cos.Client

// initTencentCOS 初始化微信配置
func initTencentCOS(bucketURL string, secretID string, secretKey string) {
	u, _ := url.Parse(bucketURL)
	b := &cos.BaseURL{BucketURL: u}
	tencentCOSClent = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretID,
			SecretKey: secretKey,
		},
	})
}

// tencentCOSPutObject  上传文件
func tencentCOSPutObject(path string, r io.Reader) error {
	_, err := tencentCOSClent.Object.Put(context.Background(), path, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// tencentCOSListObject 对象列表
func tencentCOSListObject() {

	opt := &cos.BucketGetOptions{
		Prefix:  "",
		MaxKeys: 3,
	}
	v, _, err := tencentCOSClent.Bucket.Get(context.Background(), opt)
	if err != nil {
		fmt.Println("ere:", err)
		panic(err)
	}
	for _, c := range v.Contents {
		fmt.Printf("%s, %d\n", c.Key, c.Size)
	}
}
