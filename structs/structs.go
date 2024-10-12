package structs

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {

	// User represents a user in the application
	var connectionString = "host=localhost user=postgres password=postgrespw dbname=your_db port=32770 sslmode=disable"

	// Connect to the database
	var err error
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed to connect to the database: %v", err)
		return
	}
	db.AutoMigrate(&Employee{})
	db.AutoMigrate(&EmployeeSchedule{})
}

type Employee struct {
	ID      int    `json:"id" gorm:"primaryKey"`           // Unique identifier
	Name    string `json:"name" binding:"required"`        // User's name
	Email   string `json:"email" binding:"required,email"` // User's email
	Age     int    `json:"age" binding:"required,min=0"`   // User's age
	StoreID int    `json:"storeid" binding:"required,min=0"`
}

func (user Employee) CreateEmployee() {
	db.Create(&user)
}

func GetEmployee(userId int) (*Employee, string) {
	var employee Employee
	if err := db.First(&employee, userId).Error; err != nil {
		return nil, "User not found"

	}
	return &employee, ""
}

// func (employee Employee) GetSchedule() EmployeeSchedule {
// 	storeId := employee.StoreID
// 	eId := employee.ID
// 	return EmployeeSchedule{StoreID: storeId, EmployeeID: eId, Day: Sunday, Week: 3, Year: 2024, StartTime: 1243, EndTime: 2345}
// }

type DayOfWeek int
type WeekOfYear int
type Year int

const (
	Saturday DayOfWeek = iota
	Sunday
	Monday
	Tuesday
	Wednesday
	Thrusday
	Friday
)

type DaySchedule struct {
	Day       DayOfWeek
	Week      WeekOfYear
	Year      Year
	StartTime int
	EndTime   int
}

type EmployeeSchedule struct {
	StoreID    int
	EmployeeID int `gorm:"primaryKey"`
	Schedule   []DaySchedule
}

type StoreSchedule struct {
	StoreId          int
	EmployeSchedules []EmployeeSchedule
}
