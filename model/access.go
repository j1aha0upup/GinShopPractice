package model

type Access struct {
	Id          int      `json:"id"`
	ModuleName  string   `json:"module_name"`
	ActionName  string   `json:"action_name"`
	Type        int      `json:"type"` //1：模块 2：菜单 3：动作
	Url         string   `json:"url"`
	Status      int      `json:"status"`
	ModuleId    int      `json:"module_id" `
	Sort        int      `json:"sort"`
	Description string   `json:"description"`
	AddTime     int      `json:"add_time"`
	AccessList  []Access `gorm:"foreignkey:ModuleId"`
	RoleList    []Role   `gorm:"many2many:role_access;"`
	IsChecked   bool     `gorm:"-"`
}

func (access *Access) TableName() string {
	return "access"
}
