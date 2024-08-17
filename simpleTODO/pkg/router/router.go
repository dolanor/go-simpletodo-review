package router

import (
	"fmt"
	"net/http"
	"strings"
	"todoproject/pkg/home"
	"todoproject/pkg/login"
	"todoproject/pkg/session"
	"todoproject/pkg/signup"
	"todoproject/pkg/task"
)

func RoutingHandler(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	if urlPath == "/" || urlPath == "/home" {
		home.HomePageHandler(w, r)
	} else if urlPath == "/signup" {
		signup.SignupPageHander(w, r)
	} else if urlPath == "/logout" {
		session.Logout(w, r)
	} else if urlPath == "/login" {
		login.LoginPageHandler(w, r)
	} else if urlPath == "/signupprocess" {
		signup.SignupProcessHandler(w, r)
	} else if urlPath == "/loginprocess" {
		login.LoginProcessHandler(w, r)
	} else if urlPath == "/createtaskprocess" {
		task.CreateTaskProcessor(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/deletetask/") {
		task.DeleteTaskProcessor(w, r)
	} else if urlPath == "/edittaskprocessor" {
		task.UpdateTaskProcessor(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/edittask/") {
		task.UpdateTaskPageHandler(w, r)
	} else {
		fmt.Fprintf(w, "Page not found")
	}

}
