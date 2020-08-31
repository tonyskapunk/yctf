package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndex)

	// Handle GET flag0
	router.GET("/flag0", getFlag0)

	// Handle GET flag1
	router.GET("/flag1", getFlag1)

	// Handle GET flag2
	router.GET("/flag2", getFlag2)

	// Handle GET flag3
	router.GET("/flag3", getFlag3)
}
