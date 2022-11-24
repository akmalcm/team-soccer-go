package configs

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	gormysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConn() (*sql.DB, error) {
	env, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Capture connection properties.
	cfg := mysql.Config{
		User:                 env["DBUSER"],
		Passwd:               env["DBPASS"],
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               env["DBNAME"],
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	mysqldb, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db, err := gorm.Open(gormysql.New(gormysql.Config{
		Conn: mysqldb,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	DB = db
	return mysqldb, nil
}

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page == 0 {
			page = 1
		}

		limit, _ := strconv.Atoi(c.Query("limit"))
		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 10
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}
