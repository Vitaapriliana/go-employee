package main

import (
	"go-web-employee/config"
	"go-web-employee/controller/dashboard"
	"log"
	"net/http"
)

func main() {

	config.ConnectDB()

	// Login Rooute
	//http.HandleFunc("/login", authController.Login)
	//http.HandleFunc("/logout", authController.Logout)
	//http.HandleFunc("/register", authController.Register)

	//Employee Route
	http.HandleFunc("/", dashboard.Index)
	http.HandleFunc("/add", dashboard.Add)
	http.HandleFunc("/edit", dashboard.Edit)
	http.HandleFunc("/delete", dashboard.Delete)

	log.Println("Server Running on Port 8080")
	http.ListenAndServe(":8080", nil)

}
