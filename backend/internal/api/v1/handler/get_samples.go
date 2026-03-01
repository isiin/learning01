package handler

import (
	"react-ts/backend/internal/domain"
	"react-ts/backend/internal/errs"
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
	ID   string `json:"id" example:"11111"`
	Name string `json:"name" example:"サンプル1"`
}

// GetSamples godoc
//
//	@Summary		実験用
//	@Tags			samples
//	@Param			req	query	GetSampleRequest true	"検索条件"
//	@Success		200	{object}	GetSampleResponse	"取得結果"
//	@Failure		400	{object}	ErrorResponse	"不正なリクエスト"
//	@Failure		500	{object}	ErrorResponse	"想定外のエラー"
//	@Router			/samples [get]
func GetSamples(uc domain.SamplesUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p GetSampleRequest
		if err := c.ShouldBind(&p); err != nil {
			details := createValidationDetails(err)
			err := errs.NewBusinessError(errs.InvalidRequest, details...)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}

		md, err := uc.GetSamples()
		if err != nil {
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}

		res := make([]GetSampleResponse, 0, len(md))
		for _, m := range md {
			r := GetSampleResponse{
				ID:   m.ID,
				Name: m.Name,
			}
			res = append(res, r)
		}
		c.JSON(200, res)
	}
}
