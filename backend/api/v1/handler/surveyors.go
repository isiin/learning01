package handler

import (
	"net/http"
	"react-ts/backend/domain"

	"github.com/gin-gonic/gin"
)

type SurveyorRequest struct {
	Q string `form:"q" binding:"required" example:"XXX"`
}

type Surveyor struct {
	ID   string `json:"id" example:"000001"`
	Name string `json:"name" example:"調査員1"`
}

// QuerySurveyors godoc
//
//	@Summary		指定した事業所の調査員のリストを返す
//	@Tags			surveyors
//	@Param			q	query		SurveyorRequest	true	"検索条件"
//	@Success		200	{array}		Surveyor "調査員のリスト"
//	@Failure		400	{object}	ErrorResponse "検証エラー"
//	@Router			/surveyors/query [get]
func (h *Handler) QuerySurveyors(c *gin.Context) {

	var p SurveyorRequest
	if err := c.ShouldBind(&p); err != nil {
		details := createValidationDetails(err)
		ResponseErrorJson(c, http.StatusBadRequest, domain.ECInvalidRequest, details)
		return
	}

	// TODO
	surveyors := []Surveyor{
		{ID: "001", Name: "調査員1"},
		{ID: "002", Name: "調査員2"},
		{ID: "003", Name: "調査員3"},
	}

	c.JSON(http.StatusOK, surveyors)
}
