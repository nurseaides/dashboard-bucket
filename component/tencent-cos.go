package component

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
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

// TencentCOSGetCredential 获取腾迅云对象存储临时密钥
func TencentCOSGetCredential() (*sts.CredentialResult, error) {
	c := sts.NewClient(
		config.SecretID,
		config.SecretKey,
		nil,
	)
	opt := &sts.CredentialOptions{
		DurationSeconds: int64(300),
		Region:          "ap-nanjing",
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					Action: []string{
						"name/cos:PostObject",
						"name/cos:PutObject",
						"name/cos:GetObject",
					},
					Effect: "allow",
					Resource: []string{
						//这里改成允许的路径前缀，可以根据自己网站的用户登录态判断允许上传的具体路径，例子： a.jpg 或者 a/* 或者 * (使用通配符*存在重大安全风险, 请谨慎评估使用)
						"*",
					},
				},
			},
		},
	}
	return c.GetCredential(opt)
}
