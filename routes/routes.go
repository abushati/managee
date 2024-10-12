package routes

import (
	"fmt"

	"managee/structs"
	"net/http"
	"strconv"
	"time"

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
	r.POST("/employee", func(c *gin.Context) {
		//curl -X POST http://localhost:8080/employee -H "Content-Type: application/json" -d '{"name": "John Doe", "email": "john@example.com", "age": 24, "storeid": 4}'
		var user structs.Employee // Use the User struct
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user.CreateEmployee()
		// c.JSON(http.StatusOK, gin.H{
	})
	r.GET("/employee/:e_id", func(c *gin.Context) {
		idParam := c.Param("e_id")
		sidParam, _ := string_to_int(idParam)
		// employee, _ := structs.GetEmployee(sidParam)
		employee, _ := structs.GetEmployee(sidParam)
		fmt.Printf("%+v", employee)
		c.JSON(http.StatusOK, employee)
	})

	r.GET("/employee/:e_id/schedule", func(c *gin.Context) {
		idParam := c.Param("e_id")
		currentTime := time.Now()
		year, week := currentTime.ISOWeek()

		dayQ := c.DefaultQuery("day", strconv.Itoa(int(currentTime.Day())))
		weekQ := c.DefaultQuery("week", strconv.Itoa(week))
		yearQ := c.DefaultQuery("year", strconv.Itoa(year))
		fmt.Printf("yeaer %+v", yearQ)
		// c.JSON(http.StatusOK, gin.H{
		// 	"message": "Error" + fmt.Sprintf("%+v", yearQ) + fmt.Sprintf("%+v", weekQ) + fmt.Sprintf("%+v", dayQ) + idParam,
		// })
		// return

		idUint, err := string_to_int(idParam)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Error" + fmt.Sprintf("%+v", err),
			})
			return
		}

		// 	user, _ := structs.GetEmployee(idUint)

		// 	u_sch := user.GetSchedule(day, week, year)

		// 	c.JSON(http.StatusOK, gin.H{
		// 		"message": "Data received for " + fmt.Sprintf("%+v", u_sch),
		// 	})
	})

	r.POST("/employee/:e_id/schedule", func(c *gin.Context) {
		// idParam := c.Param("e_id")
		// idUint, err := string_to_int(idParam)
		// if err != nil {
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"message": "Error" + fmt.Sprintf("%+v", err),
		// 	})
		// 	return
		// }

		// user, _ := structs.GetEmployee(idUint)
		// u_sch := user.GetSchedule()

		// c.JSON(http.StatusOK, gin.H{
		// 	"message": "Data received for " + fmt.Sprintf("%+v", u_sch),
		// })
	})

}
