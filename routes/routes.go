package routes

import (
	"fmt"
    "net/http"
    "strings"
	"strconv"
    "github.com/gin-gonic/gin"
	"managee/structs"
)

// SetupRoutes sets up the routes for the application
func SetupRoutes(r *gin.Engine) {
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, world!",
            "upper":   strings.ToUpper("hello"),
        })
    })

    r.GET("/greet", func(c *gin.Context) {
        name := c.Query("name")
        if name == "" {
            name = "stranger"
        }
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, " + name + "!",
        })
    })

	r.POST("/employee", func(c *gin.Context) {
        var user structs.User // Use the User struct
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{
            "message": "Data received for " + user.Name,
            "email":   user.Email,
            "age":     user.Age,
        })
    })

	r.GET("/employee/:e_id/schedule", func(c *gin.Context) {
		idParam := c.Param("e_id")
		idUint, _ := strconv.Atoi(idParam)
		user := structs.Get_user(idUint)

		u_sch := structs.UserSch{UserID: user.ID, Day: structs.Saturday, Week: 4, Year: 2024, StartTime: 12123432, EndTime: 12123432}

        c.JSON(http.StatusOK, gin.H{
            "message": "Data received for " + fmt.Sprintf("%+v", u_sch),
        })
    })

	r.GET("/employee/:e_id", func(c *gin.Context) {
		idParam := c.Param("e_id")
		c.JSON(http.StatusOK, gin.H{
            "message": idParam,
        })
	})
}
