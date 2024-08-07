package auth

import (
	"fmt"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	// 请求对象
	type PhoneExistRequest struct {
		Phone string `json:"phone"`
	}

	request := PhoneExistRequest{}

	// 解析 JSON 请求
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败,返回 422 状态吗和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		// 打印错误数据
		fmt.Println(err.Error())
		// 出错了，中断请求
		return
	}

	// 表单验证
	errs := requests.ValidateSignupPhoneExist(&request, c)
	// errs 返回长度等于零通过
	if len(errs) > 0 {
		// 验证失败
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	// 检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exit": user.IsphoneExist(request.Phone),
	})
}
