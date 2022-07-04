package main

import (
	"jwt-auth/config"
	"jwt-auth/controller"
	docs "jwt-auth/docs"
	"jwt-auth/repository"
	"jwt-auth/service"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	//docs "./docs"
)

var (
	db                *gorm.DB                     = config.SetupDatabaseConnection()
	jwtAuthRepository repository.JwtAuthRepository = repository.NewJwtAuthRepository(db)
	jwtAuthService    service.JwtAuthService       = service.NewJwtAuthService(jwtAuthRepository)
	jwtAuthController controller.JwtAuthController = controller.NewJwtAuthController(jwtAuthService)
)

// @title           Swagger API Notification
// @version         1.0
// @description     This is a microservice jwt-auth.
// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization
func main() {
	// open log file
	logPath := "jwt-auth.log"
	logPathFromEnv := os.Getenv("LOG_FILE_PATH")
	if logPathFromEnv != "" {
		logPath = logPathFromEnv + logPath
	}
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
		return
	}
	defer logFile.Close()

	Formatter := new(logrus.JSONFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"

	logrus.SetFormatter(Formatter)
	logrus.SetOutput(logFile)

	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	//cors
	var allowedOrigins string
	allowedOrigins = os.Getenv("CORS_ALLOWED_ORIGINS")
	logrus.Warn("allow cors origin:", allowedOrigins)
	var allowedOriginsArr []string
	for _, allowed := range strings.Split(allowedOrigins, ",") {
		allowedOriginsArr = append(allowedOriginsArr, allowed)
	}
	defaultConfig := cors.DefaultConfig()
	defaultConfig.AllowOrigins = allowedOriginsArr
	defaultConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(defaultConfig))

	r.POST("/login", jwtAuthController.Login)
	r.GET("/customer/detail", jwtAuthController.GetCustomerDetail)
	r.GET("/voucher-group", jwtAuthController.GetVoucherGroupList)
	r.POST("/voucher-purchase", jwtAuthController.PurchaseVoucher)
	r.POST("/token/refresh", jwtAuthController.RefreshToken)

	docs.SwaggerInfo.BasePath = os.Getenv("SWAGGER_BASE_PATH")
	v1 := r.Group("/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}

func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
