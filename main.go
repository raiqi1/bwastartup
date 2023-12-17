package main

import (
	"bwastart/auth"
	"bwastart/campaign"
	"bwastart/handler"
	"bwastart/helper"
	"bwastart/user"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "raiqi:difan12345@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)

	campaigns, err := campaignRepository.FindByUserID(1)
	fmt.Println("DEBUG")
	fmt.Println("DEBUG")
	fmt.Println("DEBUG")
	fmt.Println(len(campaigns))
	for _, campaign := range campaigns {
		fmt.Println(campaign.Name)
		if len(campaign.CampaignImages) > 0 {
			fmt.Println("jumlah gambar")
			fmt.Println(len(campaign.CampaignImages))
			fmt.Println(campaign.CampaignImages[0].FileName)
		}
	}

	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			repsonse := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, repsonse)
			return
		}
		//Bearer tokentokentoken
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			repsonse := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, repsonse)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			repsonse := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, repsonse)
			return
		}
		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(userID)
		if err != nil {
			repsonse := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, repsonse)
			return
		}
		c.Set("currentUser", user)
	}

}

// ambil nilai header Authorization
// dari header Authorization, kita ambil nilai token
// kita validasi token
// kita ambil user_id
// kita ambil user dari db berdasarkan user_id lewat service
// kita set context isinya user
