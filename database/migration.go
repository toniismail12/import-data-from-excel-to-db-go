package database

import (
	"import-excel/models"
)

func Migration() {

	DB.AutoMigrate(

		&models.Test_bareng_wtr{},

	)

}
