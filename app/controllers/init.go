package controllers

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/pdmp/amai/app/models"

	jwt "github.com/dgrijalva/jwt-go"
	gormdb "github.com/revel/modules/orm/gorm/app"
	"github.com/revel/revel"
)

var (
	errAuthHeaderNotFound = errors.New("authorization header not found")
	errInvalidTokenFormat = errors.New("token format is invalid")
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Println(msg)
	}
}

func init() {
	revel.OnAppStart(InitDB)
}
func AddLog(c *revel.Controller) revel.Result {
	log.Println("InterceptFunc Test.")
	return nil
}

// In order to valid the user.
func Authenticate(c *revel.Controller) revel.Result {
	log.Println("Authenticate!")
	log.Println(c.Params)

	tokenString, err := getTokenString(c)
	if err != nil {
		log.Println("get token string failed")
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON("get token string failed")
	}

	var claims jwt.MapClaims
	claims, err = decodeToken(tokenString)
	if err != nil {
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON("auth failed")
	}
	log.Println("claims decode:", claims)
	log.Println(claims["email"])
	email, found := claims["email"]
	if !found {
		log.Println(err)
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON("email not found in db")
	}

	log.Println("email found:", email)
	var user models.User
	err = gormdb.DB.Where("user_email = ?", email).First(&user).Error

	if err != nil {
		log.Println(err)
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON("auth failed")
	}

	log.Println("auth token success")
	return nil
}

func AuthenticateAdmin(c *revel.Controller) revel.Result {
	log.Println("Authenticate!")
	log.Println(c.Params)

	tokenString, err := getTokenString(c)
	if err != nil {
		log.Println("get token string failed")
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON("get token string failed")
	}

	var claims jwt.MapClaims
	claims, err = decodeToken(tokenString)
	if err != nil {
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON("auth failed")
	}
	log.Println("claims decode:", claims)
	log.Println(claims["email"])
	email, found := claims["email"]
	if !found {
		log.Println(err)
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON("email not found in db")
	}

	log.Println("email found:", email)
	var user models.User
	err = gormdb.DB.Where("user_email = ? AND id_role_user = 1", email).First(&user).Error

	if err != nil {
		log.Println(err)
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON("auth failed, need permission")
	}

	return nil
}

func getTokenString(c *revel.Controller) (tokenString string, err error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		log.Println(errAuthHeaderNotFound)
		return "", errAuthHeaderNotFound
	}

	tokenSlice := strings.Split(authHeader, " ")
	if len(tokenSlice) != 2 {
		return "", errInvalidTokenFormat
	}
	tokenString = tokenSlice[1]
	return tokenString, nil

}
func InitDB() {
	revel.INFO.Println("Doing DB Migrations...")

	gormdb.DB.AutoMigrate(&models.AcademicLevel_User{})
	gormdb.DB.AutoMigrate(&models.Assistance{})
	gormdb.DB.AutoMigrate(&models.Category{})
	gormdb.DB.AutoMigrate(&models.Classroom{})
	gormdb.DB.AutoMigrate(&models.Gender_User{})
	gormdb.DB.AutoMigrate(&models.PhonexUser{})
	gormdb.DB.AutoMigrate(&models.Place_Session{})
	gormdb.DB.AutoMigrate(&models.Role_User{})
	gormdb.DB.AutoMigrate(&models.Session{})
	gormdb.DB.AutoMigrate(&models.Subject{})
	gormdb.DB.AutoMigrate(&models.Subscription{})
	gormdb.DB.AutoMigrate(&models.Topic{})
	gormdb.DB.AutoMigrate(&models.TypePayment_Session{})
	gormdb.DB.AutoMigrate(&models.Type_Session{})
	gormdb.DB.AutoMigrate(&models.User{})
	gormdb.DB.AutoMigrate(&models.UserxClassroom{})

	revel.InterceptFunc(AddLog, revel.BEFORE, &App{})
	//	revel.InterceptFunc(Authenticate, revel.BEFORE, &App{})
	revel.InterceptFunc(Authenticate, revel.BEFORE, &User{})
	revel.InterceptFunc(Authenticate, revel.BEFORE, &Session{})
	revel.InterceptFunc(AuthenticateAdmin, revel.BEFORE, &Admin{})
}
