package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurseaides/dashboard-bucket/component"
)

//TmpCredentForTencentCos 设备列表
func TmpCredentForTencentCos(c *gin.Context) {

	credential, err := component.TencentCOSGetCredential()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "服务器出错",
			"errors":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "",
			"data":    credential,
		})
	}
}
