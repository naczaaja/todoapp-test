package orm

type User struct {
	Email  string `json:"email" binding:"required"`
}