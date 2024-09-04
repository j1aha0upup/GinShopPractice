package model

type Order struct {
	Id          int
	Uid         int
	OrderSn     string
	TotalPrice  float64
	Name        string
	Phone       string
	Address     string
	PayStatus   int
	PayType     int
	OrderStatus int
	AddTime     int
	OrderItem   []OrderItem `gorm:"foreignKey:OrderId"`
}

func (order *Order) TableName() string {
	return "order"
}
