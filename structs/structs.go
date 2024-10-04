package structs

// User represents a user in the application
type User struct {
    ID    uint   `json:"id" binding:"required"` // Unique identifier
    Name  string `json:"name" binding:"required"` // User's name
    Email string `json:"email" binding:"required,email"` // User's email
    Age   int    `json:"age" binding:"required,min=0"` // User's age
}