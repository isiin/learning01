package handler

import (
	"react-ts/backend/internal/errs"

	"github.com/gin-gonic/gin"
)

type GetSurveyorsRequest struct {
	ID       string `form:"id" binding:"omitempty,max=6,numeric" example:"000001"`
	OfficeID string `form:"office-id" binding:"omitempty,max=2,alphanum" example:"XX"`
	View     string `form:"view" binding:"omitempty,oneof=basic detail" example:"detail"` // 表示形式を指定するパラメータ
}

type GetSurveyorsResponse struct {
	ID         string  `json:"id" example:"000001"`
	Name       string  `json:"name" example:"調査員1"`
	OfficeID   *string `json:"office-id,omitempty" example:"XX"`
	OfficeName *string `json:"office,omitempty" example:"札幌事業所"` // omitemptyをつけると、nilの場合はJSONに出力されません
}

// GetSurveyors godoc
//
//	@Summary		指定した事業所の調査員のリストを返す
//	@Tags			surveyors
//	@Param			q	query		GetSurveyorsRequest	true	"検索条件"
//	@Success		200	{array}		GetSurveyorsResponse "調査員のリスト"
//	@Failure		400	{object}	ErrorResponse "不正なリクエスト"
//	@Failure		500	{object}	ErrorResponse	"想定外のエラー"
//	@Router			/surveyors [get]
func GetSurveyors() gin.HandlerFunc {
	return func(c *gin.Context) {

		var p GetSurveyorsRequest
		if err := c.ShouldBind(&p); err != nil {
			details := createValidationDetails(err)
			err := errs.NewBusinessError(errs.InvalidRequest, details...)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}

		// TODO
		surveyors := []GetSurveyorsResponse{
			{ID: "000001", Name: "調査員1"},
			{ID: "000002", Name: "調査員2"},
			{ID: "000003", Name: "調査員3"},
		}

		// view=detail が指定された場合のみ部署名を設定する
		if p.View == "detail" {
			oid := "XX"
			o := "札幌事業所"
			for i := range surveyors {
				surveyors[i].OfficeID = &oid
				surveyors[i].OfficeName = &o
			}
		}

		c.JSON(200, surveyors)
	}
}
