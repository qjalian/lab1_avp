package main

import (
	"fmt"
	"net/http"
)

var databaseUsers = map[string]string{
	"login1": "123",
	"login2": "rewq1",
}

func main() {
	http.HandleFunc("/", loginPageHandler)
	http.HandleFunc("/time", timePageHandler)

	fmt.Println("Server is running at http://localhost:5000")

	for login, password := range databaseUsers {
		fmt.Printf("Login - %s, Password - %s\n", login, password)
		break
	}

	http.ListenAndServe(":5000", nil)
}

func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	message := ""

	if r.Method == http.MethodPost {
		login := r.FormValue("login")
		password := r.FormValue("password")

		if storedPassword, ok := databaseUsers[login]; ok && storedPassword == password {
			http.Redirect(w, r, "/time", http.StatusSeeOther)
			return
		} else {
			message = "Wrong! Please try again"
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<title>Login</title>
			<style>
				body {
					display: flex;
					justify-content: center;
					align-items: flex-start;
					height: 100vh;
					margin: 0;
				}

				.container {
					margin: auto;
					text-align: center;
				}

				.login-form {
					margin-top: 20px;
				}

				.input-field {
					height: 28px;
					font-size:20px;
					width: 150px;
				}

				.custom-button {
					appearance: none;
					border: 0;
					border-radius: 8px;
					background: #26252D;
					color: #fff;
					padding: 10px 20px;
					font-size: 18px;
					cursor: pointer;
				}

				.custom-button:hover {
					background: #252850;
				}

				.custom-button:focus {
					 outline: none;
					 box-shadow: 0 0 0 2px 0.7 #252850;
				}

				.error-message {
					color: red;
					font-size: 20px;
					margin-top: 10px;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>Login</h1>
				<p class="error-message">%s</p>
				<form class="login-form" method="post" action="/">
					<input class="input-field" type="text" placeholder="login" id="login" name="login" required><br><br>
					<input class="input-field" type="password" placeholder="Password" id="password" name="password" required><br><br>
					<button class="custom-button" type="submit">Login</button>
				</form>
			</div>
		</body>
		</html>
	`, message)
}

func timePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<title>Time</title>
			<style>
				body {
					display: flex;
					justify-content: center;
					align-items: flex-start;
					height: 100vh;
					margin: 0;
				}

				.container {
					margin: auto;
					text-align: center;
				}

				.time-display {
					font-size: 24px;
					color: #333;
					margin-top: 20px;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<br>
				<p class="time-display">Current time: <span id="time-display"></span></p>

				<script>
				    function displayUpdatedTime() {
				        var displayElement = document.getElementById("time-display");
				        var updatedTime = new Date().toLocaleTimeString();
				        displayElement.textContent = updatedTime;
				    }

				    displayUpdatedTime();
				    setInterval(displayUpdatedTime, 1000);
				</script>
			</div>
		</body>
		</html>
	`)
}
