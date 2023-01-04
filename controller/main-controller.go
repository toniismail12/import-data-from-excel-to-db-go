package controller

import (
	"log"
	"fmt"
	"import-excel/models"
	"import-excel/database"
	"github.com/gofiber/fiber/v2"
	"github.com/360EntSecGroup-Skylar/excelize"
	"time"
)

func Import(c *fiber.Ctx) error {

	fileName := "excel.xlsx"

	xlsx, err := excelize.OpenFile("./files/" + fileName)
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	for _, sheetName := range xlsx.GetSheetMap() {

		rows := xlsx.GetRows(sheetName)

		jr := len(rows)

		for i := 2; i < jr+1; i++ {

			// if i==11 {
			// 	break
			// }
			
			periode := xlsx.GetCellValue(sheetName, fmt.Sprintf("K%d", i))
			date := periode[:4] +"-"+ periode[4:]+"-01"
			fmt.Println(date)
			
			insert := models.Test_bareng_wtr{

				Badge: xlsx.GetCellValue(sheetName, fmt.Sprintf("A%d", i)),
				Cost_center: xlsx.GetCellValue(sheetName, fmt.Sprintf("M%d", i)),
				Payroll_area: xlsx.GetCellValue(sheetName, fmt.Sprintf("G%d", i)),
				Periode: date,
				Wage_type: xlsx.GetCellValue(sheetName, fmt.Sprintf("P%d", i)),
				Amount: xlsx.GetCellValue(sheetName, fmt.Sprintf("S%d", i)),
				Created_at: time.Now().Local(),
				Updated_at: time.Now().Local(),
				File_code: "X1X1",
				
			}

			err := database.DB.Create(&insert)
			if err != nil {
				log.Println(err)
			}

		}

	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "success import",
	})

}

