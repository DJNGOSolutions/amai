// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tAdmin struct {}
var Admin tAdmin


func (_ tAdmin) CreateSubject(
		name string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "name", name)
	return revel.MainRouter.Reverse("Admin.CreateSubject", args).URL
}

func (_ tAdmin) CreateCategory(
		name string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "name", name)
	return revel.MainRouter.Reverse("Admin.CreateCategory", args).URL
}

func (_ tAdmin) CreateAcademicLevel(
		name string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "name", name)
	return revel.MainRouter.Reverse("Admin.CreateAcademicLevel", args).URL
}

func (_ tAdmin) Pop(
		id uint,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Admin.Pop", args).URL
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).URL
}

func (_ tApp) Login(
		email string,
		pass string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "pass", pass)
	return revel.MainRouter.Reverse("App.Login", args).URL
}

func (_ tApp) Register(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("App.Register", args).URL
}

func (_ tApp) Help(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Help", args).URL
}

func (_ tApp) HashTest(
		pass string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "pass", pass)
	return revel.MainRouter.Reverse("App.HashTest", args).URL
}


type tSession struct {}
var Session tSession


func (_ tSession) Sessions(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Session.Sessions", args).URL
}

func (_ tSession) MySessions(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Session.MySessions", args).URL
}

func (_ tSession) SessionsByTopic(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Session.SessionsByTopic", args).URL
}


type tUser struct {}
var User tUser


func (_ tUser) Show(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Show", args).URL
}

func (_ tUser) Login(
		email string,
		pass string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "pass", pass)
	return revel.MainRouter.Reverse("User.Login", args).URL
}

func (_ tUser) GetUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.GetUser", args).URL
}

func (_ tUser) Gender(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Gender", args).URL
}

func (_ tUser) AcademicLevel(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.AcademicLevel", args).URL
}

func (_ tUser) Category(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Category", args).URL
}

func (_ tUser) Subjects(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Subjects", args).URL
}


type tController struct {}
var Controller tController



type tTxnController struct {}
var TxnController tTxnController


func (_ tTxnController) Begin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TxnController.Begin", args).URL
}

func (_ tTxnController) Commit(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TxnController.Commit", args).URL
}

func (_ tTxnController) Rollback(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TxnController.Rollback", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeDir(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeDir", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}

func (_ tStatic) ServeModuleDir(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModuleDir", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


