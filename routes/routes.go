package routes

import (
	"fmt"
	"managee/structs"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func string_to_int(stringInt string) (int, error) {
	i, err := strconv.Atoi(stringInt)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return 0, err
	}
	return i, nil
}

var db *gorm.DB

// SetupRoutes sets up the routes for the application
func SetupRoutes(r *gin.Engine) {
	connectionString := "host=localhost user=postgres password=postgrespw dbname=your_db port=32768 sslmode=disable"

	// Connect to the database
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed to connect to the database: %v", err)
	}
	db.AutoMigrate(&structs.Employee{})
	db.AutoMigrate(&structs.EmployeeSchedule{})

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
		//curl -X POST http://localhost:8080/employee -H "Content-Type: application/json" -d '{"name": "John Doe", "email": "john@example.com", "age": 24, "storeid": 4}'
		var user structs.Employee // Use the User struct
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&user)
		// c.JSON(http.StatusOK, gin.H{
		// 	"message": "Data received for " + user.Name,
		// 	"email":   user.Email,
		// 	"age":     user.Age,
		// })

	})
	r.GET("/employee/:e_id", func(c *gin.Context) {
		idParam := c.Param("e_id")
		var employee structs.Employee
		if err := db.First(&employee, idParam).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, employee)
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

	r.POST("/employee/:e_id/schedule", func(c *gin.Context) {
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

}
