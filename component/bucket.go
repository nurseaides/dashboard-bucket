package component

import (
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
func PutObject(path string, r io.Reader) (err error) {
	if config.Type == "tencent-cos" {
		return tencentCOSPutObject(path, r)
	}
	return err
}

// ListObject 对象列表
func ListObject() {
	if config.Type == "tencent-cos" {
		tencentCOSListObject()
	}
}
