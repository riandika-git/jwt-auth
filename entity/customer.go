package entity

type CustomerTable interface {
	TableName() string
}

type Customer struct {
	Id       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Point    int    `json:"point"`
}

func (Customer) TableName() string {
	return "customer"
}
