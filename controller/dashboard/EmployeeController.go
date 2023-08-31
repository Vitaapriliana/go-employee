package dashboard

import (
	"go-web-employee/entity"
	"go-web-employee/model/employeeModel"
	"go-web-employee/model/employeeStatus"
	"html/template"
	"net/http"
	"strconv"
)

func Index(response http.ResponseWriter, request *http.Request) {

	employee := employeeModel.GetAll()
	data := map[string]any{
		"employee": employee,
	}

	temp, err := template.ParseFiles("view/employee/Index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(response, data)
}

func Add(Response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		temp, err := template.ParseFiles("view/employee/Create.html")
		if err != nil {
			panic(err)
		}

		employeeStatuses := employeeStatus.GetAll()
		data := map[string]any{
			"employeeStatus": employeeStatuses,
		}

		temp.Execute(Response, data)
	}

	if request.Method == "POST" {
		var employee entity.Employee

		salary, err := strconv.Atoi(request.FormValue("salary"))
		if err != nil {
			panic(err)
		}

		employee.Name = request.FormValue("name")
		employee.BirthDate = request.FormValue("birth_date")
		employee.EmployeeStatusCode.Code = request.FormValue("employee_status_code")
		employee.Salary = uint(salary)
		employee.IsDeleted = 0
		if ok := employeeModel.Create(employee); !ok {
			http.Redirect(Response, request, request.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(Response, request, "/employee", http.StatusSeeOther)
	}
}

func Detail(Response http.ResponseWriter, request *http.Request) {
	idString := request.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	employee := employeeModel.Detail(id)
	data := map[string]any{
		"employee": employee,
	}

	temp, err := template.ParseFiles("view/employee/Detail.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(Response, data)
}

func Edit(Response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		temp, err := template.ParseFiles("view/employee/Edit.html")
		if err != nil {
			panic(err)
		}

		idString := request.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		employee := employeeModel.Detail(id)
		employeeStatuses := employeeStatus.GetAll()

		data := map[string]any{
			"employee":       employee,
			"employeeStatus": employeeStatuses,
		}

		temp.Execute(Response, data)
	}

	if request.Method == "POST" {
		var employee entity.Employee

		idString := request.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		salary, err := strconv.Atoi(request.FormValue("salary"))
		if err != nil {
			panic(err)
		}

		employee.Name = request.FormValue("name")
		employee.EmployeeStatusCode.Code = request.FormValue("employee_status_code")
		employee.Salary = uint(salary)
		employee.BirthDate = request.FormValue("birth_date")

		if ok := employeeModel.Update(id, employee); !ok {
			http.Redirect(Response, request, request.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}
		http.Redirect(Response, request, "/employee", http.StatusSeeOther)
	}
}

func Delete(Response http.ResponseWriter, request *http.Request) {
	idString := request.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := employeeModel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(Response, request, "/employee", http.StatusSeeOther)
}
