package service

import (
	"jwt-auth/entity"
	"jwt-auth/repository"
	"time"
)

type JwtAuthService interface {
	CustomerDetail(username string) entity.Customer
	UpdateCustomerDetail(customer entity.Customer)
	GetVoucherPurchase(customerId uint64) []entity.VoucherPurchase
	GetVoucherPurchaseByCustomerAndVoucherGroup(customerId uint64, voucherGroupId uint64) entity.VoucherPurchase
	InsertVoucherPurchase(customerId uint64, voucherGroupId uint64, voucherCode string, purchaseDate time.Time)
	GetVoucherGroup(id uint64) entity.VoucherGroup
	GetVoucherGroupList() []entity.VoucherGroup
	UpdateVoucherGroup(voucherGroup entity.VoucherGroup)
}

type jwtAuthService struct {
	jwtAuthRepository repository.JwtAuthRepository
}

func NewJwtAuthService(jwtAuthRepo repository.JwtAuthRepository) JwtAuthService {
	return &jwtAuthService{
		jwtAuthRepository: jwtAuthRepo,
	}
}

func (service *jwtAuthService) CustomerDetail(username string) entity.Customer {
	return service.jwtAuthRepository.CustomerDetail(username)
}

func (service *jwtAuthService) GetVoucherPurchase(customerId uint64) []entity.VoucherPurchase {
	return service.jwtAuthRepository.GetVoucherPurchase(customerId)
}

func (service *jwtAuthService) GetVoucherPurchaseByCustomerAndVoucherGroup(customerId uint64, voucherGroupId uint64) entity.VoucherPurchase {
	return service.jwtAuthRepository.GetVoucherPurchaseByCustomerAndVoucherGroup(customerId, voucherGroupId)
}

func (service *jwtAuthService) InsertVoucherPurchase(customerId uint64, voucherGroupId uint64, voucherCode string, purchaseDate time.Time) {
	service.jwtAuthRepository.InsertVoucherPurchase(customerId, voucherGroupId, voucherCode, purchaseDate)
}

func (service *jwtAuthService) GetVoucherGroup(id uint64) entity.VoucherGroup {
	return service.jwtAuthRepository.GetVoucherGroup(id)
}

func (service *jwtAuthService) GetVoucherGroupList() []entity.VoucherGroup {
	return service.jwtAuthRepository.GetVoucherGroupList()
}

func (service *jwtAuthService) UpdateVoucherGroup(voucherGroup entity.VoucherGroup) {
	service.jwtAuthRepository.UpdateVoucherGroup(voucherGroup)
}

func (service *jwtAuthService) UpdateCustomerDetail(customer entity.Customer) {
	service.jwtAuthRepository.UpdateCustomer(customer)
}
