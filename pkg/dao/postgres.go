package dao

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	// import postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// PgConn it is safe for goroutines. Use PgConn for all postgres request.
	PgConn *gorm.DB
)

func ConnectPG(host, port, user, dbname, passwd string) (err error) {
	dbEndPoint := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=2", host, port, user, passwd, dbname)
	PgConn, err = gorm.Open("postgres", dbEndPoint)
	if err != nil {
		return
	}

	PgConn.DB().SetConnMaxLifetime(time.Minute * 5)
	PgConn.DB().SetMaxIdleConns(20)
	PgConn.DB().SetMaxOpenConns(500)
	return
}
