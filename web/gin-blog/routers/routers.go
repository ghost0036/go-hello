package routers

import (
	"github.com/ghost0036/go-hello/web/gin-blog/middleware/jwt"
	"github.com/ghost0036/go-hello/web/gin-blog/pkg/setting"
	"github.com/ghost0036/go-hello/web/gin-blog/routers/api"
	"github.com/ghost0036/go-hello/web/gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.Use(jwt.JWT())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("api/v1")
	{
		/**
		Tag操作
		*/
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.DELETE("/tags/:id", v1.EditTag)
		apiv1.PUT("/tags/:id", v1.EditTag)

		/**
		文章操作
		*/
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
