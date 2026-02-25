package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Surveyor struct {
	ID   string `json:"id" example:"000001"`
	Name string `json:"name" example:"調査員1"`
}

// QuerySurveyors godoc
//
//	@Summary		指定した事業所の調査員のリストを返す
//	@Tags			surveyors
//	@Param			q	query		string	true	"事業所のid"
//	@Success		200	{array}		Surveyor "調査員のリスト"
//	@Failure		400	{object}	ErrorResponse "検証エラー"
//	@Router			/surveyors/query [get]
func (h *Handler) QuerySurveyors(c *gin.Context) {
	// TODO
	q := c.Query("q")
	if len(q) < 1 {
		// TODO
		ErrorJson(c, http.StatusBadRequest, errors.New("事業所idが未設定です"))
		return
	}

	// TODO
	log.Println(q)
	surveyors := []Surveyor{
		{ID: "001", Name: "調査員1"},
		{ID: "002", Name: "調査員2"},
		{ID: "003", Name: "調査員3"},
	}

	c.JSON(http.StatusOK, surveyors)
}
