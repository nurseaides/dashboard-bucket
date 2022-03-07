package component

import (
	"fmt"
	"io"
)

var config *BucketConfig

type bucketObject struct {
	Type string
}

// InitBucket 初始化微信配置
func InitBucket(conf BucketConfig) {
	config = &conf
	config.Type = config.Type
	if config.Type == "tencent-cos" {
		initTencentCOS(config.BucketUrl, config.SecretID, config.SecretKey)
	}
}

// PutObject  上传文件
func PutObject(path string, r io.Reader) (uri string, err error) {
	if config.Type == "tencent-cos" {
		return fmt.Sprintf("%s%s", config.BucketUrl, path), tencentCOSPutObject(path, r)
	}
	return uri, err
}

// ListObject 对象列表
func ListObject() {
	if config.Type == "tencent-cos" {
		tencentCOSListObject()
	}
}
