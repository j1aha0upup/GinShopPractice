package model

type AdminUser struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
	RoleId   int    `json:"role_id"`
	AddTime  int    `json:"add_time"`
	IsSuper  bool   `json:"is_supper"`
	ActRole  Role   `json:"role" gorm:"foreignKey:RoleId"`
}

func (user *AdminUser) TableName() string {
	return "admin_user"
}
