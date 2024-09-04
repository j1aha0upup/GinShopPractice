package model

type RoleAccess struct {
	Id       int `json:"id"`
	AccessId int `json:"access_id"`
	RoleId   int `json:"role_id"`
}

func (roleAccess *RoleAccess) TableName() string {
	return "role_access"
}
