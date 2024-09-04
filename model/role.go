package model

type Role struct {
	Id          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Status      int      `json:"status"`
	AddTime     int      `json:"add_time"`
	AccessList  []Access `gorm:"many2many:role_access;"`
}

func (role *Role) TableName() string {
	return "role"
}
