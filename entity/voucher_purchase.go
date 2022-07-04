package entity

import "time"

type VoucherPurchaseTable interface {
	TableName() string
}

type VoucherPurchase struct {
	Id             uint64    `gorm:"primary_key:auto_increment" json:"id"`
	CustomerId     uint64    `json:"customer_id"`
	VoucherGroupId uint64    `json:"voucher_group_id"`
	VoucherCode    string    `json:"voucher_code"`
	PurchaseDate   time.Time `json:"purchase_date"`
}

func (VoucherPurchase) TableName() string {
	return "voucher_purchase"
}
