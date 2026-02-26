package handler

import (
	"log"
	"net/http"
	"react-ts/backend/domain"
	"time"

	"github.com/gin-gonic/gin"
)

type GetSampleRequest struct {
	Q1       string    `form:"q1" binding:"required,alphanum,len=3"`
	Q2       string    `form:"q2" binding:"omitempty,max=3,min=2"`
	UUID     string    `form:"uuid" binding:"omitempty,uuid"`
	Email    string    `form:"email" binding:"omitempty,email"`
	IntArray []int     `form:"intArray" collection_format:"csv"`
	DateUtc  time.Time `form:"dateUtc" time_format:"2006-01-02" time_utc:"1"`
}

type GetSampleResponse struct {
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
	var p GetSampleRequest
	if err := c.ShouldBind(&p); err != nil {
		details := createValidationDetails(err)
		ResponseErrorJson(c, http.StatusBadRequest, domain.ECInvalidRequest, details)
		return
	}

	// TODO
	log.Println(p)

	res := GetSampleResponse{Name: "調査員1"}
	c.JSON(http.StatusOK, res)
}
