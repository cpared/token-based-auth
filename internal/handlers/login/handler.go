package login

import (
	"context"
	"net/http"
	model "token-based-auth/internal/domain"

	"github.com/gin-gonic/gin"
)

type Body struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type Service interface {
	Login(ctx context.Context, secret, ID, pass string) (*model.Credentials, error)
}

type Handler struct{
	serv Service
}

func New(s Service) *Handler {
	return &Handler{
		serv: s,
	}
}

func (h *Handler) Login(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body Body
		if err := c.ShouldBindBodyWithJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid json payload"})
			return
		}

		// TODO: login service call
		cred, err := h.serv.Login(c.Request.Context(), secret, body.User, body.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": cred})
	}
}
