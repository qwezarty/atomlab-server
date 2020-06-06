package forms

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/qwezarty/atomsrv/models"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine
var w *httptest.ResponseRecorder
var r *http.Request

func setup() {
	w = httptest.NewRecorder()

	router = gin.Default()
	Configure(router)
}

func teardown() {
}

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	teardown()
	os.Exit(retCode)
}

func mockTitanic() *models.Titanic {
	// mock := []float64{0., 22., 1., 0., 7.25, 0., 0., 0., 0.}
	// {"sex":"0", "age":"22", "n_siblings_spouses":"1", "parch":"0", "fare":"7.25", "class":"0", "deck":"0", "embark_town":"0", "alone":"0"}
	return &models.Titanic{
		Sex:              0,
		Age:              22,
		NSiblingsSpouses: 1,
		Parch:            0,
		Fare:             7.25,
		Class:            0,
		Deck:             0,
		EmbarkTown:       0,
		Alone:            0,
	}
}

func TestTitanicHandler(t *testing.T) {
	titanic := mockTitanic()
	data, _ := json.Marshal(titanic)

	r, _ = http.NewRequest("POST", "/forms/titanic", bytes.NewBuffer(data))
	r.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "17.93", w.Body.String())
}
