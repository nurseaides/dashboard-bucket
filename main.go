package bucket

import (
	"sync"

	"github.com/gin-gonic/gin"
	. "github.com/logrusorgru/aurora"
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
		Name:      "bucket",
		Type:      PLUGIN_GW,
		Config:    config,
		Run:       run,
		BindRoute: bindActions,
	}
	InstallPlugin(plugin)
}

func run() {
	//初始化sdk
	component.InitBucket(config.Cos)

}

//路由设置
func bindActions(g *gin.Engine, wg *sync.WaitGroup) {
	g.GET("/bucket/tmp-credent-for-tencent-cos", controllers.TmpCredentForTencentCos)

	Print(Green("bindActions bucket success"))

	wg.Done()
}
