package threads

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ThreadHandler struct {
    ThreadService *ThreadService
}

// NewThreadHandler создает новый экземпляр ThreadHandler
func NewThreadHandler(threadService *ThreadService) *ThreadHandler {
    return &ThreadHandler{ThreadService: threadService}
}

// CreateThread создает новый тред
func (h *ThreadHandler) CreateThread(c *gin.Context) {
    var thread Thread
    if err := c.ShouldBindJSON(&thread); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    if err := h.ThreadService.CreateThread(&thread); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, thread)
}

// GetThread получает тред по ID
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

// GetThread получает тред по ID
func (h *ThreadHandler) GetManyThreads(c *gin.Context) {
    threads, err := h.ThreadService.GetManyThreads()
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, threads)
}

// UpdateThread обновляет данные тред
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

    c.JSON(http.StatusOK, thread)
}

// DeleteThread удаляет тред по ID
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
