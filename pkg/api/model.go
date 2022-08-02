package api

import "time"

type Employee struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Waktu    time.Time `json:"input_date"`
}

type Employees []Employee

type Company struct {
	CompanyName  string    `json:"company_name"`
	EmployeeList Employees `json:"employee_list"`
}
