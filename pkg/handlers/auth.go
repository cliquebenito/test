package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/internal/models"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	id, err := h.services.Auth.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "error responce")
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	token, err := h.services.Auth.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "error responce")
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}
