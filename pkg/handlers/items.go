package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/internal/models"
	"test/pkg/service"
)

func (h *Handler) getItemsHandler(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if userId == -1 {
		newErrorResponse(c, http.StatusUnauthorized, "user not found")
		return
	}
	var input models.Input
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	item, err := service.GetInfoForId(input.ID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}
