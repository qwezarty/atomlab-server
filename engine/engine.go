package engine

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var Sqlite = "./engine/engine.db"

func Startup(dialect string, values ...interface{}) *gorm.DB {
	// open connection
	engine, err := gorm.Open(dialect, getConn(dialect))
	if err != nil {
		// give a panic and exit immediately
		log.Fatalf("faltal error occour when conn to db: %v", err)
	}

	// sync tables to db
	if len(values) == 0 {
	}
	return engine.AutoMigrate(values...)
}

func getConn(dialect string) string {
	if got := os.Getenv("DB"); got != "" {
		dialect = got
	}

	switch dialect {
	case "sqlite3":
		return Sqlite
	default:
		return ""
	}
}
