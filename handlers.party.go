// handlers.party.go

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	parties := getAllParties()

	loggedInInterface, _ := c.Get("is_logged_in")
	is_logged_in := loggedInInterface.(bool)

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":        "Home Page",
		"payload":      parties,
		"is_logged_in": is_logged_in}, "index.html")
}

func getParty(c *gin.Context) {
	// Check if the party ID is valid
	if partyID, err := strconv.Atoi(c.Param("party_id")); err == nil {
		// Chekc if party exists
		if party, err := getPartyByID(partyID); err == nil {
			loggedInInterface, _ := c.Get("is_logged_in")
			is_logged_in := loggedInInterface.(bool)

			// Call the render function with the name of the template to render
			render(c, gin.H{
				"title":        party.Title,
				"payload":      party,
				"is_logged_in": is_logged_in}, "party.html")
		} else {
			// If the party is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// If an invalid party ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func showPartyCreationPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Create New Party"}, "create-party.html")
}

func createParty(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	loggedInInterface, _ := c.Get("is_logged_in")
	is_logged_in := loggedInInterface.(bool)

	if a, err := createNewParty(title, content); err == nil {
		render(c, gin.H{
			"title":        "Submission Successfull",
			"content":      a,
			"is_logged_in": is_logged_in}, "submission-successful.html")
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
