package controllers

import (
	"fmt"

	"github.com/pdmp/amai/app/models"
	"golang.org/x/crypto/bcrypt"

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
				`).Scan(&model)

	return c.RenderJSON(model)
}

func (c User) GenHash(pass string, id uint) revel.Result {

	var user models.User
	c.DB.Where("id = ?", id).Find(&user)

	user.Hash, _ = bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	c.DB.Save(&user)
	return c.RenderJSON(user.Hash)
}

func (c User) Login(email string, pass string) revel.Result {
	var user models.User
	//return c.RenderJSON(c.DB.Where("user_email = ?", email).Find(&user))
	c.DB.Where("user_email = ?", email).Find(&user)
	//return c.RenderJSON(&user)
	fmt.Println("hash", user.Hash, "\npass -> "+pass, "\nbyte", []byte(pass))
	if user.UserEmail != "" {
		err := bcrypt.CompareHashAndPassword(user.Hash, []byte(pass))
		if err == nil {
			if user.IdRoleUser == 1 {
				c.Session["user"] = user.UserName
				c.Session.SetDefaultExpiration()
			}
			token, _ := bcrypt.GenerateFromPassword([]byte(user.UserPassword+user.UserName), bcrypt.DefaultCost)
			return c.RenderJSON(token)
		}
	}
	return c.RenderFileName("bento/dist/index.html", revel.NoDisposition)
}

func (c User) Crud() revel.Result {
	return c.RenderFileName("bento/dist/index.html", revel.NoDisposition)
}

func (c User) GetUser(code string) revel.Result {
	var user models.User
	c.DB.Raw("SELECT * FROM public.user WHERE id = ?;", code).Scan(&user)
	//return c.RenderJSON(product) //debug
	return c.RenderJSON(user)
}

func (c User) Delete() revel.Result {
	var users []models.User
	c.DB.Raw("SELECT UserName FROM public.user").Scan(&users)
	//return c.RenderJSON(users) //debugging
	return c.Render(users)
}

func (c User) Pop(id uint) revel.Result {
	//var user models.User
	//c.DB.Raw(`DELETE FROM public."user" WHERE id = ?;`, id)
	//return c.RenderJSON(c.DB.Raw(`DELETE FROM public."user" WHERE id = ?;`, id))
	return c.RenderJSON(c.DB.Unscoped().Delete(models.User{}, "id = ?", id))
	//return c.Redirect("/crud")
	//return c.RenderJSON(product) //debuging
}

func (c User) Insert(user models.User) revel.Result {
	//gender, level, role := parseData(userModel)

	/*
		user := models.User{IdGenderUser: gender,
			IdAcademicLevelUser: level, IdRoleUser: role,
			IdRateUser: 1, UserPhoto: "None", UserName: userModel.UserName,
			UserAge: userModel.UserAge, UserEmail: userModel.UserEmail,
			UserDescription: ""}
	*/
	/*
		log := c.Log.New("insert", 1)
		log.Debug("Inserting:", userModel)
		log.Debug("Inserting:", user)*/

	c.DB.Create(&user)
	//log.Debug("Inserting:", user.UserName, user.UserAge, user.UserEmail, nil)

	//return nil
	return c.RenderJSON(c.Response.Status)
	//return c.Redirect(routes.App.Index())
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
