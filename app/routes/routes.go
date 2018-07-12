// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tAdmin struct {}
var Admin tAdmin



type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).URL
}

func (_ tApp) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Login", args).URL
}

func (_ tApp) Register(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Register", args).URL
}

func (_ tApp) Insert(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Insert", args).URL
}

func (_ tApp) Image(
		name string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "name", name)
	return revel.MainRouter.Reverse("App.Image", args).URL
}


type tSession struct {}
var Session tSession



type tUser struct {}
var User tUser


func (_ tUser) Show(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Show", args).URL
}

func (_ tUser) GetUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.GetUser", args).URL
}

func (_ tUser) Session(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Session", args).URL
}

func (_ tUser) SubscribeSession(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.SubscribeSession", args).URL
}

func (_ tUser) DeleteSession(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.DeleteSession", args).URL
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


