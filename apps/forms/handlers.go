package forms

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qwezarty/atomsrv/models"
	"github.com/qwezarty/atomsrv/utils"
)

// initialize all the parameters we need when it startup
var parameters, scales, _ = utils.InitializeParameters("titanic_neural_network.json")

/* final testing using curl
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"sex":"0", "age":"22", "n_siblings_spouses":"1", "parch":"0", "fare":"7.25", "class":"0", "deck":"0", "embark_town":"0", "alone":"0"}' \
  http://localhost:30096/forms/titanic
*/
func TitanicHandler(c *gin.Context) {
	titanic := &models.Titanic{}
	if errBind := c.Bind(titanic); errBind != nil {
		log.Println(errBind.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}
	// todo value cheking
	// transform titanic object to array
	X := models.Titanic2Slice(titanic)

	// predict the propability of surviving
	propability, errPredict := utils.PredictLogisticInstance(X, parameters, scales)
	if errPredict != nil {
		log.Println(errPredict.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": errPredict.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"propability": fmt.Sprintf("%.2f", propability*100)})
}
