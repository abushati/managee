package structs

// User represents a user in the application

type User struct {
    ID    int   `json:"id" binding:"required"` // Unique identifier
    Name  string `json:"name" binding:"required"` // User's name
    Email string `json:"email" binding:"required,email"` // User's email
    Age   int    `json:"age" binding:"required,min=0"` // User's age
}

func Get_user(userId int) User{
	user := User{
		ID: userId,
		Name: "test user",
		Email: "test@user.com",
		Age: 12,
	}
	return user
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

type UserSch struct {
	UserID int
	Day DayOfWeek
	Week WeekOfYear
	Year Year
	StartTime int
	EndTime int
}



