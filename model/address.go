package model

type Address struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Uid            int    `json:"uid"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
	DefaultAddress int    `json:"default_address"`
	AddTime        int    `json:"add_time"`
}

func (address *Address) TableName() string {
	return "address"
}
