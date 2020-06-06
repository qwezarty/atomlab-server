package forms

import (
	"github.com/gin-gonic/gin"
)

// Configure registers all handler
func Configure(r *gin.Engine) {
	r.POST("/forms/titanic", TitanicHandler)
}
