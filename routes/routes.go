package routes

import (
	"fmt"
	"managee/structs"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func string_to_int(stringInt string) (int, error) {
	i, err := strconv.Atoi(stringInt)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return 0, err
	}
	return i, nil
}

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
		var user structs.Employee // Use the User struct
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
		// dayQ := c.Query("day")
		// weekQ := c.Query("week")
		// yearQ := c.Query("year")
		// c.JSON(http.StatusOK, gin.H{
		// 	"message": "Error" + fmt.Sprintf("%+v", yearQ) + fmt.Sprintf("%+v", weekQ) + fmt.Sprintf("%+v", dayQ),
		// })
		// return

		idUint, err := string_to_int(idParam)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Error" + fmt.Sprintf("%+v", err),
			})
			return
		}

		user := structs.GetEmployee(idUint)
		u_sch := user.GetSchedule()

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
