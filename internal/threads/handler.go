package threads

import (
	"net/http"

	"github.com/Potagashev/breddit/internal/config"
	"github.com/Potagashev/breddit/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ThreadHandler struct {
	ThreadService *ThreadService
	Cfg *config.Config
}

func NewThreadHandler(threadService *ThreadService, cfg *config.Config) *ThreadHandler {
	return &ThreadHandler{ThreadService: threadService, Cfg: cfg}
}

// @Summary Create thread
// @Tags threads
// @Param requestBody body ThreadWrite true "Create data"
// @Param Access-Token header string true "Access Token"
// @Success 201 {object} ThreadWriteResponse
// @Router /threads [post]
func (h *ThreadHandler) CreateThread(c *gin.Context) {
	authHeader := c.GetHeader("Access-Token")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Access-Token header"})
		return
	}

	const bearerPrefix = "Bearer "
	if len(authHeader) <= len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return
	}

	tokenString := authHeader[len(bearerPrefix):]

	user, errr := utils.ParseToken(tokenString, *h.Cfg)

	if errr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	var thread ThreadWrite
	if err := c.ShouldBindJSON(&thread); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var threadId, err = h.ThreadService.CreateThread(&thread, user.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response = ThreadWriteResponse{ID: threadId}
	c.JSON(http.StatusCreated, response)
}

// @Summary Get thread
// @Tags threads
// @Param id path string true "Thread ID"
// @Success 200 {object} Thread
// @Router /threads/{id} [get]
func (h *ThreadHandler) GetThread(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid thread ID"})
		return
	}

	thread, err := h.ThreadService.GetThreadByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, thread)
}

// @Summary List threads
// @Tags threads
// @Success 200 {object} []Thread
// @Router /threads [get]
func (h *ThreadHandler) GetManyThreads(c *gin.Context) {
	threads, err := h.ThreadService.GetManyThreads()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, threads)
}

// @Summary Update thread
// @Tags threads
// @Param id path string true "Thread ID"
// @Param requestBody body ThreadWrite true "Update data"
// @Success 200 {object} ThreadWriteResponse
// @Router /threads/{id} [put]
func (h *ThreadHandler) UpdateThread(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid thread ID"})
		return
	}

	var thread Thread
	if err := c.ShouldBindJSON(&thread); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	thread.ID = id
	if err := h.ThreadService.UpdateThread(&thread); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response = ThreadWriteResponse{ID: id}
	c.JSON(http.StatusOK, response)
}

// @Summary Delete thread
// @Tags threads
// @Param id path string true "Thread ID"
// @Router /threads/{id} [delete]
func (h *ThreadHandler) DeleteThread(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid thread ID"})
		return
	}

	if err := h.ThreadService.DeleteThread(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
