package router

import (
    "github.com/gin-gonic/gin"
	"github.com/Potagashev/breddit/internal/threads"
)

func NewRouter(userService *threads.ThreadService) *gin.Engine {
    router := gin.Default()

    threadHandler := threads.NewThreadHandler(userService)

	router.POST("/threads", threadHandler.CreateThread)
    router.GET("/threads/:id", threadHandler.GetThread)
	router.GET("/threads", threadHandler.GetManyThreads)
    router.PUT("/threads/:id", threadHandler.UpdateThread)
    router.DELETE("/threads/:id", threadHandler.DeleteThread)

    return router
}