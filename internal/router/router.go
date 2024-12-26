package router

import (
    "github.com/gin-gonic/gin"
    _ "github.com/Potagashev/breddit/docs"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/Potagashev/breddit/internal/threads"
)

func NewRouter(threadHandler *threads.ThreadHandler) *gin.Engine {
    router := gin.Default()

    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    threadRoutes := router.Group("/api/v1/threads")
    {
        threadRoutes.POST("", threadHandler.CreateThread)
        threadRoutes.GET("/:id", threadHandler.GetThread)
        threadRoutes.GET("", threadHandler.GetManyThreads)
        threadRoutes.PUT("/:id", threadHandler.UpdateThread)
        threadRoutes.DELETE("/:id", threadHandler.DeleteThread)
    }

    return router
}