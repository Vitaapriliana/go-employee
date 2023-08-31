package employeeModel

import (
	"go-web-employee/config"
	"go-web-employee/entity"
)

func GetAll() []entity.Employee {
	rows, err := config.DB.Query(
		`SELECT
   			employee.id,
      			employee.name,
      			employee.birth_date,
      			employeeStatus.name AS employe_status_code,
      			employee.salary FROM t2_employee employee
				JOIN t1_employee_status employeeStatus ON employeeStatus.code = employee.employee_status_code
				WHERE employee.is_deleted = 0`)

	if err != nil {
		panic(err)
	}

	var employees []entity.Employee

	for rows.Next() {
		var employee entity.Employee
		if err := rows.Scan(
			&employee.Id,
			&employee.Name,
			&employee.BirthDate,
			&employee.EmployeeStatusCode.Code,
			&employee.Salary,
		); err != nil {
			panic(err)
		}

		employees = append(employees, employee)
	}

	return employees
}

func Create(employee entity.Employee) bool {
	result, err := config.DB.Exec(
		`INSERT INTO t2_employee (name, birth_date, employee_status_code, salary, is_deleted) values(?,?,?,?,?)`,
		employee.Name,
		employee.BirthDate,
		employee.EmployeeStatusCode.Code,
		employee.Salary,
		employee.IsDeleted)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Update(id int, employee entity.Employee) bool {
	query, err := config.DB.Exec(`
		UPDATE t2_employee SET
			name = ?,
			birth_date = ?,
			employee_status_code = ? ,
			salary = ?,
			is_deleted = ?
		WHERE id = ?`,
		employee.Name,
		employee.BirthDate,
		employee.EmployeeStatusCode.Code,
		employee.Salary,
		employee.IsDeleted,
		id,
	)

	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Detail(id int) entity.Employee {
	row := config.DB.QueryRow(`
			SELECT
			employee.id,
			employee.name,
			employee.birth_date,
			employeeStatus.code as employee_status_code,
			employee.salary FROM t2_employee employee
		JOIN t1_employee_status employeeStatus ON employee.employee_status_code = employeeStatus.code
		WHERE employee.id = ?
	`, id)

	var employee entity.Employee

	err := row.Scan(
		&employee.Id,
		&employee.Name,
		&employee.BirthDate,
		&employee.EmployeeStatusCode.Code,
		&employee.Salary,
	)

	if err != nil {
		panic(err)
	}

	return employee
}
func Delete(id int) error {

	_, err := config.DB.Exec(`
		UPDATE t2_employee SET
			is_deleted = ?
		WHERE id = ?`,
		1,
		id)
	return err
}
