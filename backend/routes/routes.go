package routes

import (
	"fmt"

	"managee/structs"
	"net/http"
	"strconv"

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
	r.POST("/store", func(c *gin.Context) {
		var store structs.Store // Use the User struct
		if err := c.ShouldBindJSON(&store); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//Todo: need to get the user thats logged in to save portaluser id to store
		store.CreateStore()
		store.PortalUser = 1
		c.JSON(http.StatusOK, store)
	})
	r.GET("/store", func(c *gin.Context) {
		//Todo: need to get the user thats logged in to save portaluser id to store
		PortalUser := 1
		stores := structs.GetPortalUserStores(PortalUser)
		c.JSON(http.StatusOK, stores)
	})

	r.GET("/store/:s_id", func(c *gin.Context) {
		idParam := c.Param("s_id")
		sidParam, _ := string_to_int(idParam)
		store, _ := structs.GetStore(sidParam)
		employees := store.GetEmployees()
		dynamicMap := make(map[string]interface{})
		dynamicMap["store"] = store
		dynamicMap["employees"] = employees
		c.JSON(http.StatusOK, dynamicMap)
	})

	r.POST("/employee", func(c *gin.Context) {
		var user structs.Employee // Use the User struct
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user.CreateEmployee()
	})

	r.GET("/employee/:e_id", func(c *gin.Context) {
		idParam := c.Param("e_id")
		sidParam, _ := string_to_int(idParam)

		employee, _ := structs.GetEmployee(sidParam)
		fmt.Printf("%+v", employee)
		c.JSON(http.StatusOK, employee)
	})

	r.GET("/employee/:e_id/schedule", func(c *gin.Context) {
		idParam := c.Param("e_id")

		dayQ, _ := string_to_int(c.DefaultQuery("day", "0"))
		weekQ, _ := string_to_int(c.DefaultQuery("week", "0"))
		yearQ, _ := string_to_int(c.DefaultQuery("year", "0"))
		idUint, err := string_to_int(idParam)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Error" + fmt.Sprintf("%+v", err),
			})
			return
		}

		user, _ := structs.GetEmployee(idUint)
		u_sch := user.Schedule(dayQ, weekQ, yearQ)

		c.JSON(http.StatusOK, u_sch)
	})

	r.POST("/employee/:e_id/schedule", func(c *gin.Context) {
		// body, err := ioutil.ReadAll(c.Request.Body)
		// if err != nil {
		// 	log.Println("Error reading body:", err)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		// 	return
		// }
		// fmt.Println("Request Body:", string(body))

		var schs []structs.EmployeeSchedule
		if err := c.ShouldBindJSON(&schs); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		idParam := c.Param("e_id")
		idUint, _ := string_to_int(idParam)
		user, _ := structs.GetEmployee(idUint)
		if user != nil {
			storeId := user.StoreID
			eId := user.ID
			for i := range schs {
				schs[i].StoreID = storeId
				schs[i].EmployeeID = eId
			}
		}
		user.SetSchedule(schs)
	})

	r.PUT("/employee/:e_id/schedule", func(c *gin.Context) {
		var schs []structs.EmployeeSchedule
		if err := c.ShouldBindJSON(&schs); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		idParam := c.Param("e_id")
		idUint, _ := string_to_int(idParam)
		user, _ := structs.GetEmployee(idUint)
		if user != nil {
			storeId := user.StoreID
			eId := user.ID
			for i := range schs {
				schs[i].StoreID = storeId
				schs[i].EmployeeID = eId
			}
		}

		user.SetSchedule(schs)

	})
}
