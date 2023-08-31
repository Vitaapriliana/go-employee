package entity

type Employee struct {
	Id                 uint
	Name               string
	BirthDate          string
	EmployeeStatusCode EmployeeStatus
	Salary             uint
	IsDeleted          uint
}
