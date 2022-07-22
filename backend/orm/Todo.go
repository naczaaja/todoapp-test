package orm

type Todo struct {
	Id       int    `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	Message  string `json:"message"`
	CreateBy int    `json:"createBy" gorm:"not null"`
}