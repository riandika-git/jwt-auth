package entity

type VoucherGroupTable interface {
	TableName() string
}

type VoucherGroup struct {
	Id               uint64 `gorm:"primary_key:auto_increment" json:"id"`
	VoucherGroupName string `json:"voucher_group_name"`
	Qty              int    `json:"qty"`
}

func (VoucherGroup) TableName() string {
	return "voucher_group"
}
