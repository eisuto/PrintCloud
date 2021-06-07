package api

import (
	"PrintCloud/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Tags 打印机
// @Summary 获取全部打印机
// @Accept  json
// @Produce json
// @Success 200 {string} "{"code":1,"data":[{"id":1,"status":"ok"}],"msg": "操作成功"}"
// @Failure 400 {string} "{"code":1,"data":[],"msg": "操作失败"}"
// @Router /api/printers [get]
func Printers(c *gin.Context) {
	var printer model.Printer
	var code int64 = 1
	result, err := printer.Printers()

	if err != nil {
		code = -1
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": result,
		"msg":  "操作成功",
	})

}
