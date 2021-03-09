/*
 * @Author: Kim
 * @Date: 2021-03-08 14:00:59
 * @LastEditTime: 2021-03-09 10:44:00
 * @LastEditors: Kim
 * @Description:
 * @FilePath: /scaffold_go/internal/routers/router.go
 */
package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	v1 "demo/internal/routers/api/v1"

	_ "demo/docs" // docs is generated by Swag CLI, you have to import it.
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	tag := v1.NewTag()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags", tag.List)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
