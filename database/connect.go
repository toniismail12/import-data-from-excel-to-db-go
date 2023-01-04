package database

import (
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"import-excel/config"
)

var DB *gorm.DB

func Connect() {
	
	host := config.Environment("DB_HOST")
	user := config.Environment("DB_USERNAME")
	pass := config.Environment("DB_PASSWORD")
	dbport := config.Environment("DB_PORT")
	db_name := config.Environment("DB_DATABASE")

	// koneksi ke database
	dsn := "sqlserver://" + user + ":" + pass + "@" + host + ":" + dbport + "?" + "database=" + db_name
	
	database, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("gagal connect database")
	} else {
		log.Println("connect success")
	}

	// migrate tabel
	DB = database
	
	// Migration()

}
