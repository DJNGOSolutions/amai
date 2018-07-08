package controllers

import (
	"fmt"

	"github.com/pdmp/amai/app/models"

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
			  "user".user_email, 
			  "user".user_age, 
			  "user".user_name, 
			  "user".id, 
			  academic_level_user.academic_level, 
			  gender_user.gender, 
			  role_user.role, 
			  rate_user.rate
			FROM 
			  public."user", 
			  public.academic_level_user, 
			  public.role_user, 
			  public.gender_user, 
			  public.rate_user
			WHERE 
			  academic_level_user.id = "user".id_academic_level_user AND
			  role_user.id = "user".id_role_user AND
			  gender_user.id = "user".id_gender_user AND
			  rate_user.id = "user".id_rate_user;
				`).Scan(&model)

	return c.RenderJSON(model)
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
	c.DB.Raw("DELETE FROM public.user WHERE id = ?;", id)
	fmt.Println("id", id)
	return nil
	//return c.Redirect("/crud")
	//return c.RenderJSON(product) //debuging
}

/*
func parseData(userModel UserModel) (uint, uint, uint) {
	var data [3]uint
	if userModel.Gender == "Hombre" {
		data[0] = 1
	} else {
		data[0] = 2
	}

	if userModel.AcademicLevel == "Educacion Media" {
		data[1] = 1
	} else if userModel.AcademicLevel == "Bachiller" {
		data[1] = 2
	} else {
		data[1] = 3
	}

	if userModel.Role == "Alumno" {
		data[2] = 1
	} else {
		data[2] = 2
	}

	return data[0], data[1], data[2]
}
*/
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
