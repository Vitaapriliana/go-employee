package authController

//import (
//	"go-web-employee/library"
//	"go-web-employee/model/UserModel"
//	"net/http"
//)
//
//type UserInput struct {
//	Username string `validate:"required"`
//	Password string `validate:"required"`
//}
//
//var userModel = UserModel.NewUserModel()
//var validation = library.NewValidation()
//
//func Login(response http.ResponseWriter, request *http.Request) {

//if request.Method == http.MethodGet {
//	temp, _ := template.ParseFiles("view/employee/Login.html")
//	temp.Execute(response, nil)
//} else if request.Method == http.MethodPost {
//	// proses login
//	request.ParseForm()
//	UserInput := &UserInput{
//		Username: request.Form.Get("username"),
//		Password: request.Form.Get("password"),
//	}
//
//	errorMessages := validation.Struct(UserInput)
//
//	if errorMessages != nil {
//
//		data := map[string]interface{}{
//			"validation": errorMessages,
//		}
//
//		temp, _ := template.ParseFiles("view/employee/Login.html")
//		temp.Execute(response, data)
//
//	} else {
//
//	var user entity.User
//	userModel.Where(&user, "username", UserInput.Username)
//
//	var message error
//	if user.Username == "" {
//		message = errors.New("Username atau Password salah!")
//	} else {
//		// pengecekan password
//		errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))
//		if errPassword != nil {
//			message = errors.New("Username atau Password salah!")
//		}
//	}
//
//	if message != nil {
//
//		data := map[string]interface{}{
//			"error": message,
//		}
//
//		temp, _ := template.ParseFiles("view/employee/Login.html")
//		temp.Execute(response, data)
//	} else {
//		// set session
//		session, _ := config.Store.Get(request, config.SESSION_ID)
//
//		session.Values["loggedIn"] = true
//		session.Values["name"] = user.Name
//		session.Values["username"] = user.Username
//
//		session.Save(request, response)
//
//		http.Redirect(response, request, "/", http.StatusSeeOther)
//	}
//}
//
//}

//	if request.Method == http.MethodGet {
//		temp, _ := template.ParseFiles("view/employee/Login.html")
//		temp.Execute(response, nil)
//	} else if request.Method == http.MethodPost {
//		// proses login
//		request.ParseForm()
//		UserInput := &UserInput{
//			Username: request.Form.Get("username"),
//			Password: request.Form.Get("password"),
//		}
//
//		errorMessages := validation.Struct(UserInput)
//
//		if errorMessages != nil {
//
//			data := map[string]interface{}{
//				"validation": errorMessages,
//			}
//
//			temp, _ := template.ParseFiles("view/employee/Login.html")
//			temp.Execute(response, data)
//
//		} else {
//
//			var user entity.User
//			userModel.Where(&user, "username", UserInput.Username)
//
//			var message error
//			if user.Username == "" {
//				message = errors.New("Username atau Password salah!")
//			} else {
//				// pengecekan password
//				errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))
//				if errPassword != nil {
//					message = errors.New("Username atau Password salah!")
//				}
//			}
//
//			if message != nil {
//
//				data := map[string]interface{}{
//					"error": message,
//				}
//
//				temp, _ := template.ParseFiles("view/employee/Login.html")
//				temp.Execute(response, data)
//			} else {
//				// set session
//				session, _ := config.Store.Get(request, config.SESSION_ID)
//
//				session.Values["loggedIn"] = true
//				session.Values["name"] = user.Name
//				session.Values["username"] = user.Username
//
//				session.Save(request, response)
//
//				http.Redirect(response, request, "/employee", http.StatusSeeOther)
//			}
//		}
//
//	}
//
//}
//
//func Logout(response http.ResponseWriter, request *http.Request) {
//	session, _ := config.Store.Get(request, config.SESSION_ID)
//	// delete session
//	session.Options.MaxAge = -1
//	session.Save(request, response)
//
//	http.Redirect(response, request, "/Login", http.StatusSeeOther)
//}
//
//func Register(response http.ResponseWriter, request *http.Request) {
//
//	if request.Method == http.MethodGet {
//
//		temp, _ := template.ParseFiles("view/employee/Register.html")
//		temp.Execute(response, nil)
//
//	} else if request.Method == http.MethodPost {
//		// melakukan proses registrasi
//
//		// mengambil inputan form
//		request.ParseForm()
//
//		user := entity.User{
//			Name:      request.Form.Get("name"),
//			Username:  request.Form.Get("username"),
//			Password:  request.Form.Get("password"),
//			IsDeleted: 0,
//		}
//
//		errorMessages := validation.Struct(user)
//
//		if errorMessages != nil {
//
//			data := map[string]interface{}{
//				"validation": errorMessages,
//				"user":       user,
//			}
//
//			temp, _ := template.ParseFiles("view/employee/Register.html")
//			temp.Execute(response, data)
//		} else {
//
//			// hashPassword
//			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
//			user.Password = string(hashPassword)
//
//			// insert ke database
//			userModel.Create(user)
//
//			data := map[string]interface{}{
//				"pesan": "Registrasi berhasil",
//			}
//			temp, _ := template.ParseFiles("view/employee/Register.html")
//			temp.Execute(response, data)
//		}
//	}

//}
