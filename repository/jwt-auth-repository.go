package repository

import (
	"jwt-auth/entity"
	"time"

	"gorm.io/gorm"
)

type JwtAuthRepository interface {
	CustomerDetail(username string) entity.Customer
	UpdateCustomer(customer entity.Customer)
	GetVoucherPurchase(customerId uint64) []entity.VoucherPurchase
	GetVoucherPurchaseByCustomerAndVoucherGroup(customerId uint64, voucherGroupId uint64) entity.VoucherPurchase
	InsertVoucherPurchase(customerId uint64, voucherGroupId uint64, voucherCode string, purchaseDate time.Time)
	GetVoucherGroup(id uint64) entity.VoucherGroup
	GetVoucherGroupList() []entity.VoucherGroup
	UpdateVoucherGroup(voucherGroup entity.VoucherGroup)
}

type jwtAuthConnection struct {
	connection *gorm.DB
}

//NewJwtAuthRepository creates an instance JwtAuthRepository
func NewJwtAuthRepository(dbConn *gorm.DB) JwtAuthRepository {
	return &jwtAuthConnection{
		connection: dbConn,
	}
}

func (db *jwtAuthConnection) CustomerDetail(username string) entity.Customer {
	var customer entity.Customer
	db.connection.Where("username = ?", username).First(&customer)

	return customer
}

func (db *jwtAuthConnection) GetVoucherPurchase(customerId uint64) []entity.VoucherPurchase {
	var voucherPurchase []entity.VoucherPurchase
	db.connection.Where("customer_id = ?", customerId).Find(&voucherPurchase)

	return voucherPurchase
}

func (db *jwtAuthConnection) GetVoucherPurchaseByCustomerAndVoucherGroup(customerId uint64, voucherGroupId uint64) entity.VoucherPurchase {
	var voucherPurchase entity.VoucherPurchase
	db.connection.Where("customer_id = ? AND voucher_group_id = ?", customerId, voucherGroupId).First(&voucherPurchase)

	return voucherPurchase
}

func (db *jwtAuthConnection) InsertVoucherPurchase(customerId uint64, voucherGroupId uint64, voucherCode string, purchaseDate time.Time) {
	voucherPurchase := entity.VoucherPurchase{CustomerId: customerId, VoucherGroupId: voucherGroupId, VoucherCode: voucherCode, PurchaseDate: purchaseDate}

	db.connection.Create(&voucherPurchase)
}

func (db *jwtAuthConnection) GetVoucherGroup(id uint64) entity.VoucherGroup {
	var voucherGroup entity.VoucherGroup
	db.connection.Where("id = ?", id).First(&voucherGroup)

	return voucherGroup
}

func (db *jwtAuthConnection) GetVoucherGroupList() []entity.VoucherGroup {
	var voucherGroup []entity.VoucherGroup
	db.connection.Find(&voucherGroup)

	return voucherGroup
}

func (db *jwtAuthConnection) UpdateVoucherGroup(voucherGroup entity.VoucherGroup) {
	voucherGroupEdit := entity.VoucherGroup{Id: voucherGroup.Id}
	db.connection.Model(&voucherGroupEdit).Updates(map[string]interface{}{"voucher_group_name": voucherGroup.VoucherGroupName, "qty": voucherGroup.Qty})
}

func (db *jwtAuthConnection) UpdateCustomer(customer entity.Customer) {
	editCustomer := entity.Customer{Id: customer.Id}

	db.connection.Model(&editCustomer).Updates(map[string]interface{}{"point": customer.Point})
}
