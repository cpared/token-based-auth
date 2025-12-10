package tasks

import (
	"context"
	"net/http"
	model "token-based-auth/internal/domain"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Get(ctx context.Context, tID string) *model.Task
	Create(ctx context.Context, i model.Info) *model.Task
}

type Handler struct {
	serv Service
}

type BodyRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func New(s Service) *Handler {
	return &Handler{
		serv: s,
	}
}

func (h *Handler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, found := c.Params.Get("id")
		if !found {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid ID"})
			return
		}

		var body BodyRequest
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid body"})
			return
		}

		t := h.serv.Get(c.Request.Context(), id)

		c.JSON(http.StatusOK, gin.H{"result": t})
	}
}

func (h *Handler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body BodyRequest
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid body"})
			return
		}

		t := h.serv.Create(c.Request.Context(), model.Info{
			Description: body.Description,
			Title:       body.Title,
		})

		c.JSON(http.StatusOK, gin.H{"result": t})
	}
}
