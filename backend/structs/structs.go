package structs

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	db.AutoMigrate(&Store{})
	db.AutoMigrate(&EmployeeForcast{})
}

type Store struct {
	ID         int    `json:"id" gorm:"primaryKey"`        // Unique identifier
	Name       string `json:"name" binding:"required"`     // User's name
	Location   string `json:"location" binding:"required"` // User's email
	PortalUser int
}

func (store Store) GetEmployees() []Employee {
	var employee []Employee
	db.Where(&Employee{StoreID: store.ID}).Find(&employee)
	return employee
}

func (store Store) CreateStore() {
	db.Create(&store)
}

func GetPortalUserStores(portalUser int) []Store {
	var userStores []Store
	db.Where(&Store{PortalUser: portalUser}).Find(&userStores)
	return userStores

}

func GetStore(storeId int) (*Store, string) {
	var store Store
	if err := db.First(&store, storeId).Error; err != nil {
		return nil, "User not found"

	}
	return &store, ""
}

type Employee struct {
	ID         int     `json:"id" gorm:"primaryKey"`           // Unique identifier
	Name       string  `json:"name" binding:"required"`        // User's name
	Email      string  `json:"email" binding:"required,email"` // User's email
	Age        int     `json:"age" binding:"required,min=0"`   // User's age
	StoreID    int     `json:"storeid" binding:"required,min=0"`
	HourlyRate float32 `json:"hourly_rate" binding:"requried"`
}

func GetEmployee(userId int) (*Employee, string) {
	var employee Employee
	if err := db.First(&employee, userId).Error; err != nil {
		return nil, "User not found"

	}
	return &employee, ""
}

func (user Employee) CreateEmployee() {
	db.Create(&user)
}

func (employee Employee) Schedule(day int, week int, year int) []EmployeeSchedule {
	storeId := employee.StoreID
	eId := employee.ID
	var sch []EmployeeSchedule

	base := db.Where(&EmployeeSchedule{StoreID: storeId, EmployeeID: eId})
	if day != 0 {
		base = base.Where("day = ?", day)
	}
	if week != 0 {
		base = base.Where("week = ?", week)
	}
	if year != 0 {
		base = base.Where("year = ?", year)
	}

	base.Order(clause.OrderBy{Columns: []clause.OrderByColumn{
		{Column: clause.Column{Name: "year"}, Desc: true},
		{Column: clause.Column{Name: "week"}, Desc: true},
		{Column: clause.Column{Name: "day"}, Desc: true},
	}})

	base.Find(&sch)

	return sch
}

func GetSchedules(storeId int, day int, week int, year int) []EmployeeSchedule {
	var sch []EmployeeSchedule

	base := db.Where(&EmployeeSchedule{StoreID: storeId})
	if day != 0 {
		base = base.Where("day = ?", day)
	}
	if week != 0 {
		base = base.Where("week = ?", week)
	}
	if year != 0 {
		base = base.Where("year = ?", year)
	}

	base.Order(clause.OrderBy{Columns: []clause.OrderByColumn{
		{Column: clause.Column{Name: "year"}, Desc: true},
		{Column: clause.Column{Name: "week"}, Desc: true},
		{Column: clause.Column{Name: "day"}, Desc: true},
	}})

	base.Find(&sch)

	return sch
}

func (employee Employee) SetSchedule(schs []EmployeeSchedule) {
	for _, sch := range schs {
		db.Create(&sch)
	}
}

type DayOfWeek int
type WeekOfYear int
type Year int

type EmployeeSchedule struct {
	StoreID    int        `json:"sid" `
	EmployeeID int        `json:"eid" `
	Day        DayOfWeek  `json:"day"`
	Week       WeekOfYear `json:"week"`
	Year       Year       `json:"year"`
	StartTime  int        `json:"starttime"`
	EndTime    int        `json:"endtime"`
}

type StoreSchedule struct {
	StoreId           int
	EmployeeSchedules []EmployeeSchedule
}

type EmployeeForcast struct {
	StoreId               int
	EmployeeId            int
	ForcastId             int
	TotalHours            int
	TotalRegularHours     int
	OverTimeHours         int
	SpreadOfPay           float32
	TotalRegularWage      float32
	TotalBaseWage         float32
	PayrollTaxEstimated   float32
	TotalForcastHourlyPay float32
	Salaries              float32
}

// Todo: get Forcast with a hash of storeid/week/year
func GenerateEmployeeForcast(eId int, eschls []EmployeeSchedule) {
	regularHoursinmins := 2400
	weeklyTotalMins := 0
	totalRegularMins := 0
	overTimeMin := 0
	spreadOfPayBase := 15
	for _, esch := range eschls {
		if esch.EmployeeID != eId {
			fmt.Printf("Employee IDs don't match: expected %d, got %+v\n", eId, esch)
			continue
		}
		endTimeInMins := esch.EndTime
		startTimeInMins := esch.StartTime
		totalMins := endTimeInMins - startTimeInMins
		weeklyTotalMins += totalMins

	}
	if weeklyTotalMins > totalRegularMins {
		totalRegularMins = regularHoursinmins
		overTimeMin = weeklyTotalMins - regularHoursinmins
	} else {
		totalRegularMins = weeklyTotalMins
	}

	employee, err := GetEmployee(eId)
	if err == "" {
		fmt.Print("here")
	}

	spreadOfPay := employee.HourlyRate * float32(spreadOfPayBase)
	toalRegularWage := employee.HourlyRate * float32(totalRegularMins)
	employeeForcast := EmployeeForcast{
		StoreId:           employee.StoreID,
		EmployeeId:        employee.ID,
		TotalHours:        weeklyTotalMins,
		TotalRegularHours: totalRegularMins,
		OverTimeHours:     overTimeMin,
		SpreadOfPay:       spreadOfPay,
		TotalRegularWage:  toalRegularWage,
	}

	fmt.Printf("Employee IDs %d total mins %d", eId, weeklyTotalMins)
	fmt.Printf("Employee %+v\n", employeeForcast)
}
