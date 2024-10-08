package structs

// User represents a user in the application

type Employee struct {
	ID      int    `json:"id" binding:"required"`          // Unique identifier
	Name    string `json:"name" binding:"required"`        // User's name
	Email   string `json:"email" binding:"required,email"` // User's email
	Age     int    `json:"age" binding:"required,min=0"`   // User's age
	StoreID int
}

func GetEmployee(eId int) Employee {
	user := Employee{
		ID:    eId,
		Name:  "test user",
		Email: "test@user.com",
		Age:   12,
	}
	return user
}

func (employee Employee) GetSchedule() EmployeeSchedule {
	storeId := employee.StoreID
	eId := employee.ID
	return EmployeeSchedule{StoreID: storeId, EmployeeID: eId, Day: Sunday, Week: 3, Year: 2024, StartTime: 1243, EndTime: 2345}
}

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

type EmployeeSchedule struct {
	StoreID    int
	EmployeeID int
	Day        DayOfWeek
	Week       WeekOfYear
	Year       Year
	StartTime  int
	EndTime    int
}
