package orm

type Todo struct {
	Message  string `json:"message"`
	CreateBy int    `json:"createBy" gorm:"not null"`
}