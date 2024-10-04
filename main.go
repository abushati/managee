package main

import (
    "github.com/gin-gonic/gin"
	"managee/routes"
)



func main() {
    r := gin.Default()
  
    // Setup the routes
    routes.SetupRoutes(r)

    // Start the server on port 8080
    r.Run(":8080")
}
