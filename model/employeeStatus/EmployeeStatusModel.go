package employeeStatus

import (
	"go-web-employee/config"
	"go-web-employee/entity"
)

func GetAll() []entity.EmployeeStatus {
	rows, err := config.DB.Query(`SELECT * FROM t1_employee_status`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var employeeStatuses []entity.EmployeeStatus

	for rows.Next() {
		var employeeStatus entity.EmployeeStatus
		if err := rows.Scan(&employeeStatus.Id, &employeeStatus.Code, &employeeStatus.Name, &employeeStatus.IsDeleted); err != nil {
			panic(err)
		}

		employeeStatuses = append(employeeStatuses, employeeStatus)
	}

	return employeeStatuses
}

//func GetByName(name string) entity.EmployeeStatus {
//	row := config.DB.QueryRow(`
//        SELECT id, code, name, is_deleted FROM t1_employee_status WHERE name = ?`, name)
//
//	var employeeStatus entity.EmployeeStatus
//	err := row.Scan(
//		&employeeStatus.Id,
//		&employeeStatus.Code,
//		&employeeStatus.Name,
//		&employeeStatus.IsDeleted,
//	)
//
//	if err != nil {
//		panic(err)
//	}
//
//	return employeeStatus
//}

func GetCodeById(id string) (string, error) {
	row := config.DB.QueryRow(`
        SELECT code FROM t1_employee_status WHERE id = ?`, id)

	var code string
	err := row.Scan(&code)
	if err != nil {
		return "", err
	}

	return code, nil
}
