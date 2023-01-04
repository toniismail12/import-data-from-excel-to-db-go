package models

import "time"

type Test_bareng_wtr struct {
	Id                      uint `json:"id"`
	Badge 		string
	Cost_center 	string
	Payroll_area 	string
	Periode 	string
	Wage_type 	string
	Amount 	string
	Created_at 	time.Time
	Updated_at 	time.Time
	File_code 	string

}
