package controllers

/*
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

type SessionsModel struct {
	UserEmail         string
	SessionDate       string
	SessionTime_Start string
	SessionTime_End   string
}

var hmacSecret = []byte{97, 48, 97, 50, 97, 98, 105, 49, 99, 102, 83, 53, 57, 98, 52, 54, 97, 102, 99, 12, 12, 13, 56, 34, 23, 16, 78, 67, 54, 34, 32, 21}

var query = `
			SELECT
  "user".id,
  role_user.role,
  "user".user_name,
  gender_user.gender,
  "user".user_email,
  academic_level_user.academic_level,
  "user".user_description,
  "user".user_photo
FROM
  public.session,
  public."user",
  public.gender_user,
  public.academic_level_user,
  public.role_user
WHERE
  gender_user.id = "user".id_gender_user AND
  academic_level_user.id = "user".id_academic_level_user AND
  role_user.id = "user".id_role_user
`

func (c User) Show() revel.Result {
	//var users []*models.User
	var model []*UserModel
	//var prods models.Product
	//Gdb.Select("code", "price").Find(&prods).Scan(&users)
	//c.DB.Raw("SELECT * FROM public.user ;").Scan(&users)
	c.DB.Raw(query).Scan(&model).GetErrors()

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

func (c User) GetUser() revel.Result {
	token := c.Params.Form.Get("token")
	var claims jwt.MapClaims
	var err error
	claims, err = decodeToken(token)
	if err != nil {
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON("auth failed")
	}

	email, found := claims["email"]
	if !found {
		log.Println(err)
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON("email not found in db")
	}

	log.Println("email found:", email)
	var user UserModel
	err = c.DB.Raw(query+" AND user_email = ?", email).Scan(&user).Error
	log.Println(query+" AND user_email =", email)
	//err = gormdb.DB.Where("user_email = ?", email).First(&user).Error

	if err != nil {
		log.Println(err)
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON("auth failed")
	}

	return c.RenderJSON(&user)
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

func (c User) Sessions() revel.Result {
	var sessions []*SessionsModel

	c.DB.Raw(`
SELECT user_email, session_date, session_time_start, session_time_end
FROM public.user, assistance, session
WHERE public.user.id = assistance.id_user AND assistance.id_user = session.id
	`).Find(&sessions)

	return c.RenderJSON(sessions)
}

type ClassroomModel struct {
	Id          int
	UserName    string
	TopicName   string
	IdSubject   int
	SubjectName string
}

func (c User) Classroom() revel.Result {
	var classroom []*ClassroomModel

	c.DB.Raw(`
		SELECT
		  userx_classroom.id,
		  "user".user_name,
		  topic.topic_name,
		  classroom.id_subject,
		  subject.subject_name
		FROM
		  public.userx_classroom,
		  public.classroom,
		  public."user",
		  public.topic,
		  public.subject
		WHERE
		  classroom.id = userx_classroom.id_classroom AND
		  classroom.id_topic = topic.id AND
		  "user".id = userx_classroom.id_user AND
		  topic.id_subject = subject.id;
	`).Scan(&classroom)

	return c.RenderJSON(classroom)

}*/
