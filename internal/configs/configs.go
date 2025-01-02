package configs

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func GetPort() string {
	return os.Getenv("PORT")
}

func GetDBConnStr() string {
	return os.Getenv("DB_CONNECTION_STRING")
}

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}

type AdminUser struct {
	Name     string
	Email    string
	Password string
}

func GetAdminInfo() *AdminUser {
	adminUser := new(AdminUser)

	adminUser.Name = os.Getenv("ADMIN_NAME")
	adminUser.Email = os.Getenv("ADMIN_EMAIL")
	adminUser.Password = os.Getenv("ADMIN_PASSWORD")

	return adminUser
}

func GetAllowedOrigins() []string {
	return strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
}
