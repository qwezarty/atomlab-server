package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/qwezarty/atomsrv/apps"
	"github.com/qwezarty/atomsrv/apps/forms"
	"github.com/qwezarty/atomsrv/engine"
)

func main() {
	// startup all managers
	var router = gin.Default()
	var db = engine.Startup("sqlite3")

	// register all sub-routes
	apps.Configure(db)
	forms.Configure(router, db) // singleton, pass by pointer

	log.Fatal(router.Run("0.0.0.0:30096"))
}
