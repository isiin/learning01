package v1

import (
	"react-ts/backend/api/v1/handler"

	"github.com/gin-gonic/gin"
)

// @title	react-ts backend API
// @version	1.0
// @BasePath	/v1
func Route(r *gin.Engine) {
	// 各エンドポイントのルーティング
	h := handler.NewHandler()

	v1 := r.Group("/v1")
	{
		surveyors := v1.Group("/surveyors")
		{
			surveyors.GET("query", h.QuerySurveyors)
		}
		samples := v1.Group("/samples")
		{
			samples.GET("", h.GetSamples)
		}
	}
}
