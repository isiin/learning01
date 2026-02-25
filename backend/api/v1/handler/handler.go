package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

type ErrorResponse struct {
	Status  int    `json:"status" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// エラー情報を格納したJSONレスポンスを生成する
func ErrorJson(c *gin.Context, status int, err error) {
	er := ErrorResponse{
		Status:  status,
		Message: err.Error(),
	}
	c.JSON(status, er)
}
