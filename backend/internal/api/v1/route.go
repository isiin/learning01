package v1

import (
	"react-ts/backend/internal/api/v1/handler"
	"react-ts/backend/internal/bootstrap"

	"github.com/gin-gonic/gin"
)

// @title	react-ts backend API
// @version	1.0
// @BasePath	/v1
func Route(r *gin.Engine, cp *bootstrap.Components) {

	v1 := r.Group("/v1")

	v1.Use(handler.ErrorHandler())
	v1.GET("/surveyors", handler.GetSurveyors(cp.SurveyUC))
	v1.GET("/samples", handler.GetSamples(cp.SampleUC))
}
