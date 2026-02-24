package controller

import (
	"log"
	"net/http"
	"react-ts/backend/model"

	"github.com/gin-gonic/gin"
)

// QuerySurveyors godoc
//
//	@Summary		指定した事業所の調査員のリストを返す
//	@Tags			surveyors
//	@Param			q	query		string	false	"事業所のid"
//	@Success		200	{array}		model.Surveyor
//	@Failure		default	{object}	httputil.HTTPError
//	@Router			/surveyors/query [get]
func (c *Controller) QuerySurveyors(ctx *gin.Context) {
	q := ctx.Query("q")

	// TODO
	log.Println(q)
	surveyors := []model.Surveyor{
		{ID: "001", Name: "調査員1"},
		{ID: "002", Name: "調査員2"},
		{ID: "003", Name: "調査員3"},
	}

	ctx.JSON(http.StatusOK, surveyors)
}
