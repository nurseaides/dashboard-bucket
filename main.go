package bucket

import (
	"github.com/nurseaides/dashboard-bucket/component"
	. "github.com/nurseaides/dashboard-engine"
)

type ModuleConfig struct {
	Type      string `json:"bucket_type"`
	BucketUrl string `json:"bucket_url"`
	SecretID  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`
}

var config = new(ModuleConfig)

func init() {
	plugin := &PluginConfig{
		Name:   "bucket",
		Type:   PLUGIN_APP,
		Config: config,
		Run:    run,
	}
	InstallPlugin(plugin)
}

func run() {
	//初始化sdk
	if config.Type == "tencert-cos" {
		component.InitBucket(config.BucketUrl, config.SecretID, config.SecretKey)
	}

	//路由设置
	bindActions()
}

func bindActions() {
}
