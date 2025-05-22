package __controller

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine, h *UserHandler) {
	r.POST("/register", h.Register)
	r.GET("/users", h.GetAll)
}
