package dto

import "time"

type LoginDetail struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomerLoginDTO struct {
	AccessUuid   string `json:"accessUuid"`
	Id           uint64 `json:"id"`
	Username     string `json:"username"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type CustomerDetailDTO struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Point    int    `json:"point"`
}

type VoucherGroupDTO struct {
	Id               uint64 `json:"id"`
	VoucherGroupName string `json:"voucherGroupName"`
	Qty              int    `json:"qty"`
}

type PurchaseVoucherRequestDTO struct {
	VoucherGroupId uint64 `json:"voucherGroupId"`
}

type VoucherPurchaseDTO struct {
	Id               uint64    `json:"id"`
	CustomerId       uint64    `json:"customerId"`
	CustomerName     string    `json:"customerName"`
	VoucherGroupId   uint64    `json:"voucherGroupId"`
	VoucherGroupName string    `json:"voucherGroupName"`
	VoucherCode      string    `json:"voucherCode"`
	PurchaseDate     time.Time `json:"purchaseDate"`
}

type RefreshToken struct {
	Token string `json:"refreshToken"`
}

type FcmTokenDTO struct {
	FcmToken string `json:"fcmToken"`
}

type UserLoginDTO struct {
	AccessUuid   string  `json:"accessUuid"`
	UserId       uint64  `json:"userId"`
	Username     string  `json:"username"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Phone        string  `json:"phone"`
	Whatsapp     string  `json:"whatsapp"`
	Address      string  `json:"address"`
	Deposit      float64 `json:"deposit"`
	FirstLogin   bool    `json:"firstLogin"`
	AccessToken  string  `json:"accessToken"`
	RefreshToken string  `json:"refreshToken"`
}

type UserProfileDTO struct {
	UserId   uint64 `json:"userId"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Whatsapp string `json:"whatsapp"`
	Address  string `json:"address"`
	FcmToken string `json:"fcmToken"`
}

type UserPasswordDTO struct {
	OldPassword       string `json:"oldPassword"`
	NewPassword       string `json:"newPassword"`
	NewPasswordRetype string `json:"newPasswordRetype"`
}

type ProfileAccountDTO struct {
	Name     string `json:"name"  validate:"required,min=2,max=200"`
	Email    string `json:"email"  validate:"required,min=2,max=200"`
	Phone    string `json:"phone"  validate:"required,min=10,max=15"`
	Whatsapp string `json:"whatsapp"  validate:"required,min=10,max=15"`
}

type ProfileFinanceDTO struct {
	Bank        string `json:"bank"  validate:"required,min=2,max=50"`
	AccountNo   string `json:"accountNo"  validate:"required,min=10,max=20"`
	AccountName string `json:"accountName"  validate:"required,min=2,max=150"`
}

type UserIdDTO struct {
	UserId uint64 `json:"userId"`
}

type CmsLoginDTO struct {
	AccessUuid   string `json:"accessUuid"`
	UserId       uint64 `json:"userId"`
	Username     string `json:"username"`
	Name         string `json:"name"`
	Whatsapp     string `json:"whatsapp"`
	Level        string `json:"level"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type CmsProfileDTO struct {
	UserId   uint64 `json:"userId"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Whatsapp string `json:"whatsapp"`
	Level    string `json:"level"`
}

type DepositDTO struct {
	Id                uint64    `json:"id"`
	TransactionDate   time.Time `json:"transactionDate"`
	TransactionType   string    `json:"transactionType"`
	TransactionAmount float64   `json:"transactionAmount"`
}
