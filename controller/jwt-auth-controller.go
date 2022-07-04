package controller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"crypto/md5"
	"encoding/hex"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"github.com/twinj/uuid"

	"jwt-auth/dto"
	"jwt-auth/entity"
	"jwt-auth/helper"
	"jwt-auth/service"
	"jwt-auth/utility"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

type JwtAuthController interface {
	Login(context *gin.Context)
	GetCustomerDetail(context *gin.Context)
	GetVoucherGroupList(context *gin.Context)
	PurchaseVoucher(context *gin.Context)
	RefreshToken(context *gin.Context)
}

type jwtAuthController struct {
	jwtAuthService service.JwtAuthService
}

func NewJwtAuthController(jwtAuthServ service.JwtAuthService) JwtAuthController {
	return &jwtAuthController{
		jwtAuthService: jwtAuthServ,
	}
}

// Login godoc
// @Summary Login
// @Schemes
// @Description Login
// @Accept json
// @Produce json
// @Param Login body dto.LoginDetail true "Login"
// @Success 200 {object} helper.Response{data=dto.CustomerLoginDTO}
// @Router /login [POST]
func (c *jwtAuthController) Login(context *gin.Context) {
	var loginDetail dto.LoginDetail
	context.BindJSON(&loginDetail)

	passByte := []byte(loginDetail.Password)
	passEncr := md5.Sum(passByte)

	pass := hex.EncodeToString(passEncr[:])

	logrus.WithFields(logrus.Fields{
		"username": loginDetail.Username,
		"password": pass,
	}).Info("Login")
	var customerDetail = c.jwtAuthService.CustomerDetail(loginDetail.Username)

	//compare the user from the request with data from db:
	if customerDetail.Username != loginDetail.Username || customerDetail.Password != pass {
		res := helper.BuildErrorResponse("Unauthorized", "Username atau Password salah", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	ts, err := CreateToken(customerDetail.Id, customerDetail.Username, customerDetail.Name, customerDetail.Email, customerDetail.Phone, customerDetail.Address)
	if err != nil {
		res := helper.BuildErrorResponse("Unable to process", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)

		return
	}

	var customerLogin dto.CustomerLoginDTO
	customerLogin.AccessUuid = ts.AccessUuid
	customerLogin.Id = customerDetail.Id
	customerLogin.Username = customerDetail.Username
	customerLogin.Name = customerDetail.Name
	customerLogin.Email = customerDetail.Email
	customerLogin.Phone = customerDetail.Phone
	customerLogin.Address = customerDetail.Address

	customerLogin.AccessToken = ts.AccessToken
	customerLogin.RefreshToken = ts.RefreshToken

	res := helper.BuildResponse(true, "OK", customerLogin)
	context.JSON(http.StatusOK, res)
}

// GetCustomerDetail godoc
// @Summary Get Customer Detail
// @Schemes
// @Description Get Customer Detail
// @Accept json
// @Produce json
// @Security BearerToken
// @Success 200 {object} dto.CustomerDetailDTO
// @Router /customer/detail [GET]
func (c *jwtAuthController) GetCustomerDetail(context *gin.Context) {
	tokenAuth, err := utility.ValidateJwtToken(context.Request)
	if err != nil {
		res := helper.BuildErrorResponse("Unauthorized", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var customerDetailEntity = c.jwtAuthService.CustomerDetail(tokenAuth.Username)

	var customerDetail dto.CustomerDetailDTO
	customerDetail.Id = customerDetailEntity.Id
	customerDetail.Username = customerDetailEntity.Username
	customerDetail.Name = customerDetailEntity.Name
	customerDetail.Email = customerDetailEntity.Email
	customerDetail.Phone = customerDetailEntity.Phone
	customerDetail.Address = customerDetailEntity.Address
	customerDetail.Point = customerDetailEntity.Point

	logrus.WithFields(logrus.Fields{
		"userId":   tokenAuth.UserId,
		"username": tokenAuth.Username,
	}).Info("GetCustomerDetail")

	resp := helper.BuildResponse(true, "OK", customerDetail)
	context.JSON(http.StatusOK, resp)
}

// GetVoucherGroup List godoc
// @Summary Voucher Group List
// @Schemes
// @Description Voucher Group List
// @Accept json
// @Produce json
// @Success 200 {object} []dto.VoucherGroupDTO
// @Router /voucher-group [GET]
func (c jwtAuthController) GetVoucherGroupList(context *gin.Context) {
	var voucherGroupList = c.jwtAuthService.GetVoucherGroupList()

	responses := make([]dto.VoucherGroupDTO, 0)
	for _, voucherGroup := range voucherGroupList {
		response := new(dto.VoucherGroupDTO)
		response.Id = voucherGroup.Id
		response.VoucherGroupName = voucherGroup.VoucherGroupName
		response.Qty = voucherGroup.Qty

		responses = append(responses, *response)
	}

	resp := helper.BuildResponse(true, "OK", responses)
	context.JSON(http.StatusOK, resp)
}

// PurchaseVoucher godoc
// @Summary Purchase Voucher
// @Schemes
// @Description Purchase Voucher
// @Accept json
// @Produce json
// @Security BearerToken
// @Param Finance body dto.PurchaseVoucherRequestDTO true "Voucher Group ID"
// @Success 200 {object} dto.VoucherPurchaseDTO
// @Router /voucher-purchase [POST]
func (c *jwtAuthController) PurchaseVoucher(context *gin.Context) {

	tokenAuth, err := utility.ValidateJwtToken(context.Request)
	if err != nil {
		res := helper.BuildErrorResponse("Unauthorized", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	var purchaseVoucherRequest dto.PurchaseVoucherRequestDTO
	context.BindJSON(&purchaseVoucherRequest)

	validate = helper.ValidateStruct()

	if err := validate.Struct(purchaseVoucherRequest); err != nil {
		logrus.Error(err)
		helper.ErrorValidation(err, context)
	}

	var existingVoucherPurchase = c.jwtAuthService.GetVoucherPurchaseByCustomerAndVoucherGroup(tokenAuth.UserId, purchaseVoucherRequest.VoucherGroupId)
	var checkVoucherPurchase entity.VoucherPurchase
	if existingVoucherPurchase != checkVoucherPurchase {
		res := helper.BuildErrorResponse("Bad Request", tokenAuth.Username+" already purchased this voucher!", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var voucherCode = time.Now().Format("2006-01-02-150405") + "V" + strconv.FormatUint(purchaseVoucherRequest.VoucherGroupId, 10) + strconv.FormatUint(tokenAuth.UserId, 10)
	c.jwtAuthService.InsertVoucherPurchase(tokenAuth.UserId, purchaseVoucherRequest.VoucherGroupId, voucherCode, time.Now())

	var customerDetailEntity = c.jwtAuthService.CustomerDetail(tokenAuth.Username)
	var point = customerDetailEntity.Point
	customerDetailEntity.Point = point - 1
	c.jwtAuthService.UpdateCustomerDetail(customerDetailEntity)

	var voucherGroup = c.jwtAuthService.GetVoucherGroup(purchaseVoucherRequest.VoucherGroupId)
	var qty = voucherGroup.Qty
	voucherGroup.Qty = qty - 1
	c.jwtAuthService.UpdateVoucherGroup(voucherGroup)

	var voucherPurchaseList = c.jwtAuthService.GetVoucherPurchase(customerDetailEntity.Id)

	responses := make([]dto.VoucherPurchaseDTO, 0)
	for _, voucherPurchase := range voucherPurchaseList {
		response := new(dto.VoucherPurchaseDTO)
		response.Id = voucherPurchase.Id
		response.CustomerId = customerDetailEntity.Id
		response.CustomerName = customerDetailEntity.Name
		response.VoucherGroupId = voucherPurchase.VoucherGroupId
		response.VoucherGroupName = voucherGroup.VoucherGroupName
		response.VoucherCode = voucherPurchase.VoucherCode

		responses = append(responses, *response)
	}

	logrus.WithFields(logrus.Fields{
		"userId":   tokenAuth.UserId,
		"username": tokenAuth.Username,
	}).Info("PurchaseVoucher")

	resp := helper.BuildResponse(true, "OK", responses)
	context.JSON(http.StatusOK, resp)
}

func CreateToken(userId uint64, username string, name string, email string, phone string, address string) (*entity.TokenDetails, error) {
	td := &entity.TokenDetails{}
	duration, err := strconv.Atoi(os.Getenv("TOKEN_DURATION"))
	if err != nil {
		duration = 30
	}

	td.AtExpires = time.Now().Add(time.Minute * time.Duration(duration)).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["accessUuid"] = td.AccessUuid
	atClaims["userId"] = userId
	atClaims["username"] = username
	atClaims["name"] = name
	atClaims["email"] = email
	atClaims["phone"] = phone
	atClaims["address"] = address
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refreshUuid"] = td.RefreshUuid
	rtClaims["username"] = username
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

// RefreshToken godoc
// @Summary Refresh Access Token
// @Schemes
// @Description Refresh Access Token
// @Accept json
// @Produce json
// @Param RefreshToken body dto.RefreshToken true "Refresh Token"
// @Success 200 {object} helper.Response{data=dto.UserLoginDTO}
// @Router /token/refresh [POST]
func (c *jwtAuthController) RefreshToken(context *gin.Context) {
	var requestToken dto.RefreshToken
	context.BindJSON(&requestToken)
	refreshToken := requestToken.Token

	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})

	//if there is an error, the token must have expired
	if err != nil {
		res := helper.BuildErrorResponse("Refresh token expired", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		res := helper.BuildErrorResponse("Unauthorized", "Token invalid", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		if !ok {
			res := helper.BuildErrorResponse("Unauthorized", "Token invalid", helper.EmptyObj{})
			context.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
		username := claims["username"].(string)

		//Create new pairs of refresh and access tokens
		var customerDetail = c.jwtAuthService.CustomerDetail(username)
		ts, createErr := CreateToken(customerDetail.Id, customerDetail.Username, customerDetail.Name, customerDetail.Email, customerDetail.Phone, customerDetail.Address)
		if createErr != nil {
			res := helper.BuildErrorResponse("Unable to process", createErr.Error(), helper.EmptyObj{})
			context.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		var customerLogin dto.CustomerLoginDTO
		customerLogin.AccessUuid = ts.AccessUuid
		customerLogin.Id = customerDetail.Id
		customerLogin.Username = customerDetail.Username
		customerLogin.Name = customerDetail.Name
		customerLogin.Email = customerDetail.Email
		customerLogin.Phone = customerDetail.Phone
		customerLogin.Address = customerDetail.Address

		customerLogin.AccessToken = ts.AccessToken
		customerLogin.RefreshToken = ts.RefreshToken

		logrus.WithFields(logrus.Fields{
			"userId":   customerLogin.Id,
			"username": customerLogin.Username,
		}).Info("RefreshToken")

		resp := helper.BuildResponse(true, "OK", customerLogin)
		context.JSON(http.StatusOK, resp)

	} else {
		res := helper.BuildErrorResponse("Unauthorized", "Token invalid", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	}
}
