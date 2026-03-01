package handler

import (
	"react-ts/backend/internal/domain"
	"react-ts/backend/internal/errs"

	"github.com/gin-gonic/gin"
)

type GetSurveyorsRequest struct {
	OfficeID string `form:"office-id" binding:"omitempty,alphanum,max=2" example:"XX"`
}

type GetSurveyorsResponse struct {
	ID   string `json:"id" example:"000001"`
	Name string `json:"name" example:"調査員1"`
}

// GetSurveyors godoc
//
//	@Summary		指定条件の調査員のリストを返す
//	@Tags			surveyors
//	@Param			q	query		GetSurveyorsRequest	true	"検索条件"
//	@Success		200	{array}		GetSurveyorsResponse "調査員のリスト"
//	@Failure		400	{object}	ErrorResponse "リクエスト形式不正"
//	@Failure		500	{object}	ErrorResponse	"想定外のエラー"
//	@Router			/surveyors [get]
func GetSurveyors(uc domain.SurveyUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {

		var p GetSurveyorsRequest
		if err := c.ShouldBind(&p); err != nil {
			details := createValidationDetails(err)
			err := errs.NewBusinessError(errs.InvalidRequest, details...)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}

		md, err := uc.GetSurveyors(domain.SurveyorFilter{OfficeID: p.OfficeID})
		if err != nil {
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}

		res := make([]GetSurveyorsResponse, 0, len(md))
		for _, m := range md {
			r := GetSurveyorsResponse{
				ID:   m.ID,
				Name: m.Name,
			}
			res = append(res, r)
		}
		c.JSON(200, res)
	}
}
