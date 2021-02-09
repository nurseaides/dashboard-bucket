package bucket

import (
	"github.com/nurseaides/dashboard-bucket/component"
	"github.com/nurseaides/dashboard-bucket/controllers"
	. "github.com/nurseaides/dashboard-engine"
)

type moduleConfig struct {
	Cos component.BucketConfig `json:"cos"`
}

var config = new(moduleConfig)

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
	component.InitBucket(config.Cos)

	//路由设置
	bindActions()
}

func bindActions() {
	GinR.GET("/bucket/tmp-credent-for-tencent-cos", controllers.TmpCredentForTencentCos)
}
