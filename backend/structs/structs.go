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
	db.AutoMigrate(&StoreForcast{})
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

type Position int

const (
	HeadChef Position = iota //BOH
	SousChef
	LineCook
	PrepCook
	PastryChef
	Dishwasher
	Expediter
	GeneralManager //FOH
	AssistantManager
	Host
	Server
	Busser
	Bastender
	Barback
)

var BackOfHousePositions = []Position{HeadChef, SousChef, LineCook, PrepCook, PastryChef, Dishwasher, Expediter}

type Employee struct {
	ID      int    `json:"id" gorm:"primaryKey"`           // Unique identifier
	Name    string `json:"name" binding:"required"`        // User's name
	Email   string `json:"email" binding:"required,email"` // User's email
	Age     int    `json:"age" binding:"required,min=0"`   // User's age
	StoreID int    `json:"storeid" binding:"required,min=0"`
	//Todo: validated that its hourly or salary
	CompensationType string   `json:"compensation_type" binding:"required,min=0"`
	HourlyRate       float64  `json:"hourly_rate" default:"0"`
	Salary           float64  `json:"salary" default:"0"`
	Position         Position `json:"position"`
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

type EmployeeSchedule struct {
	StoreID    int `json:"sid" `
	EmployeeID int `json:"eid" `
	Day        int `json:"day"`
	Week       int `json:"week"`
	Year       int `json:"year"`
	StartTime  int `json:"starttime"`
	EndTime    int `json:"endtime"`
}

type StoreSchedule struct {
	StoreId           int
	EmployeeSchedules []EmployeeSchedule
}

type EmployeeForcast struct {
	YearWeek                  string
	StoreId                   int
	EmployeeId                int
	ForcastId                 int `gorm:"primaryKey;autoIncrement"`
	TotalHours                float64
	TotalRegularHours         float64
	OverTimeHours             float64
	SpreadOfPay               float64
	TotalRegularWage          float64
	OvertimeWage              float64
	TotalBaseWage             float64
	PayrollTaxEstimated       float64
	GrandTotalPayrollForecast float64
	Salary                    float64
}

func (ef EmployeeForcast) save() {
	db.Create(&ef)
}

type StoreForcast struct {
	StoreId                              int
	ForcastId                            int `gorm:"primaryKey;autoIncrement"`
	YearWeek                             string
	FrontHouseTotalHours                 float64
	FrontHouseTotalRegHours              float64
	FrontHouseOverTimeHours              float64
	FrontHouseSpreadofPay                float64
	FrontHouseTotalRegWage               float64
	FrontHouseOvertimeWage               float64
	FrontHouseTotalBaseWage              float64
	FrontHousePayrollTaxEstimate         float64
	FrontHouseTotalForecastHourlyPayRate float64
	FrontHouseSalary                     float64
	BackHouseTotalHours                  float64
	BackHouseTotalRegHours               float64
	BackHouseOverTimeHours               float64
	BackHouseSpreadofPay                 float64
	BackHouseTotalRegWage                float64
	BackHouseOvertimeWage                float64
	BackHouseTotalBaseWage               float64
	BackHousePayrollTaxEstimate          float64
	BackHouseTotalForecastHourlyPayRate  float64
	BackHouseSalary                      float64
}

func minsToHoursConverter(mins int) float64 {
	return float64(mins) / 60
}

// Todo: get Forcast with a hash of storeid/week/year
func GenerateEmployeeForcast(done chan EmployeeForcast, employeeId int, employeeSchedules []EmployeeSchedule, year int, week int) (*EmployeeForcast, string) {
	regularHoursinmins := 2400
	weeklyTotalMins := 0
	totalRegularMins := 0
	overTimeMin := 0
	spreadOfPayBase := 15.0

	employee, err := GetEmployee(employeeId)
	if err != "" {
		fmt.Printf("Employee IDs %d don't exist\n", employeeId)
		return nil, "employee nonexist"
	}

	for _, dailySchedule := range employeeSchedules {
		if dailySchedule.EmployeeID != employeeId {
			fmt.Printf("Employee IDs don't match: expected %d, got %+v\n", employeeId, dailySchedule)
			continue
		} else if dailySchedule.Week != week || dailySchedule.Year != year {
			fmt.Printf("Schedule week/year doesn't match expected week %d, year %d got schedule with week %d year %d \n", week, year, dailySchedule.Week, dailySchedule.Year)
			continue
		}

		endTimeInMins := dailySchedule.EndTime
		startTimeInMins := dailySchedule.StartTime
		totalMins := endTimeInMins - startTimeInMins
		weeklyTotalMins += totalMins
	}

	if weeklyTotalMins > regularHoursinmins {
		totalRegularMins = regularHoursinmins
		overTimeMin = weeklyTotalMins - regularHoursinmins
	} else {
		totalRegularMins = weeklyTotalMins
	}
	totalHours := minsToHoursConverter(weeklyTotalMins)
	totalRegularHours := minsToHoursConverter(totalRegularMins)
	overTimeHours := minsToHoursConverter(overTimeMin)

	spreadOfPay := employee.HourlyRate * spreadOfPayBase
	toalRegularWage := employee.HourlyRate * totalRegularHours
	overtimeWage := employee.HourlyRate * 1.5 * overTimeHours
	totalBaseWage := spreadOfPay + toalRegularWage + overtimeWage

	payrollTaxEstimated := totalBaseWage * .1
	grandTotalPayrollForecast := totalBaseWage + payrollTaxEstimated

	employeeForcast := EmployeeForcast{
		YearWeek:                  fmt.Sprintf("%d/%d", year, week),
		StoreId:                   employee.StoreID,
		EmployeeId:                employee.ID,
		TotalHours:                totalHours,
		TotalRegularHours:         totalRegularHours,
		OverTimeHours:             overTimeHours,
		SpreadOfPay:               spreadOfPay,
		TotalRegularWage:          toalRegularWage,
		OvertimeWage:              overtimeWage,
		TotalBaseWage:             totalBaseWage,
		PayrollTaxEstimated:       payrollTaxEstimated,
		GrandTotalPayrollForecast: grandTotalPayrollForecast,
		Salary:                    employee.Salary,
	}
	employeeForcast.save()
	done <- employeeForcast
	fmt.Printf("Employee IDs %d total mins %d", employeeId, weeklyTotalMins)
	fmt.Printf("EmployeeForcast %+v\n", employeeForcast)
	return &employeeForcast, ""
}

func GenerateStoreForcast(employeeForcasts []EmployeeForcast) {

	for forcast := range employeeForcasts {
		employee, err := GetEmployee(forcast)
		if err != "" {
			fmt.Print("bad)")
		}
		if contains(BackOfHousePositions, employee.Position) {

		}
	}

}
