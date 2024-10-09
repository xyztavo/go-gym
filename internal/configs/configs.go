package configs

import (
	"fmt"
	"log"
	"os"

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
	return fmt.Sprintf("postgres://%v:%v@localhost/%v?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
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
