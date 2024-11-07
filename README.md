cat > README.md <<EOF
# My Auth App

This is a simple Go web application for user authentication. It checks a username and password against hardcoded values and responds with login success or failure.

## Project Structure

- **main.go**: The main Go file containing the server and login logic.
- **go.mod**: Go module file.
- **.gitignore**: Lists files and directories that Git should ignore.
- **README.md**: Project documentation.

## Setup and Run

1. Clone the repository:
   \`\`\`
   git clone https://github.com/yourusername/my-auth-app.git
   cd my-auth-app
   \`\`\`

2. Initialize Go dependencies:
   \`\`\`
   go mod tidy
   \`\`\`

3. Run the server:
   \`\`\`
   go run main.go
   \`\`\`

The server will start on \`localhost:8080\`. Send a POST request to \`/login\` with JSON containing the username and password.

## Test the API

Use the following \`curl\` command to test:
\`\`\`
curl -X POST -d '{"username":"admin", "password":"password123"}' -H "Content-Type: application/json" http://localhost:8080/login
\`\`\`

- Valid credentials will return \`Login successful!\`.
- Invalid credentials will return \`Invalid username or password\`.
EOF
