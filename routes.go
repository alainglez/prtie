// routes.go

package main

func initializeRoutes() {

	router.Use(setUserStatus())

	// Handle the index route
	router.GET("/", showIndexPage)

	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)

		userRoutes.POST("/register", ensureNotLoggedIn(), register)

		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)

		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)

		userRoutes.GET("/logout", ensureLoggedIn(), logout)
	}

	partyRoutes := router.Group("/party")
	{
		// Handle GET requests at /party/view/some_party_id
		partyRoutes.GET("/view/:party_id", getParty)

		partyRoutes.GET("/create", ensureLoggedIn(), showPartyCreationPage)

		partyRoutes.POST("/create", ensureLoggedIn(), createParty)
	}
}
