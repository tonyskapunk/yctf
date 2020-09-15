package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndex)

	// Handle flag0
	router.GET("/flag0", getFlag0)

	// Handle flag1
	router.GET("/flag1", getFlag1)

	// Handle flag2
	router.GET("/flag2", getFlag2)

	// Handle flag3
	router.GET("/flag3", getFlag3)

	// Handle flag4
	router.GET("/flag4", getFlag4)

	// Handle flag5
	router.GET("/flag5", getFlag5)

	// Handle flag6
	router.GET("/flag6", getFlag6)

	// Handle flag7
	router.GET("/flag7", getFlag7)
}
