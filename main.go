package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var validCredentials = map[string]string{
	"username": "admin",
	"password": "password123",
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	// Check if the method is POST
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Validate credentials
		if username == validCredentials["username"] && password == validCredentials["password"] {
			http.Redirect(w, r, "/success", http.StatusFound)
			return
		} else {
			// If credentials are wrong, show an error message
			tmpl, err := template.ParseFiles("templates/login.html")
			if err != nil {
				http.Error(w, "Error loading page", http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, "Invalid username or password")
			return
		}
	}

	// Display the login page if method is GET
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func successPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login Successful!")
}

func main() {
	http.HandleFunc("/", loginPage)       // Handle login page
	http.HandleFunc("/success", successPage) // Handle success page

	// Start the web server
	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
