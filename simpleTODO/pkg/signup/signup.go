package signup

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/session"
	"todoproject/pkg/tools"
	"todoproject/pkg/user"
)

var csrft string

func SignupPageHander(w http.ResponseWriter, r *http.Request) {
	csrft = tools.GenerateUUID()
	http.SetCookie(w, &http.Cookie{Name: "csrft", Value: csrft, HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode})
	template, _ := template.ParseFiles("../../pkg/signup/template/signup.html")
	template.Execute(w, nil)
}

func SignupProcessHandler(w http.ResponseWriter, r *http.Request) {
	sent_csrf_token, err := r.Cookie("csrft")
	if err != nil || sent_csrf_token == nil {
		fmt.Println("csrft not found")
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	} else {
		if sent_csrf_token.Value != csrft {
			fmt.Println("invalid csrft")
			http.Redirect(w, r, "/signup", http.StatusSeeOther)
		} else {
			_, err := r.Cookie("session_id")
			if err != nil {
				fmt.Println("we reached here 1")
				userId := tools.GenerateUUID()
				fmt.Println("we reached here 2")
				sessionId := tools.GenerateUUID()
				fmt.Println("we reached here 3")
				http.SetCookie(w, &http.Cookie{Name: "session_id", Value: sessionId, Expires: time.Now().Add(time.Hour * 168)})
				fmt.Println("we reached here 4")
				query, arguments := databasetools.QuerryMaker("insert", []string{"userID", "username", "password", "firstName", "lastName", "email", "phoneNumber"}, "users", map[string]string{}, map[string]string{"userId": userId, "username": r.Form.Get("username"), "password": tools.HashThis(r.Form.Get("password")), "firstName": r.Form.Get("firstName"), "lastName": r.Form.Get("lastName"), "email": r.Form.Get("email"), "phoneNumber": r.Form.Get("phoneNumber")})
				fmt.Println("we reached here 5")
				fmt.Println("*******************************************")
				fmt.Println(query)
				fmt.Println("*******************************************")
				fmt.Println(arguments...)
				fmt.Println("*******************************************")
				user.CreateUser(databasetools.DataBase, query, arguments)
				fmt.Println("we reached here 6")
				session.CreateSession(databasetools.DataBase, query)
				fmt.Println("we reached here 7")
				http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
			} else {
				http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
				http.Redirect(w, r, "/home", http.StatusSeeOther)
			}
			http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
			http.Redirect(w, r, "/home", http.StatusSeeOther)
		}
	}

}
