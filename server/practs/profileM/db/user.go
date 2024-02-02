package db

type User struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	Logname  string `json:"logname"`
	Password string `json:"password"`
}
