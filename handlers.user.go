// handlers.user.go

package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func generateSessionToken() string {
	// We're usnig a random 16 character string as the session token
	// This is NOT a secure way of generating session tokens
	// DO NOT USE THIS IN PRODUCTION
	return strconv.FormatInt(rand.Int63(), 16)
}

func showRegistrationPage(c *gin.Context) {
	loggedInInterface, _ := c.Get("is_logged_in")
	is_logged_in := loggedInInterface.(bool)

	render(c, gin.H{
		"title":        "Register",
		"is_logged_in": is_logged_in}, "register.html")
}

func register(c *gin.Context) {
	// Obtain POSTed username and password values
	username := c.PostForm("username")
	password := c.PostForm("password")

	if _, err := registerNewUser(username, password); err == nil {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		loggedInInterface, _ := c.Get("is_logged_in")
		is_logged_in := loggedInInterface.(bool)

		render(c, gin.H{
			"title":        "Successful registration & Login",
			"is_logged_in": is_logged_in}, "login-successful.html")
	} else {
		// if the username/password combination is invalid,
		// show the error message on the login page
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error()})
	}

}

func showLoginPage(c *gin.Context) {
	loggedInInterface, _ := c.Get("is_logged_in")
	is_logged_in := loggedInInterface.(bool)

	render(c, gin.H{
		"title":        "Login",
		"is_logged_in": is_logged_in}, "login.html")
}

func performLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if isUserValid(username, password) {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		loggedInInterface, _ := c.Get("is_logged_in")
		is_logged_in := loggedInInterface.(bool)

		render(c, gin.H{
			"title":        "Successful Login",
			"is_logged_in": is_logged_in}, "login-successful.html")
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid Credentials Provided"})
	}
}

func logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)

	c.Redirect(http.StatusTemporaryRedirect, "/")
}
