package model

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	LastIp   string `json:"last_ip"`
	Email    string `json:"email"`
	AddTime  int    `json:"add_time"`
	Status   int    `json:"status"`
}

func (user *User) TableName() string {
	return "User"
}
