package forms

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Configure registers all handler
func Configure(r *gin.Engine, engine *gorm.DB) {
	db = engine

	r.POST("/forms/titanic", Titanic)
}
