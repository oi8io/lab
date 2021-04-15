package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oi.io/apps/discover/errcode"
	"oi.io/apps/discover/global"
	"oi.io/apps/discover/model"
)

func RegisterHandler(c *gin.Context) {
	var req model.RequestRegister
	if e := c.ShouldBindJSON(&req); e != nil {
		err := errcode.ParamError
		c.JSON(http.StatusOK, gin.H{
			"code":    err.Code(),
			"message": err.Error(),
		})
		return
	}
	//bind instance
	instance := model.NewInstance(&req)
	if instance.Status == 0 || instance.Status > 2 {
		err := errcode.ParamError
		c.JSON(http.StatusOK, gin.H{
			"code":    err.Code(),
			"message": err.Error(),
		})
		return
	}
	//dirtytime
	if req.DirtyTimestamp > 0 {
		instance.DirtyTimestamp = req.DirtyTimestamp
	}
	global.Discovery.Registry.Register(instance, req.LatestTimestamp)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "",
		"data":    "",
	})
}

