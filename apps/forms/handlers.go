package forms

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qwezarty/atomsrv/models"
)

// load parameters from h5df file
// var titanicParams :=

func Titanic(c *gin.Context) {
	titanic := &models.Titanic{}
	if err := c.Bind(titanic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// value cheking
	// transform to matrix
	// do the computation
	return
}
