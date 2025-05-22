package __controller

import (
	"github.com/gin-gonic/gin"
	"internal/__entity"
	"internal/__usecase"
	"net/http"
)

type UserHandler struct {
	uc *__usecase.UserUsecase
}

func NewUserHandler(uc *__usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) Register(c *gin.Context) {
	var user __entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.RegisterUser(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.uc.GetAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
