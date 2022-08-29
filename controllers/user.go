package controllers

import (
	"gin-demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetUserResponse struct {
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	User    models.User `json:"user"`
}

func GetUser(c *gin.Context) {
	res := GetUserResponse{
		Code:    200,
		Message: "ok",
		User: models.User{
			Id:       1,
			UserName: "username",
			Password: "1",
		},
	}
	c.JSON(http.StatusOK, &res)
}
