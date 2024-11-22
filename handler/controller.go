package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sachin-gautam/gin-api/constants"
	"github.com/sachin-gautam/gin-api/database"
)

type Handler struct {
	db *database.Database
}

func NewHandler(db *database.Database) *Handler {
	return &Handler{
		db,
	}
}

type Error struct {
	Message string `json:"message"`
}

type Response struct {
	Data string `json:"data"`
}

func (h *Handler) GetDetail(c *gin.Context) {
	tenantID, exists := c.Get(constants.TenantID)
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, &Error{
			Message: "tenant not found",
		})
	}
	detail := h.db.FindDetailByTenantID(tenantID.(string))

	c.JSON(http.StatusOK, &Response{
		Data: detail,
	})
}
