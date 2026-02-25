package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type GetSampleRequest struct {
	Q1 string    `form:"q1" binding:"required,alphanum,len=3"`
	Q2 string    `form:"q2" binding:"max=3"`
	Q3 string    `form:"q3" binding:"uuid"`
	Q4 string    `form:"q4" binding:"email"`
	Q5 []int     `form:"q5" collection_format:"csv"`
	Q6 time.Time `form:"q6" time_format:"2006-01-02" time_utc:"1"`
}

type GetSampleResponse struct {
	ID   string `json:"id" example:"000001"`
	Name string `json:"name" example:"調査員1"`
}

// GetSamples godoc
//
//	@Summary		実験用
//	@Tags			samples
//	@Param			req	query	GetSampleRequest true	"検索条件"
//	@Success		200	{object}	GetSampleResponse	"取得結果"
//	@Failure		400	{object}	ErrorResponse	"検証エラー"
//	@Router			/samples [get]
func (h *Handler) GetSamples(c *gin.Context) {
	// TODO
	var req GetSampleRequest
	if err := c.ShouldBind(&req); err != nil {
		ErrorJson(c, http.StatusBadRequest, err)
		return
	}

	// TODO
	log.Println(req)

	res := GetSampleResponse{ID: "001", Name: "調査員1"}
	c.JSON(http.StatusOK, res)
}
