cat > main.go <<EOF
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

// User represents a user with a username and password
type User struct {
    Username string \`json:"username"\`
    Password string \`json:"password"\`
}

// hardcoded username and password
var validUsername = "admin"
var validPassword = "password123"

// loginHandler handles the login request
func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    // Check the username and password
    if user.Username == validUsername && user.Password == validPassword {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintln(w, "Login successful!")
    } else {
        w.WriteHeader(http.StatusUnauthorized)
        fmt.Fprintln(w, "Invalid username or password")
    }
}

func main() {
    http.HandleFunc("/login", loginHandler)

    fmt.Println("Starting server on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
EOF
