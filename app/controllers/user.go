package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/pdmp/amai/app/models"
	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
)

type User struct {
	gormc.Controller
}

type UserModel struct {
	Id            uint
	UserName      string
	UserAge       uint8
	Gender        string
	UserEmail     string
	AcademicLevel string
	Role          string
	Rate          string
}

var hmacSecret = []byte{97, 48, 97, 50, 97, 98, 105, 49, 99, 102, 83, 53, 57, 98, 52, 54, 97, 102, 99, 12, 12, 13, 56, 34, 23, 16, 78, 67, 54, 34, 32, 21}

func (c User) Show() revel.Result {
	//var users []*models.User
	var model []*UserModel
	//var prods models.Product
	//Gdb.Select("code", "price").Find(&prods).Scan(&users)
	//c.DB.Raw("SELECT * FROM public.user ;").Scan(&users)
	c.DB.Raw(`
			SELECT 
  "user".id, 
  role_user.role, 
  "user".user_name, 
  gender_user.gender, 
  "user".user_email, 
  academic_level_user.academic_level, 
  ratex_user.rate, 
  "user".user_description, 
  "user".user_photo
FROM 
  public.ratex_user, 
  public."user", 
  public.gender_user, 
  public.academic_level_user, 
  public.role_user
WHERE 
  ratex_user.id = "user".id_rate_user AND
  gender_user.id = "user".id_gender_user AND
  academic_level_user.id = "user".id_academic_level_user AND
  role_user.id = "user".id_role_user;
				`).Scan(&model).GetErrors()

	return c.RenderJSON(model)
}

func (c User) Login(email string, pass string) revel.Result {
	var user models.User
	//return c.RenderJSON(c.DB.Where("user_email = ?", email).Find(&user))
	c.DB.Where("user_email = ?", email).Find(&user)
	//return c.RenderJSON(&user)
	fmt.Println("hash", user.UserPassword, "\npass -> "+pass, "\nbyte", []byte(pass))
	if user.UserEmail != "" {
		err := bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(pass))
		if err == nil {
			if user.IdRoleUser == 1 {
				c.Session["user"] = user.UserName
				c.Session.SetDefaultExpiration()
			}
			token := encodeToken(email)
			return c.RenderJSON(token)
		}
	}
	return c.RenderFileName("bento/dist/index.html", revel.NoDisposition)
}

func encodeToken(email string) string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"nbf":   time.Date(2018, 07, 07, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSecret)

	fmt.Println(tokenString, err)

	return tokenString
}

func decodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSecret, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		fmt.Println("email and nbf:", claims["email"], claims["nbf"])
	} else {
		log.Println(err)
		return nil, err
	}
	return claims, nil
	// return claims["email"].(string), claims["nbf"].(string)
}

func (c User) getUser(email string) error {
	var user models.User
	return c.DB.Where("user_email = ?", email).Find(user).Error
	//return c.RenderJSON(product) //debug
}

func (c User) Gender() revel.Result {
	var genders []*models.Gender_User

	c.DB.Find(&genders)

	return c.RenderJSON(genders)
}

func (c User) AcademicLevel() revel.Result {
	var levels []*models.AcademicLevel_User

	c.DB.Find(&levels)

	return c.RenderJSON(levels)
}

func (c User) Category() revel.Result {
	var topics []*models.Topic

	c.DB.Find(&topics)

	return c.RenderJSON(topics)
}

func (c User) Subjects() revel.Result {
	var subjects []*models.Subject

	c.DB.Find(&subjects)

	return c.RenderJSON(subjects)
}

func (c User) getUserByGender(id uint) revel.Result {
	var user models.User
	c.DB.Raw("SELECT User_Photo, User_Name, User_Age, User_Email, User_Description FROM public.User WHERE Id_Gender_User = ?;", id).Scan(&user)
	return c.RenderJSON(c.Response.Status)
}

func (c User) getUserAcademicLevel(id uint) revel.Result {
	var user models.User
	c.DB.Raw("SELECT User_Photo, User_Name, User_Age, User_Email, User_Description FROM public.User WHERE Id_Academic_Level_User = ?;", id).Scan(&user)
	return c.RenderJSON(c.Response.Status)
}

func (c User) getUserByRole(id uint) revel.Result {
	var user models.User
	c.DB.Raw("SELECT User_Photo, User_Name, User_Age, User_Email, User_Description FROM public.User WHERE Id_Role_User = ?;", id).Scan(&user)
	return c.RenderJSON(c.Response.Status)
}

func (c User) getUserByRate(id uint) revel.Result {
	var user models.User
	c.DB.Raw("SELECT User_Photo, User_Name, User_Age, User_Email, User_Description FROM public.User WHERE Id_Rate_User = ?;", id).Scan(&user)
	return c.RenderJSON(c.Response.Status)
}
